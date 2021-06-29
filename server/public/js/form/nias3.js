let form_nias3 = ` 
<div v-if="vi.nias3">

    <form v-if="selected" class="cfgform">                        
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">        
                       
        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">                       
        
        <input type="button" value="new">
    </form>

    <div v-for="(cn, i) in names_in">
        
        <form class="cfgform">
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">            
                             
            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">            
            
            <input type="button" value="update">
        </form>
    </div>

</div>`;

export function getForm_Nias3() {
  return form_nias3;
}
