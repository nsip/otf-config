let form_weight = `
<div v-if="mPV.get('Weight').value" v-for="(cn, i) in input">    

    <button v-if="i==0" type="button" class="collapsible" @click="collapse(i)">create a new config</button> 
    <button v-if="i>0" type="button" class="collapsible" @click="collapse(i)">{{input[i].name}}</button>   

    <form class="cfgform" v-if="vf[i]">

      <label v-if="i==0" class="lb">{{label.name[0]}}:</label>
      <input v-if="i==0" v-model.trim="input[i].name" type="text" :placeholder="label.name[1]" :readonly="i>0">   

      <label class="lb">{{label.path[0]}}:</label>
      <input v-model.trim="input[i].path" type="text" :placeholder="label.path[1]">   

      <label class="lb">{{label.args[0]}}:</label>
      <input v-model.trim="input[i].args" type="text" :placeholder="label.args[1]">  

      <label class="lb">{{label.delay[0]}}:</label>
      <input v-model.trim="input[i].delay" type="text" :placeholder="label.delay[1]"> 
      
      <label class="lb">{{label.in[0]}}:</label>
      <input v-model.trim="input[i].in" type="text" :placeholder="label.in[1]"> 

      <label class="lb">{{label.out[0]}}:</label>
      <input v-model.trim="input[i].out" type="text" :placeholder="label.out[1]"> 

      <label class="lb">{{label.svrname[0]}}:</label>
      <input v-model.trim="input[i].service.svrname" type="text" :placeholder="label.svrname[1]">

      <label class="lb">{{label.svrid[0]}}:</label>
      <input v-model.trim="input[i].service.svrid" type="text" :placeholder="label.svrid[1]"> 
            
      <label class="lb">{{label.port[0]}}:</label>
      <input v-model.trim="input[i].service.port" type="text" :placeholder="label.port[1]"> 

      <label class="lb">{{label.api[0]}}:</label>
      <input v-model.trim="input[i].service.api" type="text" :placeholder="label.api[1]"> 

      <label class="lb">{{label.wt_studentIDPath[0]}}:</label>
      <input v-model.trim="input[i].weighting.studentIDPath" type="text" :placeholder="label.wt_studentIDPath[1]"> 

      <label class="lb">{{label.wt_progressionLevelPath[0]}}:</label>
      <input v-model.trim="input[i].weighting.progressionLevelPath" type="text" :placeholder="label.wt_progressionLevelPath[1]"> 

      <label class="lb">{{label.wt_timePath0[0]}}:</label>
      <input v-model.trim="input[i].weighting.timePath0" type="text" :placeholder="label.wt_timePath0[1]"> 

      <label class="lb">{{label.wt_timePath1[0]}}:</label>
      <input v-model.trim="input[i].weighting.timePath1" type="text" :placeholder="label.wt_timePath1[1]"> 

      <label class="lb">{{label.wt_scorePath[0]}}:</label>
      <input v-model.trim="input[i].weighting.scorePath" type="text" :placeholder="label.wt_scorePath[1]"> 

      <!--
      <label class="lb">***:</label>
      <input v-model.trim="input[i].***" type="text" :placeholder="***">            
      -->
        
      <input v-if="i==0" type="button" value="new" @click="btn_new(selproj)">
      <input v-if="i>0" type="button" value="delete" @click="btn_delete(selproj, i)">
      <input v-if="i>0" type="button" value="update" @click="btn_update(selproj, i)">
      
    </form>
</div>
`;

export function getForm_Weight() {
  return form_weight;
}
