
// Only ONE Hub Form
let form_hub = `
<div v-if="mPV.get('Hub').value">    

    <form class="cfgform">

      <label class="lb">{{label.name[0]}}: </label>
      <input v-model="input[0].name" type="text" :placeholder="label.name[1]">   

      <label class="lb">{{label.path[0]}}: </label>
      <input v-model="input[0].path" type="text" :placeholder="label.path[1]">   

      <label class="lb">{{label.args[0]}}:</label>
      <input v-model="input[0].args" type="text" :placeholder="label.args[1]">  

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

      <br/> 

      <label class="lb">{{label.sel_benthos[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_benthos[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Benthos').value" value="">{{cn}}</option>
      </select>

      <br/>     

      <label class="lb">{{label.sel_reader_align_map[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_reader_align_map[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Reader').value" value="">{{cn}}</option>
      </select>

      <br/>     

      <label class="lb">{{label.sel_reader_level_map[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_reader_level_map[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Reader').value" value="">{{cn}}</option>
      </select>

      <br/>     

      <label class="lb">{{label.sel_reader[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_reader[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Reader').value" value="">{{cn}}</option>
      </select>

      <br/>     

      <label class="lb">{{label.sel_txtclassifier[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_txtclassifier[1]}}</option>
        <option v-for="(cn, i) in mPN.get('TxtClassifier').value" value="">{{cn}}</option>
      </select>

      <br/>

      <label class="lb">{{label.sel_level[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_level[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Level').value" value="">{{cn}}</option>
      </select>

      <br/>

      <label class="lb">{{label.sel_weight[0]}}: </label>
      <select class="selector">
        <option value="" disabled selected>{{label.sel_weight[1]}}</option>
        <option v-for="(cn, i) in mPN.get('Weight').value" value="">{{cn}}</option>
      </select>

    </form>
</div>
`;

export function getForm_Hub() {
  return form_hub;
}
