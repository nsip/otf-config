let form_reader = `
<div v-if="mPV.get('Reader').value" v-for="(cn, i) in input">    

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

        <label class="lb">{{label.svrname[0]}}:</label>
        <input v-model.trim="input[i].svrname" type="text" :placeholder="label.svrname[1]">           

        <label class="lb">{{label.svrid[0]}}:</label>
        <input v-model.trim="input[i].svrid" type="text" :placeholder="label.svrid[1]">  

        <label class="lb">{{label.provider[0]}}:</label>
        <input v-model.trim="input[i].provider" type="text" :placeholder="label.provider[1]">    

        <label class="lb">{{label.inputfmt[0]}}:</label>
        <input v-model.trim="input[i].inputFormat" type="text" :placeholder="label.inputfmt[1]">    

        <label class="lb">{{label.alignmethod[0]}}:</label>
        <input v-model.trim="input[i].alignMethod" type="text" :placeholder="label.alignmethod[1]">    

        <label class="lb">{{label.levelmethod[0]}}:</label>
        <input v-model.trim="input[i].levelMethod" type="text" :placeholder="label.levelmethod[1]">    

        <label class="lb">{{label.gencapability[0]}}:</label>
        <input v-model.trim="input[i].capability" type="text" :placeholder="label.gencapability[1]"> 

        <label class="lb">{{label.natshost[0]}}:</label>
        <input v-model.trim="input[i].natsHost" type="text" :placeholder="label.natshost[1]">

        <label class="lb">{{label.natsport[0]}}:</label>
        <input v-model.trim="input[i].natsPort" type="text" :placeholder="label.natsport[1]"> 

        <label class="lb">{{label.natscluster[0]}}:</label>
        <input v-model.trim="input[i].natsCluster" type="text" :placeholder="label.natscluster[1]"> 
        
        <label class="lb">{{label.topic[0]}}:</label>
        <input v-model.trim="input[i].topic" type="text" :placeholder="label.topic[1]"> 

        <label class="lb">{{label.folder[0]}}:</label>
        <input v-model.trim="input[i].folder" type="text" :placeholder="label.folder[1]"> 
        
        <label class="lb">{{label.filesuffix[0]}}:</label>
        <input v-model.trim="input[i].suffix" type="text" :placeholder="label.filesuffix[1]">

        <label class="lb">{{label.interval[0]}}:</label>
        <input v-model.trim="input[i].interval" type="text" :placeholder="label.interval[1]">
                
        <label class="lb">{{label.recursive[0]}}:</label>
        <input v-model.trim="input[i].recursive" type="text" :placeholder="label.recursive[1]">
                
        <label class="lb">{{label.dotfiles[0]}}</label>
        <input v-model.trim="input[i].dotfiles" type="text" :placeholder="label.dotfiles[1]"> 
        
        <label class="lb">{{label.ignore[0]}}:</label>
        <input v-model.trim="input[i].ignore" type="text" :placeholder="label.ignore[1]"> 

        <label class="lb">{{label.concurrfiles[0]}}:</label>
        <input v-model.trim="input[i].concurrFiles" type="text" :placeholder="label.concurrfiles[1]"> 

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

export function getForm_Reader() {
    return form_reader;
}
