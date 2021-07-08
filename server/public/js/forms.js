import { getEmitter } from "./js/mitt.js";
import { get_allitem, get_cfg, post_cfg, put_cfg, delete_cfg } from "./js/fetch.js";
import { getForm } from "./form/all.js";
import { getLabels } from "./js/label.js";

const emitter = getEmitter();

const InitInput = {
  // all
  name: "",
  path: "",
  args: "",

  // all services
  svrname: "",
  svrid: "",

  // reader
  provider: "",
  inputfmt: "",
  alignmethod: "",
  levelmethod: "",
  gencapability: "",
  natshost: "",
  natsport: "",
  natscluster: "",
  topic: "",
  folder: "",
  filesuffix: "",
  interval: "",
  recursive: false,
  dotfiles: false,
  ignore: "",
  concurrfiles: 10,

  // align, level, weight ...
  port: 0,

  // align, level ...
  niashost: "",
  niasport: "",
  niastoken: "",

  // align
  tchost: "",
  tcport: "",

  // text classifier

  // weight
  failwhenerr: false,
};

function inflate_form(input, data) {
  // data fields name refer to 'config.go'
  input.value.push({
    // all
    name: data.name,
    path: data.path,
    args: data.args,

    // all services
    svrname: data.svrname,
    svrid: data.svrid,

    // reader
    provider: data.provider,
    inputfmt: data.inputFormat,
    alignmethod: data.alignMethod,
    levelmethod: data.levelMethod,
    gencapability: data.capability,
    natshost: data.natsHost,
    natsport: data.natsPort,
    natscluster: data.natsCluster,
    topic: data.topic,
    folder: data.folder,
    filesuffix: data.suffix,
    interval: data.interval,
    recursive: data.recursive,
    dotfiles: data.dotfiles,
    ignore: data.ignore,
    concurrfiles: parseInt(data.concurrFiles),

    // align, level, weight ...
    port: parseInt(data.port),

    // align, level ...
    niashost: data.niasHost,
    niasport: data.niasPort,
    niastoken: data.niasToken,

    // align
    tchost: data.tcHost,
    tcport: data.tcPort,

    // weight
    failwhenerr: data.failWhenErr,
  });
}

function clr_form(input) {
  input.value = [InitInput];
}

// ****************************************************************************************** adding more
function clr_new_form(input) {
  input.value[0].name = "";
  input.value[0].path = "";
  input.value[0].args = "";
  input.value[0].svrname = "";
  input.value[0].svrid = "";
  input.value[0].port = 0;
}

////////////////////////////////////////////////////////////////////////////////////

const mPV = new Map();
mPV.set('NatsStreaming', Vue.ref(false));
mPV.set('Nias3', Vue.ref(false));
mPV.set('Benthos', Vue.ref(false));
mPV.set('Reader', Vue.ref(false));
mPV.set('Align', Vue.ref(false));
mPV.set('TxtClassifier', Vue.ref(false));
mPV.set('Level', Vue.ref(false));
mPV.set('Weight', Vue.ref(false));
mPV.set('Hub', Vue.ref(false));

function reset_v() {
  mPV.get('NatsStreaming').value = false;
  mPV.get('Nias3').value = false;
  mPV.get('Benthos').value = false;
  mPV.get('Reader').value = false;
  mPV.get('Align').value = false;
  mPV.get('TxtClassifier').value = false;
  mPV.get('Level').value = false;
  mPV.get('Weight').value = false;
  mPV.get('Hub').value = false;
}

function visible(proj) {
  reset_v();
  mPV.get(proj).value = true;
}

////////////////////////////////////////////////////////////////////////////////////

const mPN = new Map();
mPN.set('NatsStreaming', Vue.ref([]));
mPN.set('Nias3', Vue.ref([]));
mPN.set('Benthos', Vue.ref([]));
mPN.set('Reader', Vue.ref([]));
mPN.set('Align', Vue.ref([]));
mPN.set('TxtClassifier', Vue.ref([]));
mPN.set('Level', Vue.ref([]));
mPN.set('Weight', Vue.ref([]));

function get_dropcontent(proj) {
  mPN.get(proj).value = [];
  (async () => {
    await sleep(20);
    const all = await get_allitem();
    (async function inflate(data) {
      for (let i = 0; i < data.length; i++) {
        await sleep(10);
        const b = await get_cfg(proj, data[i]);
        mPN.get(proj).value.push(b.name);
      }
    })(all[proj]);
  })();
}

////////////////////////////////////////////////////////////////////////////////////

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms))

export default {
  setup() {

    let selproj = Vue.ref("");

    let title = Vue.ref("OTF project (select from left)");

    // init an empty one for the first new form
    let input = Vue.ref([InitInput]);

    // in one page, which form can be seen
    let vf = Vue.ref([true]);

    function collapse(i) {
      vf.value[i] = !vf.value[i];
    }

    // listen to an event
    emitter.on("selected", (e) => {

      // test
      console.log("forms received:", e);

      // change title
      title.value = `OTF - ${e}`;

      // clear input form if change to another project config
      if (e != selproj.value) {
        clr_new_form(input);
      }

      // select project
      selproj.value = e;

      // set visibility of each project config     
      visible(e);

      // fetch all selected config
      (async () => {
        await sleep(20);
        const all = await get_allitem();

        // clear all existing input
        clr_form(input);

        // console.log(a);
        // console.log(a[e]);

        (async function inflate(data) {
          for (let i = 0; i < data.length; i++) {
            await sleep(10);
            const b = await get_cfg(e, data[i]);
            inflate_form(input, b);
          }
        })(all[e]);

      })();

      /////////////////////////////////////////////

      // refresh selector
      if (e == "Hub") {
        get_dropcontent('NatsStreaming');
        get_dropcontent('Nias3');
        get_dropcontent('Benthos');
        get_dropcontent('Reader');
        get_dropcontent('Align');
        get_dropcontent('TxtClassifier');
        get_dropcontent('Level');
        get_dropcontent('Weight');
      }

      // keep the 1st form visible on project changed
      vf.value = [true];
    });

    // new button
    function btn_new(selproj) {
      console.log(`new ${selproj}`);
      post_cfg(selproj, input.value[0]); // send input new form to backend
      clr_new_form(input); // clear new form
      emitter.emit("selected", selproj); // refresh current form
    }

    // update button
    function btn_update(selproj, i) {
      console.log(`update ${selproj} on ${i} form`);
      // console.log(input.value[i])
      put_cfg(selproj, input.value[i]);
      emitter.emit("selected", selproj); // refresh current form
    }

    // delete button
    function btn_delete(selproj, i) {
      console.log(`delete ${selproj} on ${i} form`);
      // console.log(input.value[i])
      delete_cfg(selproj, input.value[i].name);
      emitter.emit("selected", selproj); // refresh current form
    }

    let disable_btn = Vue.ref(false);

    return {
      selproj,
      title,
      label: getLabels(),
      input,
      disable_btn,
      btn_new,
      btn_update,
      btn_delete,
      mPV,
      mPN,
      vf,
      collapse,
    };
  },

  template: getForm(),
};
