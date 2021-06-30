let form_benthos = `
<div v-if="vi.benthos" v-for="(cn, i) in input">
    <form class="cfgform">

        <label class="lb">{{label.name[0]}}: </label>
        <input v-model="input[i].name" type="text" :placeholder="label.name[1]">            
                         
        <label class="lb">{{label.path[0]}}:</label>
        <input v-model="input[i].path" type="text" :placeholder="label.path[1]">            
        
        <input v-if="i==0" type="button" value="new">
        <input v-if="i>0" type="button" value="update">
        
    </form>
</div>
`;

export function getForm_Benthos() {
    return form_benthos;
}
