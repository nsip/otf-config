
// *** defined name from log http fetch result ***
// fields for VUE v-bind

const _NatsStreaming = {
    args: "",
    delay: "",
    name: "",
    path: "",
}

const _Nias3 = {
    args: "",
    delay: "",
    name: "",
    path: "",
}

const _Benthos = {
    args: "",
    delay: "",
    name: "",
    path: "",
}

const _Reader = {
    alignMethod: "",
    args: "",
    capability: "",
    concurrFiles: 10,
    delay: "",
    dotfiles: false,
    folder: "",
    ignore: "",
    inputFormat: "",
    interval: "",
    levelMethod: "",
    name: "",
    natsCluster: "",
    natsHost: "",
    natsPort: 0,
    path: "",
    provider: "",
    recursive: false,
    suffix: "",
    svrid: "",
    svrname: "",
    topic: "",
}

const _Align = {
    args: "",
    delay: "",
    host: "",
    name: "",
    niasHost: "",
    niasPort: 0,
    niasToken: "",
    path: "",
    port: 0,
    svrid: "",
    svrname: "",
    tcHost: "",
    tcPort: 0,
}

const _TxtClassifier = {
    args: "",
    delay: "",
    name: "",
    path: "",
    port: 0,
}

const _Level = {
    args: "",
    delay: "",
    host: "",
    name: "",
    niasHost: "",
    niasPort: 0,
    niasToken: "",
    path: "",
    port: 0,
    svrid: "",
    svrname: "",
}

const _Weight = {
    args: "",
    delay: "",
    failWhenErr: false,
    name: "",
    path: "",
}

const _Hub = {
    args: "",
    name: "",
    path: "",
    tablename: "",
}

const mInitInput = new Map();
mInitInput.set('NatsStreaming', _NatsStreaming);
mInitInput.set('Nias3', _Nias3);
mInitInput.set('Benthos', _Benthos);
mInitInput.set('Reader', _Reader);
mInitInput.set('Align', _Align);
mInitInput.set('TxtClassifier', _TxtClassifier);
mInitInput.set('Level', _Level);
mInitInput.set('Weight', _Weight);
mInitInput.set('Hub', _Hub);

export function get_init_input(proj) {
    return mInitInput.get(proj);
}