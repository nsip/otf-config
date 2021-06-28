import { getEmitter } from './js/mitt.js'

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

function reset_nav() {
  allGrp.forEach(e => {
    document.getElementById(e).innerText = e;
  })
}

const emitter = getEmitter();

export default {
  // name: 'nav',

  setup() {

    function sel2form(str) {

      // mark nav selected item
      reset_nav();
      document.getElementById(str).innerText = str + " *";

      // send to main form
      emitter.emit("selected", str);
    }

    return {
      grps: allGrp,
      sel2form,
    };
  },

  template: `
  <div id="nav" class="sidenav">
  <a href="#" v-for="grp in grps" :id="grp" @click="this.sel2form(grp)">{{grp}}</a>
  </div>
  `,
}
