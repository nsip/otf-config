
// Only ONE Hub Form

let form_hub = `
<div v-if="vi.hub">    
    <form class="cfgform">

      <label class="lb">{{label.name[0]}}: </label>
      <input v-model="input[0].name" type="text" :placeholder="label.name[1]">   

      <label class="lb">{{label.path[0]}}: </label>
      <input v-model="input[0].path" type="text" :placeholder="label.path[1]">   

      <input type="button" value="update" :disabled="disable_btn" @click="btn_update(selproj, 0)">

      <p></p>

      <label class="lb">{{label.sel_natsstreaming[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_natsstreaming[1]}}</option>
        <option v-for="(cn, i) in mPN.get('NatsStreaming').value" value="">{{cn}}</option>
      </select>

      <br/>     

      <label class="lb">{{label.sel_nias3[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_nias3[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Nias3').value" value="">{{cn}}</option>
      </select>

    </form>
</div>
`;

export function getForm_Hub() {
  return form_hub;
}
