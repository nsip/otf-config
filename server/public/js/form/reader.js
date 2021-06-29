let form_reader = `
<div v-if="vi.reader">

    <form v-if="selected" class="cfgform">
        
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">        

        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">        

        <label class="lb">service name:</label>
        <input v-model="new_cfg.svrname" type="text" placeholder="service name">        

        <label class="lb">service id:</label>
        <input v-model="new_cfg.svrid" type="text" placeholder="service id">        

        <label class="lb">provider:</label>
        <input v-model="new_cfg.provider" type="text" placeholder="provider">        

        <label class="lb">input format:</label>
        <input v-model="new_cfg.inputfmt" type="text" placeholder="input format">        

        <label class="lb">align method:</label>
        <input v-model="new_cfg.alignmethod" type="text" placeholder="align method">        

        <label class="lb">level method:</label>
        <input v-model="new_cfg.levelmthod" type="text" placeholder="level method">   

        <label class="lb">general capacity:</label>
        <input v-model="new_cfg.gencapacity" type="text" placeholder="general capacity">   
                
        <!-- 
        <label class="lb">***:</label>
        <input v-model="new_cfg.***" type="text" placeholder="***">        
        -->
        
        <input type="button" value="new">
    </form>

    <div v-for="(cn, i) in names_in">
        
        <form class="cfgform">
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">            

            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">            

            <label class="lb">service name:</label>
            <input v-model="svrnames_in[i]" type="text" placeholder="service name">            

            <label class="lb">service id:</label>
            <input v-model="svrids_in[i]" type="text" placeholder="service id">            

            <label class="lb">provider:</label>
            <input v-model="providers_in[i]" type="text" placeholder="provider">            

            <label class="lb">input format:</label>
            <input v-model="inputfmts_in[i]" type="text" placeholder="input format">            

            <label class="lb">align method:</label>
            <input v-model="alignmethods_in[i]" type="text" placeholder="align method">            

            <label class="lb">level method:</label>
            <input v-model="levelmethods_in[i]" type="text" placeholder="level method"> 
            
            <label class="lb">general capacity:</label>
            <input v-model="gencapabilities_in[i]" type="text" placeholder="general capacity"> 
                        
            <!-- 
            <label class="lb">***:</label>
            <input v-model="***[i]" type="text" placeholder="***">            
            -->            
            
            <input type="button" value="update">
        </form>
    </div>

</div>`;

export function getForm_Reader() {
  return form_reader;
}
