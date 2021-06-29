let form_natsstreaming = `
<div v-if="vi.natsstreaming">

        <div v-for="(cn, i) in names_in">            
            <form class="cfgform">
                <label class="lb">name: </label>
                <input v-model="names_in[i]" type="text" placeholder="config name">                
                                 
                <label class="lb">path to service executable:</label>
                <input v-model="paths_in[i]" type="text" placeholder="path to service executable">                
                               
                <input v-if="i==0" type="button" value="new">
                <input v-if="i>0" type="button" value="update">
            </form>
        </div>

</div>`;

export function getForm_NatsStreaming() {
  return form_natsstreaming;
}
