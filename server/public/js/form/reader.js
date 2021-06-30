let form_reader = `
<div v-if="vi.reader" v-for="(cn, i) in input">    
    <form class="cfgform">
        <label class="lb">name: </label>
        <input v-model="input[i].name" type="text" placeholder="config name">   

        <label class="lb">path to service executable:</label>
        <input v-model="input[i].path" type="text" placeholder="path to service executable">   

        <label class="lb">service name:</label>
        <input v-model="input[i].svrname" type="text" placeholder="service name">           

        <label class="lb">service id:</label>
        <input v-model="input[i].svrid" type="text" placeholder="service id">  

        <label class="lb">provider:</label>
        <input v-model="input[i].provider" type="text" placeholder="provider">    

        <label class="lb">input format:</label>
        <input v-model="input[i].inputfmt" type="text" placeholder="input format">    

        <label class="lb">align method:</label>
        <input v-model="input[i].alignmethod" type="text" placeholder="align method">    

        <label class="lb">level method:</label>
        <input v-model="input[i].levelmethod" type="text" placeholder="level method">    

        <label class="lb">general capacity:</label>
        <input v-model="input[i].gencapability" type="text" placeholder="general capacity"> 

        <label class="lb">nats host:</label>
        <input v-model="input[i].natshost" type="text" placeholder="nats host">

        <label class="lb">nats port:</label>
        <input v-model="input[i].natsport" type="text" placeholder="nats port"> 
                    
        <!--
        <label class="lb">***:</label>
        <input v-model="input[i].***" type="text" placeholder="***">            
        -->
        
        <input v-if="i==0" type="button" value="new">
        <input v-if="i>0" type="button" value="update">
    </form>
</div>
`;

export function getForm_Reader() {
    return form_reader;
}
