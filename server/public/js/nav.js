// port 1555 is set in config.toml

function url(api) {
  return `http://127.0.0.1:1555/` + api;
}

let urlAllGrp = url("allgrp");
let urlAllItems = url("allitems");
let urlReader = url("otf-config/reader?cfgName=spa_prescribed_config");

function test(url) {
  const data = fetch(url)
    .then((resp) => resp.json())
    .then((data) => {
      // console.log(data);
      return data;
    });

  //   const data = async () => {
  //     const a = await prom;
  //     console.log(a);
  //     return a;
  //   };

  return data;
}

const allGrp = [
  "NatsStreaming",
  "Nias3",
  "Benthos",
  "Reader",
  "Align",
  "TxtClassifier",
  "Level",
  "Weight",
  "Hub",
];

const app = Vue.createApp({
  data() {
    return {
      grps: allGrp,
      title: "OTF All-In-One Config",
      cfgname: "",
      nfVisible: false,
      content: "",
    };
  },
  methods: {
    show(str) {
      this.nfVisible = true;

      // console.log(str);
      // alert(str);

      switch (str) {
        case "NatsStreaming":
          str += "s";
          content = '<label class="lb">config name:</label>'
          break;
        case "Nias3":
          str += "s";
          break;
        case "Benthos":
          str += "es";
          break;
        case "Reader":
          str += "s";
          break;
        default:
          str += "s";
      }

      (async () => {
        data = await test(urlAllItems);
        console.log(data[str]);
        this.cfgname = data[str];
      })();
    },
  },
}).mount("#app");
