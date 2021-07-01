import { getEmitter } from "./js/mitt.js";
import { get_allitem, get_cfg, post_cfg, put_cfg } from "./js/fetch.js";
import { getForm } from "./form/all.js";

const emitter = getEmitter();

const VisForm = {
  natsstreaming: false,
  nias3: false,
  benthos: false,
  reader: false,
  align: false,
  txtclassifier: false,
  level: false,
  weight: false,
  hub: false,
}

function viform(proj, vi) {

  vi.natsstreaming = false;
  vi.nias3 = false;
  vi.benthos = false;
  vi.reader = false;
  vi.align = false;
  vi.txtclassifier = false;
  vi.level = false;
  vi.weight = false;
  vi.hub = false;

  switch (proj) {
    case "NatsStreaming":
      vi.natsstreaming = true;
      break;
    case "Nias3":
      vi.nias3 = true;
      break;
    case "Benthos":
      vi.benthos = true;
      break;
    case "Reader":
      vi.reader = true;
      break;
    case "Align":
      vi.align = true;
      break;
    case "TxtClassifier":
      vi.txtclassifier = true;
      break;
    case "Level":
      vi.level = true;
      break;
    case "Weight":
      vi.weight = true;
      break;
    case "Hub":
      vi.hub = true;
      break;
  }
}

const InitLabel = {
  // all
  name: ["Config File Name", "config file name, should be identical"],
  path: ["Path to Service Executable", "path/to/service/executable"],

  // all services
  svrname: ["Service Name", "name for service"],
  svrid: ["Service ID", "id for service, leave blank to auto-generate a unique id"],

  // reader
  provider: ["Provider Name", "name of product or system supplying the data"],
  inputfmt: ["Input File Format", "format of input data, one of csv|json"],
  alignmethod: ["Align Method", "method to align input data to NLPs must be one of prescribed|mapped|inferred"],
  levelmethod: ["Level Method", "method to apply common scaling this data, one of prescribed|mapped-scale|rules"],
  gencapability: ["General Capability", "General Capability for assessment results; Literacy or Numeracy"],
  natshost: ["Nats Streaming Host", "hostname/ip of nats broker"],
  natsport: ["Nats Port", "connection port for nats broker"],
  natscluster: ["Nats Cluster Name", "cluster id for nats broker"],
  topic: ["Nats Topic", "nats topic name to publish parsed data items to"],
  folder: ["Watching Folder", "folder to watch for data files"],
  filesuffix: ["Only Watch File Suffix", "filter files to read by file extension, eg. .csv or .myapp (actual data handling will be determined by input format flag)"],
  interval: ["Interval For Dealing", "watcher poll interval"],
  recursive: ["Dealing With Sub Folder", "watch folders recursively"],
  dotfiles: ["Dealing With Dot Files", "watch dot files"],
  ignore: ["Folders To Be Ignored", "comma separated list of paths to ignore"],
  concurrfiles: ["Files' Count In Once Process", "pool size for concurrent file processing"],

  // align, level, weight ...
  port: ["Service Port", "current service running port"],

  // align, level ...
  niashost: ["NIAS3 Host", "hostname/ip of nias3"],
  niasport: ["NIAS3 Port", "connection port for nias3"],
  niastoken: ["NIAS3 Token", "token to access nias3"],

  // align
  tchost: ["Text Classifier Host", "text classifier service hostname/ip"],
  tcport: ["Text Classifier Port", "text classifier service connection port"],

  // text classifier

  // weight
  failwhenerr: ["Panic If Error", "if error happens, should service abort?"],
}

const InitInput = {
  // all
  name: "",
  path: "",

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
  port: "",

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

}

function inflateform(input, data) {

  // data fields name refer to 'config.go'
  input.value.push({
    // all
    name: data.name,
    path: data.path,

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
    concurrfiles: data.concurrFiles,

    // align, level, weight ...
    port: data.port,

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
  input.value[0].svrname = "";
  input.value[0].svrid = "";
  input.value[0].port = "";
}

export default {
  setup() {

    let selected = Vue.ref(false);

    let selproj = Vue.ref("");

    let title = Vue.ref("OTF project (select from left)");

    let vi = Vue.reactive(VisForm);

    // init an empty one for the first new form
    let input = Vue.ref([InitInput]);

    // listen to an event
    emitter.on("selected", (e) => {

      // test
      console.log("forms received:", e);

      // selected
      selected.value = true;

      // change title
      title.value = `OTF - ${e}`;

      // clear input form if change to another project config
      if (e != selproj.value) {
        clr_new_form(input);
      }

      // select project
      selproj.value = e;

      // set visibility of each project config
      viform(e, vi);

      // fetch all selected config
      (async () => {
        let arg = `${e}s`;
        if (e == "Benthos") {
          arg = `${e}es`;
        }
        const all = await get_allitem();

        // clear all existing input
        clr_form(input);

        // console.log(a);
        // console.log(a[arg]);

        // fetch config content array
        all[arg].forEach((cname) => {
          (async () => {
            const b = await get_cfg(e, cname);
            console.log(b);

            // fill existing form input with fetch data
            inflateform(input, b);

          })();
        });
      })();
    });

    // new button
    function btn_new(selproj) {
      console.log(`new ${selproj}`);
      post_cfg(selproj, input.value[0]); // send input new form to backend
      clr_new_form(input);               // clear new form
      emitter.emit("selected", selproj); // refresh current form
    }

    // update button
    function btn_update(selproj, i) {
      console.log(`update ${selproj} on ${i} form`);
      put_cfg(selproj, input.value[i]);
      emitter.emit("selected", selproj); // refresh current form
    }

    let disable_btn = Vue.ref(false);

    return {
      selected,
      selproj,
      title,
      vi,
      label: InitLabel,
      input,
      disable_btn,
      btn_new,
      btn_update,
    };
  },

  template: getForm(),
};
