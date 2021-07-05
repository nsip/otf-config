
// Only ONE Hub Form

let form_hub = `
<div v-if="vi.hub">    
    <form class="cfgform">

      <label class="lb">{{label.name[0]}}: </label>
      <input v-model="input[0].name" type="text" :placeholder="label.name[1]">   

      <label class="lb">{{label.path[0]}}: </label>
      <input v-model="input[0].path" type="text" :placeholder="label.path[1]">   

      <input type="button" value="update" :disabled="disable_btn" @click="btn_update(selproj, 0)">

      <p></p>

      <div class="dropdown">
        <button @click="myFunction()" class="dropbtn">Select NatsStreaming</button>
        <div id="myDropdown" class="dropdown-content">
          <a href="#" v-for="(cn, i) in dropcontent" :id="cn" @click="dropdown(cn)">{{cn}}</a>          
        </div>
      </div>      

    </form>
</div>
`;

export function getForm_Hub() {
  return form_hub;
}
