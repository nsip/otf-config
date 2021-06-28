
// includes some variables from caller 'template:'

let formAll = `
<div id="mainform">
<h4 style="text-align: center;">{{title}}</h4>
<form v-if="selected" class="cfgform">   
    <div v-if="vi.name">     
        <label class="lb">name: </label>
        <input v-model="new_cfg.name" type="text" placeholder="config name">
        <br>
    </div>
    <div v-if="vi.path">
        <label class="lb">path to service executable:</label>
        <input v-model="new_cfg.path" type="text" placeholder="path to service executable">
        <br>
    </div>
    <div v-if="vi.svrname">
        <label class="lb">service name:</label>
        <input v-model="new_cfg.svrname" type="text" placeholder="service name">
        <br>
    </div>
    <div v-if="vi.svrid">
        <label class="lb">service id:</label>
        <input v-model="new_cfg.svrid" type="text" placeholder="service id">
        <br>
    </div>
    <div v-if="vi.provider">
        <label class="lb">provider:</label>
        <input v-model="new_cfg.provider" type="text" placeholder="provider">
        <br>
    </div>
    <div v-if="vi.inputfmt">
        <label class="lb">input format:</label>
        <input v-model="new_cfg.inputfmt" type="text" placeholder="input format">
        <br>
    </div>
    <br>
    <input type="button" value="new">
</form>
<div v-for="(cn, i) in names_in">
    <br>
    <form class="cfgform">   
        <div v-if="vi.name">     
            <label class="lb">name: </label>
            <input v-model="names_in[i]" type="text" placeholder="config name">
            <br>
        </div>
        <div v-if="vi.path"> 
            <label class="lb">path to service executable:</label>
            <input v-model="paths_in[i]" type="text" placeholder="path to service executable">
            <br>
        </div>        
        <div v-if="vi.svrname"> 
            <label class="lb">service name:</label>
            <input v-model="svrnames_in[i]" type="text" placeholder="service name">
            <br>
        </div>
        <div v-if="vi.svrid"> 
            <label class="lb">service id:</label>
            <input v-model="svrids_in[i]" type="text" placeholder="service id">
            <br>
        </div>
        <div v-if="vi.provider"> 
            <label class="lb">provider:</label>
            <input v-model="providers_in[i]" type="text" placeholder="provider">
            <br>
        </div>
        <div v-if="vi.inputfmt"> 
            <label class="lb">input format:</label>
            <input v-model="inputfmts_in[i]" type="text" placeholder="input format">
            <br>
        </div>
        <br>
        <input type="button" value="update">
    </form>
</div> 
<br> 
</div>`

export function getForm() {
    return formAll;
}