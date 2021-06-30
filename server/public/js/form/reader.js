let form_reader = `
<div v-if="vi.reader" v-for="(cn, i) in input">    
    <form class="cfgform">

        <label class="lb">{{label.name[0]}}:</label>
        <input v-model="input[i].name" type="text" :placeholder="label.name[1]">   

        <label class="lb">{{label.path[0]}}:</label>
        <input v-model="input[i].path" type="text" :placeholder="label.path[1]">   

        <label class="lb">{{label.svrname[0]}}:</label>
        <input v-model="input[i].svrname" type="text" :placeholder="label.svrname[1]">           

        <label class="lb">{{label.svrid[0]}}:</label>
        <input v-model="input[i].svrid" type="text" :placeholder="label.svrid[1]">  

        <label class="lb">{{label.provider[0]}}:</label>
        <input v-model="input[i].provider" type="text" :placeholder="label.provider[1]">    

        <label class="lb">{{label.inputfmt[0]}}:</label>
        <input v-model="input[i].inputfmt" type="text" :placeholder="label.inputfmt[1]">    

        <label class="lb">{{label.alignmethod[0]}}:</label>
        <input v-model="input[i].alignmethod" type="text" :placeholder="label.alignmethod[1]">    

        <label class="lb">{{label.levelmethod[0]}}:</label>
        <input v-model="input[i].levelmethod" type="text" :placeholder="label.levelmethod[1]">    

        <label class="lb">{{label.gencapability[0]}}:</label>
        <input v-model="input[i].gencapability" type="text" :placeholder="label.gencapability[1]"> 

        <label class="lb">{{label.natshost[0]}}:</label>
        <input v-model="input[i].natshost" type="text" :placeholder="label.natshost[1]">

        <label class="lb">{{label.natsport[0]}}:</label>
        <input v-model="input[i].natsport" type="text" :placeholder="label.natsport[1]"> 

        <label class="lb">{{label.natscluster[0]}}:</label>
        <input v-model="input[i].natscluster" type="text" :placeholder="label.natscluster[1]"> 
        
        <label class="lb">{{label.topic[0]}}:</label>
        <input v-model="input[i].topic" type="text" :placeholder="label.topic[1]"> 

        <label class="lb">{{label.folder[0]}}:</label>
        <input v-model="input[i].folder" type="text" :placeholder="label.folder[1]"> 
        
        <label class="lb">{{label.filesuffix[0]}}:</label>
        <input v-model="input[i].filesuffix" type="text" :placeholder="label.filesuffix[1]">

        <label class="lb">{{label.interval[0]}}:</label>
        <input v-model="input[i].interval" type="text" :placeholder="label.interval[1]">
                
        <label class="lb">{{label.recursive[0]}}:</label>
        <input v-model="input[i].recursive" type="text" :placeholder="label.recursive[1]">
                
        <label class="lb">{{label.dotfiles[0]}}</label>
        <input v-model="input[i].dotfiles" type="text" :placeholder="label.dotfiles[1]"> 
        
        <label class="lb">{{label.ignore[0]}}:</label>
        <input v-model="input[i].ignore" type="text" :placeholder="label.ignore[1]"> 

        <label class="lb">{{label.concurrfiles[0]}}:</label>
        <input v-model="input[i].concurrfiles" type="text" :placeholder="label.concurrfiles[1]"> 

        <!--
        <label class="lb">***:</label>
        <input v-model="input[i].***" type="text" :placeholder="***">            
        -->
        
        <input v-if="i==0" type="button" value="new">
        <input v-if="i>0" type="button" value="update">
        
    </form>
</div>
`;

export function getForm_Reader() {
    return form_reader;
}
