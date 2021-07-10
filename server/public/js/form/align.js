let form_align = `
<div v-if="mPV.get('Align').value" v-for="(cn, i) in input">

    <button v-if="i==0" type="button" class="collapsible" @click="collapse(i)">New Config</button> 
    <button v-if="i>0" type="button" class="collapsible" @click="collapse(i)">{{input[i].name}}</button>    

    <form class="cfgform" v-if="vf[i]">

      <label v-if="i==0" class="lb">{{label.name[0]}}: </label>
      <input v-if="i==0" v-model.trim="input[i].name" type="text" :placeholder="label.name[1]" :readonly="i>0">   

      <label class="lb">{{label.path[0]}}:</label>
      <input v-model.trim="input[i].path" type="text" :placeholder="label.path[1]">   

      <label class="lb">{{label.args[0]}}:</label>
      <input v-model.trim="input[i].args" type="text" :placeholder="label.args[1]">  

      <label class="lb">{{label.svrname[0]}}:</label>
      <input v-model.trim="input[i].svrname" type="text" :placeholder="label.svrname[1]">           

      <label class="lb">{{label.svrid[0]}}:</label>
      <input v-model.trim="input[i].svrid" type="text" :placeholder="label.svrid[1]"> 
      
      <label class="lb">{{label.port[0]}}:</label>
      <input v-model.trim="input[i].port" type="text" :placeholder="label.port[1]"> 

      <label class="lb">{{label.niashost[0]}}:</label>
      <input v-model.trim="input[i].niashost" type="text" :placeholder="label.niashost[1]"> 

      <label class="lb">{{label.niasport[0]}}:</label>
      <input v-model.trim="input[i].niasport" type="text" :placeholder="label.niasport[1]"> 

      <label class="lb">{{label.niastoken[0]}}:</label>
      <input v-model.trim="input[i].niastoken" type="text" :placeholder="label.niastoken[1]"> 

      <label class="lb">{{label.tchost[0]}}:</label>
      <input v-model.trim="input[i].tchost" type="text" :placeholder="label.tchost[1]"> 

      <label class="lb">{{label.tcport[0]}}:</label>
      <input v-model.trim="input[i].tcport" type="text" :placeholder="label.tcport[1]"> 

      <!--
      <label class="lb">***:</label>
      <input v-model.trim="input[i].***" type="text" placeholder="***">            
      -->
        
      <input v-if="i==0" type="button" value="new" :disabled="disable_btn" @click="btn_new(selproj)">
      <input v-if="i>0" type="button" value="delete" :disabled="disable_btn" @click="btn_delete(selproj, i)">
      <input v-if="i>0" type="button" value="update" :disabled="disable_btn" @click="btn_update(selproj, i)">
      
    </form>
</div>
`;

export function getForm_Align() {
  return form_align;
}
