import { getEmitter } from "./js/mitt.js";
import { get_allitem, get_cfg } from "./js/fetch.js";
import { getForm } from "./form/all.js";

const emitter = getEmitter();

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

function clrform(input) {
  input.value = [
    {
      name: "",
      path: "",
      svrname: "",
      svrid: "",
      provider: "",
      inputfmt: "",
      alignmethod: "",
      levelmethod: "",
      gencapability: "",
      natshost: "",
      natsport: "",
    }
  ];
}

export default {
  setup() {
    let selected = Vue.ref(false);
    let title = Vue.ref("OTF project (select from left)");

    let vi = Vue.reactive({
      natsstreaming: false,
      nias3: false,
      benthos: false,
      reader: false,
      align: false,
      txtclassifier: false,
      level: false,
      weight: false,
      hub: false,
    });

    // define all input
    let input = Vue.ref([
      // init an empty one for the first new form
      {
        name: "",
        path: "",
        svrname: "",
        svrid: "",
        provider: "",
        inputfmt: "",
        alignmethod: "",
        levelmethod: "",
        gencapability: "",
        natshost: "",
        natsport: "",
      }
    ]);

    // listen to an event
    emitter.on("selected", (e) => {
      // test
      console.log("forms received:", e);

      // selected
      selected.value = true;

      // change title
      title.value = `OTF - ${e}`;

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
        clrform(input);

        // console.log(a);
        // console.log(a[arg]);

        // fetch config content array
        all[arg].forEach((cname) => {
          (async () => {
            const b = await get_cfg(e, cname);
            console.log(b);

            // assign from fetch
            input.value.push({
              name: b.name,
              path: b.path,
              svrname: b.svrname,
              svrid: b.svrid,
              provider: b.provider,
              inputfmt: b.inputFormat,
              alignmethod: b.alignMethod,
              levelmethod: b.levelMethod,
              gencapability: b.capability,
              natshost: b.natsHost,
              natsport: b.natsPort,
            });

          })();
        });
      })();
    });

    return {
      selected,
      title,
      vi,
      input,
    };
  },

  template: getForm(),
};
