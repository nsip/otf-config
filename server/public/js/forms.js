import { getEmitter } from "./js/mitt.js";
import { get_allitem, get_cfg, post_cfg, put_cfg, delete_cfg, post_table } from "./js/fetch.js";
import { getForm } from "./form/all.js";
import { getLabels } from "./js/label.js";

const emitter = getEmitter();

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

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

let all_data;

function refresh_all_data() {
  (async () => {
    all_data = await get_allitem();
  })();
}

refresh_all_data();

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
  (async (data) => {
    for (let i = 0; i < data.length; i++) {
      await sleep(10);
      const b = await get_cfg(proj, data[i]);
      mPN.get(proj).value.push(b.name);
    }
  })(all_data[proj]);
}

////////////////////////////////////////////////////////////////////////////////////

const TableCols = {
  ExePathGrp: ["PATH_OF_SERVICE", "L"],
  ArgsGrp: ["ARGUMENTS", "L"],
  DelayGrp: ["DELAY", "M"],
  Enabled: ["ENABLED", "M"],
}

function fill_table(proj, name) {
  (async (data) => {
    for (let i = 0; i < data.length; i++) {
      await sleep(10);
      const b = await get_cfg(proj, data[i]);
      if (b.name == name) {
        TableCols.ExePathGrp.push(b.path);
        TableCols.ArgsGrp.push(b.args);
        TableCols.DelayGrp.push(b.delay);
        TableCols.Enabled.push(true);
      }
    }
  })(all_data[proj]);
}

function clear_table() {
  TableCols.ExePathGrp = ["PATH_OF_SERVICE", "L"];
  TableCols.ArgsGrp = ["ARGUMENTS", "L"];
  TableCols.DelayGrp = ["DELAY", "M"];
  TableCols.Enabled = ["ENABLED", "M"];
}

////////////////////////////////////////////////////////////////////////////////////

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
        await sleep(10);

        await refresh_all_data();

        // clear all existing input
        clr_form(input);

        (async (data) => {
          for (let i = 0; i < data.length; i++) {
            await sleep(10);
            const b = await get_cfg(e, data[i]);
            inflate_form(input, b);
          }
        })(all_data[e]);

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

    let com = {
      natsstreaming: Vue.ref(""),
      nias3: Vue.ref(""),
      benthos_align: Vue.ref(""),
      benthos_level: Vue.ref(""),
      benthos: Vue.ref(""),
      reader_align: Vue.ref(""),
      reader_level: Vue.ref(""),
      reader: Vue.ref([""]),
      align: Vue.ref(""),
      txtclassifier: Vue.ref(""),
      level: Vue.ref(""),
      weight: Vue.ref(""),
    }

    function isEmpty(str) {
      return (!str || str.length === 0);
    }

    function hasEmptyStr(arr) {
      for (let i = 0; i < nReader.value.length; i++) {
        if (isEmpty(arr[i])) {
          return true
        }
      }
      return false
    }

    function com_invalid() {
      return isEmpty(com.natsstreaming.value) ||
        isEmpty(com.nias3.value) ||
        isEmpty(com.benthos_align.value) ||
        isEmpty(com.benthos_level.value) ||
        isEmpty(com.benthos.value) ||
        isEmpty(com.reader_align.value) ||
        isEmpty(com.reader_level.value) ||
        hasEmptyStr(com.reader.value) ||
        isEmpty(com.align.value) ||
        isEmpty(com.txtclassifier.value) ||
        isEmpty(com.level.value) ||
        isEmpty(com.weight.value);
    }

    function btn_composite() {
      // alert(com.natsstreaming.value);

      clear_table();

      fill_table("NatsStreaming", com.natsstreaming.value);
      fill_table("Nias3", com.nias3.value);
      fill_table("Benthos", com.benthos_align.value);
      fill_table("Benthos", com.benthos_level.value);
      fill_table("Benthos", com.benthos.value);
      fill_table("Reader", com.reader_align.value);
      fill_table("Reader", com.reader_level.value);
      for (let i = 0; i < com.reader.value.length; i++) {
        fill_table("Reader", com.reader.value[i]);
      }
      fill_table("Align", com.align.value);
      fill_table("TxtClassifier", com.txtclassifier.value);
      fill_table("Level", com.level.value);
      fill_table("Weight", com.weight.value);

      (
        async () => {
          await sleep(200);
          post_table(TableCols);
          // console.log(TableCols);
        }
      )();
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
      com,
      com_invalid,
    };
  },

  template: getForm(),
};
