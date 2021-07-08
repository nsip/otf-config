let form_nias3 = ` 
<div v-if="mPV.get('Nias3').value" v-for="(cn, i) in input">       

    <button v-if="i==0" type="button" class="collapsible" @click="collapse(i)">New Config</button> 
    <button v-if="i>0" type="button" class="collapsible" @click="collapse(i)">{{input[i].name}}</button>   

    <form class="cfgform" v-if="vf[i]">

        <label v-if="i==0" class="lb">{{label.name[0]}}:</label>
        <input v-if="i==0" v-model="input[i].name" type="text" :placeholder="label.name[1]" :readonly="i>0">            
                         
        <label class="lb">{{label.path[0]}}:</label>
        <input v-model="input[i].path" type="text" :placeholder="label.path[1]">    
        
        <label class="lb">{{label.args[0]}}:</label>
        <input v-model="input[i].args" type="text" :placeholder="label.args[1]">  
        
        <input v-if="i==0" type="button" value="new" :disabled="disable_btn" @click="btn_new(selproj)">
        <input v-if="i>0" type="button" value="delete" :disabled="disable_btn" @click="btn_delete(selproj, i)">
        <input v-if="i>0" type="button" value="update" :disabled="disable_btn" @click="btn_update(selproj, i)">
        
    </form>
</div>
`;

export function getForm_Nias3() {
    return form_nias3;
}
