
// includes some variables from caller 'template:'

let form_head = `<div id="mainform">
<h4 style="text-align: center;">{{title}}</h4>`;

let form_tail = `<br></div>`;

let form_natsstreaming = `
<div v-if="vi.natsstreaming">

        <form v-if="selected" class="cfgform">                        
            <label class="lb">name: </label>
            <input v-model="new_cfg.name" type="text" placeholder="config name">
            <br>
                           
            <label class="lb">path to service executable:</label>
            <input v-model="new_cfg.path" type="text" placeholder="path to service executable">
            <br>
                           
            <br>
            <input type="button" value="new">
        </form>

        <div v-for="(cn, i) in names_in">
            <br>
            <form class="cfgform">
                <label class="lb">name: </label>
                <input v-model="names_in[i]" type="text" placeholder="config name">
                <br>
                                 
                <label class="lb">path to service executable:</label>
                <input v-model="paths_in[i]" type="text" placeholder="path to service executable">
                <br>
                
                <br>
                <input type="button" value="update">
            </form>
        </div>

</div>`

let form_nias3 = ` 
<div v-if="vi.nias3">

    <form v-if="selected" class="cfgform">                        
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">
        <br>
                       
        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">
        <br>
                       
        <br>
        <input type="button" value="new">
    </form>

    <div v-for="(cn, i) in names_in">
        <br>
        <form class="cfgform">
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">
            <br>
                             
            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">
            <br>
            
            <br>
            <input type="button" value="update">
        </form>
    </div>

</div>`

let form_benthos = `
<div v-if="vi.benthos">

    <form v-if="selected" class="cfgform">                        
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">
        <br>
                       
        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">
        <br>
                       
        <br>
        <input type="button" value="new">
    </form>

    <div v-for="(cn, i) in names_in">
        <br>
        <form class="cfgform">
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">
            <br>
                             
            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">
            <br>
            
            <br>
            <input type="button" value="update">
        </form>
    </div>

</div>`

let form_reader = `
<div v-if="vi.reader">

    <form v-if="selected" class="cfgform">
        
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">
        <br>

        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">
        <br>

        <label class="lb">service name:</label>
        <input v-model="new_cfg.svrname" type="text" placeholder="service name">
        <br>

        <label class="lb">service id:</label>
        <input v-model="new_cfg.svrid" type="text" placeholder="service id">
        <br>

        <label class="lb">provider:</label>
        <input v-model="new_cfg.provider" type="text" placeholder="provider">
        <br>

        <label class="lb">input format:</label>
        <input v-model="new_cfg.inputfmt" type="text" placeholder="input format">
        <br>

        <label class="lb">align method:</label>
        <input v-model="new_cfg.alignmethod" type="text" placeholder="align method">
        <br>

        <label class="lb">level method:</label>
        <input v-model="new_cfg.levelmthod" type="text" placeholder="level method">
        <br>

        <!-- 
        <label class="lb">***:</label>
        <input v-model="new_cfg.***" type="text" placeholder="***">
        <br>
        -->

        <br>
        <input type="button" value="new">
    </form>

    <div v-for="(cn, i) in names_in">
        <br>
        <form class="cfgform">
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">
            <br>

            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">
            <br>

            <label class="lb">service name:</label>
            <input v-model="svrnames_in[i]" type="text" placeholder="service name">
            <br>

            <label class="lb">service id:</label>
            <input v-model="svrids_in[i]" type="text" placeholder="service id">
            <br>

            <label class="lb">provider:</label>
            <input v-model="providers_in[i]" type="text" placeholder="provider">
            <br>

            <label class="lb">input format:</label>
            <input v-model="inputfmts_in[i]" type="text" placeholder="input format">
            <br>

            <label class="lb">align method:</label>
            <input v-model="alignmethods_in[i]" type="text" placeholder="align method">
            <br>

            <label class="lb">level method:</label>
            <input v-model="levelmethods_in[i]" type="text" placeholder="level method">
            <br>

            <!-- 
            <label class="lb">***:</label>
            <input v-model="***[i]" type="text" placeholder="***">
            <br>
            -->
            
            <br>
            <input type="button" value="update">
        </form>
    </div>

</div>`

let form_align = ``

let form_txtclassifier = ``

let form_level = ``

let form_weight = ``

let form_hub = ``

export function getForm() {
    return form_head +
        form_natsstreaming +
        form_nias3 +
        form_benthos +
        form_reader +
        form_tail;
}