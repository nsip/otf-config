import { getEmitter } from './js/mitt.js'
import { get_allitem, get_cfg } from './js/fetch.js'
import { getForm } from './js/templates.js'

const emitter = getEmitter();

function viform(proj, vi) {

    vi.name = false;
    vi.path = false;
    vi.svrname = false;
    vi.svrid = false;
    vi.provider = false;
    vi.inputfmt = false;

    switch (proj) {
        case "NatsStreaming":
            vi.name = true;
            vi.path = true;
            break;
        default:
            vi.name = true;
            vi.path = true;
            vi.svrname = true;
            vi.svrid = true;
            vi.provider = true;
            vi.inputfmt = true;
    }
}

export default {

    setup() {

        let selected = Vue.ref(false);
        let title = Vue.ref("OTF project (select from left)");

        let new_cfg = Vue.reactive({
            name: "",
            path: "",
            svrname: "",
            svrid: "",
            provider: "",
            inputfmt: "",
        });

        let vi = Vue.reactive({
            name: true,
            path: true,
            svrname: true,
            svrid: true,
            provider: true,
            inputfmt: true,
        })

        // let in_cfg = Vue.reactive({
        //     name: Vue.ref([]),
        //     path: Vue.ref([]),
        //     svrname: Vue.ref([]),
        //     svrid: Vue.ref([]),
        //     provider: Vue.ref([]),
        //     inputfmt: Vue.ref([]),
        // })

        let names_in = Vue.ref([]);
        let paths_in = Vue.ref([]);
        let svrnames_in = Vue.ref([]);
        let svrids_in = Vue.ref([]);
        let providers_in = Vue.ref([]);
        let inputfmts_in = Vue.ref([]);

        // listen to an event
        emitter.on('selected', e => {

            // test
            console.log('forms received:', e);

            // selected
            selected.value = true;

            // change title
            title.value = `OTF - ${e}`;

            // set visibility of each project config
            viform(e, vi);

            // fetch all selected config            
            (async () => {
                let arg = `${e}s`;
                if (e == 'Benthos') {
                    arg = `${e}es`;
                }
                const all = await get_allitem();

                // clear ***_in for inflating again
                names_in.value = [];
                paths_in.value = [];
                svrnames_in.value = [];
                svrids_in.value = [];
                providers_in.value = [];
                inputfmts_in.value = [];

                // in_cfg.name.value = [];

                // console.log(a);
                // console.log(a[arg]);               

                // fetch config content array
                all[arg].forEach(cname => {
                    (async () => {
                        const b = await get_cfg(e, cname);
                        console.log(b);

                        names_in.value.push(b.name);
                        paths_in.value.push(b.path);
                        svrnames_in.value.push(b.svrname);
                        svrids_in.value.push(b.svrid);
                        providers_in.value.push(b.provider);
                        inputfmts_in.value.push(b.inputFormat);

                        // in_cfg.name.value.push(b.name);

                    })();
                })

            })();
        });

        return {
            selected,
            title,
            new_cfg,
            vi,
            // in_cfg,

            names_in,
            paths_in,
            svrnames_in,
            svrids_in,
            providers_in,
            inputfmts_in,
        };
    },

    template: getForm(),
}