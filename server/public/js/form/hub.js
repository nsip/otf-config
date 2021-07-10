
// Only ONE Hub Form
let form_hub = `
<div v-if="mPV.get('Hub').value">    

    <form class="cfgform">

      <label class="lb">{{label.path[0]}}: </label>
      <input v-model.trim="input[0].path" type="text" :placeholder="label.path[1]">   

      <label class="lb">{{label.args[0]}}:</label>
      <input v-model.trim="input[0].args" type="text" :placeholder="label.args[1]">  
      
      <!-- ------------------------------------------------------------------------------------ -->

      <p></p>

      <!-- NatsStreaming -->
      <div>
        <label class="lb">{{label.sel_natsstreaming[0]}}: </label>
        <select v-model="com_natsstreaming" class="selector">
          <option value="" disabled selected>{{label.sel_natsstreaming[1]}}</option>
          <option v-for="(cn, i) in mPN.get('NatsStreaming').value" :value="cn">{{cn}}</option>
        </select>
      </div>
     

      <!-- Nias3 -->
      <div>
        <label class="lb">{{label.sel_nias3[0]}}: </label>
        <select v-model="com_nias3" class="selector">
          <option value="" disabled selected>{{label.sel_nias3[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Nias3').value" :value="cn">{{cn}}</option>
        </select>
      </div>


      <!-- Benthos for Align -->
      <div>
        <label class="lb">{{label.sel_benthos_align[0]}}: </label>
        <select v-model="com_benthos_align" class="selector">
          <option value="" disabled selected>{{label.sel_benthos_align[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Benthos').value" :value="cn">{{cn}}</option>
        </select>   
      </div>

      <!-- Benthos for Level -->
      <div>
        <label class="lb">{{label.sel_benthos_level[0]}}: </label>
        <select v-model="com_benthos_level" class="selector">
          <option value="" disabled selected>{{label.sel_benthos_level[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Benthos').value" :value="cn">{{cn}}</option>
        </select>   
      </div>

      <!-- Benthos for Data -->
      <div>
        <label class="lb">{{label.sel_benthos[0]}}: </label>
        <select v-model="com_benthos" class="selector">
          <option value="" disabled selected>{{label.sel_benthos[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Benthos').value" :value="cn">{{cn}}</option>
        </select>   
      </div>          

      <!-- Reader for Align -->
      <div>
        <label class="lb">{{label.sel_reader_align_map[0]}}: </label>
        <select v-model="com_reader_align" class="selector">
          <option value="" disabled selected>{{label.sel_reader_align_map[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Reader').value" :value="cn">{{cn}}</option>
        </select>
      </div>         

      <!-- Reader for Level -->
      <div>
        <label class="lb">{{label.sel_reader_level_map[0]}}: </label>
        <select v-model="com_reader_level" class="selector">
          <option value="" disabled selected>{{label.sel_reader_level_map[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Reader').value" :value="cn">{{cn}}</option>
        </select>
      </div>
      

      <!-- Readers -->
      <div v-for="(rd, j) in nReader" >
        <label class="lb">{{label.sel_reader[0]}}: </label>
        <select v-model="com_reader[j]" class="selector">
          <option value="" disabled selected>{{label.sel_reader[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Reader').value" :value="cn">{{cn}}</option>
        </select>
        <input v-if="j==nReader.length-1" type="button" value="remove" :disabled="nReader.length==1" @click="btn_remove_reader()">
        <input v-if="j==nReader.length-1" type="button" value="more" @click="btn_add_reader()">        
      </div>
      

      <!-- TxtClassifier -->
      <div>
        <label class="lb">{{label.sel_txtclassifier[0]}}: </label>
        <select v-model="com_txtclassifier" class="selector">
          <option value="" disabled selected>{{label.sel_txtclassifier[1]}}</option>
          <option v-for="(cn, i) in mPN.get('TxtClassifier').value" :value="cn">{{cn}}</option>
        </select>
      </div>      


      <!-- Level -->
      <div>
        <label class="lb">{{label.sel_level[0]}}: </label>
        <select v-model="com_level" class="selector">
          <option value="" disabled selected>{{label.sel_level[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Level').value" :value="cn">{{cn}}</option>
        </select>
      </div>
      

      <!-- Weight -->
      <div>
        <label class="lb">{{label.sel_weight[0]}}: </label>
        <select v-model="com_weight" class="selector">
          <option value="" disabled selected>{{label.sel_weight[1]}}</option>
          <option v-for="(cn, i) in mPN.get('Weight').value" :value="cn">{{cn}}</option>
        </select>
      </div>

      <input type="button" value="Composite" @click="btn_composite()">  

    </form>
</div>
`;

export function getForm_Hub() {
  return form_hub;
}
