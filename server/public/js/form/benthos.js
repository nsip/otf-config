let form_benthos = `
<div v-if="vi.benthos" v-for="(cn, i) in input">
    <form class="cfgform">

        <label class="lb">{{label.name[0]}}: </label>
        <input v-model="input[i].name" type="text" :placeholder="label.name[1]" :readonly="i>0">            
                         
        <label class="lb">{{label.path[0]}}:</label>
        <input v-model="input[i].path" type="text" :placeholder="label.path[1]">            
        
        <input v-if="i==0" type="button" value="new" :disabled="disable_btn" @click="btn_new(selproj)">
        <input v-if="i>0" type="button" value="delete" :disabled="disable_btn" @click="btn_delete(selproj, i)">
        <input v-if="i>0" type="button" value="update" :disabled="disable_btn" @click="btn_update(selproj, i)">
        
    </form>
</div>
`;

export function getForm_Benthos() {
    return form_benthos;
}
