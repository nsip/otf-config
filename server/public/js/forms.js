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

function reset_v() {
  mPV.set('NatsStreaming', Vue.ref(false));
  mPV.set('Nias3', Vue.ref(false));
  mPV.set('Benthos', Vue.ref(false));
  mPV.set('Reader', Vue.ref(false));
  mPV.set('Align', Vue.ref(false));
  mPV.set('TxtClassifier', Vue.ref(false));
  mPV.set('Level', Vue.ref(false));
  mPV.set('Weight', Vue.ref(false));
  mPV.set('Hub', Vue.ref(false));
}

reset_v();

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

let ExePathGrp = [];
let ArgsGrp = [];
let DelayGrp = [];

function fill_table(proj, name) {
  (async () => {
    await sleep(20);
    const all = await get_allitem();
    (async function inflate(data) {
      for (let i = 0; i < data.length; i++) {
        await sleep(10);
        const b = await get_cfg(proj, data[i]);
        if (b.name == name) {

          console.log(b.path);

          ExePathGrp.push(b.path);
          console.log(ExePathGrp);
        }
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
      post_cfg(selproj, input.value[0]); // send input new form to back-end
      clr_new_form(input); // clear new form
      emitter.emit("selected", selproj); // refresh current form
    }

    // update button
    function btn_update(selproj, i) {
      console.log(`update ${selproj} on ${i} form`);
      // console.log(input.value[i])
      put_cfg(selproj, input.value[i]);
    }

    // delete button
    function btn_delete(selproj, i) {
      console.log(`delete ${selproj} on ${i} form`);
      // console.log(input.value[i])
      delete_cfg(selproj, input.value[i].name);
      emitter.emit("selected", selproj); // refresh current form
    }

    let disable_btn = Vue.ref(false);

    let nReader = Vue.ref([""]);

    function btn_add_reader() {
      nReader.value.push("");
    }

    function btn_remove_reader() {
      nReader.value.pop();
    }

    let com_natsstreaming = Vue.ref("");
    let com_nias3 = Vue.ref("");
    let com_benthos_align = Vue.ref("");
    let com_benthos_level = Vue.ref("");
    let com_benthos = Vue.ref("");
    let com_reader_align = Vue.ref("");
    let com_reader_level = Vue.ref("");
    let com_reader = Vue.ref([""]);
    let com_txtclassifier = Vue.ref("");
    let com_level = Vue.ref("");
    let com_weight = Vue.ref("");

    function btn_composite() {
      alert(com_reader.value);

      fill_table("NatsStreaming", com_natsstreaming.value)
      fill_table("Nias3", com_nias3.value)
      fill_table("Benthos", com_benthos_align.value)
      fill_table("Benthos", com_benthos_level.value)
      fill_table("Benthos", com_benthos.value)
      fill_table("Reader", com_reader_align.value)
      fill_table("Reader", com_reader_level.value)
      for (let i = 0; i < com_reader.value.length; i++) {
        fill_table("Reader", com_reader.value[i])
      }
      fill_table("TxtClassifier", com_txtclassifier.value)

      // fill_table("", com_natsstreaming.value)
      // fill_table("", com_natsstreaming.value)

      // console.log(ExePathGrp);

    }

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
      nReader,
      btn_add_reader,
      btn_remove_reader,
      btn_composite,
      com_natsstreaming,
      com_nias3,
      com_benthos_align,
      com_benthos_level,
      com_benthos,
      com_reader_align,
      com_reader_level,
      com_reader,
      com_txtclassifier,
      com_level,
      com_weight,
    };
  },

  template: getForm(),
};
