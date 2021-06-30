let form_nias3 = ` 
<div v-if="vi.nias3" v-for="(cn, i) in input">        
    <form class="cfgform">
        <label class="lb">name: </label>
        <input v-model="input[i].name" type="text" placeholder="config name">            
                         
        <label class="lb">path to service executable:</label>
        <input v-model="input[i].path" type="text" placeholder="path to service executable">            
        
        <input v-if="i==0" type="button" value="new">
        <input v-if="i>0" type="button" value="update">
    </form>
</div>
`;

export function getForm_Nias3() {
    return form_nias3;
}
