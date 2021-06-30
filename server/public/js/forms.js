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

    let names_in = Vue.ref([""]); // init an empty one for the first new form
    let paths_in = Vue.ref([""]);
    let svrnames_in = Vue.ref([""]);
    let svrids_in = Vue.ref([""]);
    let providers_in = Vue.ref([""]);
    let inputfmts_in = Vue.ref([""]);
    let alignmethods_in = Vue.ref([""]);
    let levelmethods_in = Vue.ref([""]);
    let gencapabilities_in = Vue.ref([""]);
    let natshosts_in = Vue.ref([""]);
    let natsports_in = Vue.ref([""]);

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

        // clear ***_in for inflating again
        names_in.value = [""]; // init an empty one for the first new form
        paths_in.value = [""];
        svrnames_in.value = [""];
        svrids_in.value = [""];
        providers_in.value = [""];
        inputfmts_in.value = [""];
        alignmethods_in.value = [""];
        levelmethods_in.value = [""];
        gencapabilities_in.value = [""];
        natshosts_in = Vue.ref([""]);
        natsports_in = Vue.ref([""]);

        // console.log(a);
        // console.log(a[arg]);

        // fetch config content array
        all[arg].forEach((cname) => {
          (async () => {
            const b = await get_cfg(e, cname);
            console.log(b);

            names_in.value.push(b.name);
            paths_in.value.push(b.path);
            svrnames_in.value.push(b.svrname);
            svrids_in.value.push(b.svrid);
            providers_in.value.push(b.provider);
            inputfmts_in.value.push(b.inputFormat);
            alignmethods_in.value.push(b.alignMethod);
            levelmethods_in.value.push(b.levelMethod);
            gencapabilities_in.value.push(b.capability);
            natshosts_in.value.push(b.natsHost);
            natsports_in.value.push(b.natsPort);
          })();
        });
      })();
    });

    return {
      selected,
      title,
      vi,
      names_in,
      paths_in,
      svrnames_in,
      svrids_in,
      providers_in,
      inputfmts_in,
      alignmethods_in,
      levelmethods_in,
      gencapabilities_in,
      natshosts_in,
      natsports_in,
    };
  },

  template: getForm(),
};
