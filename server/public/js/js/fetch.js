// port 1555 is set in config.toml

// const allGrp = [
//     "NatsStreaming",
//     "Nias3",
//     "Benthos",
//     "Reader",
//     "Align",
//     "TxtClassifier",
//     "Level",
//     "Weight",
//     "Hub",
//   ];

function fetch_get(path) {

    let url = `http://127.0.0.1:1555/` + path;

    const data = fetch(url)
        .then((resp) => resp.json())
        .then((data) => {
            // console.log(data);
            return data;
        });

    //   const data = async () => {
    //     const a = await prom;
    //     console.log(a);
    //     return a;
    //   };

    return data;
}


export function get_allgrp() {
    return fetch_get("allgrp")
}

export function get_allitem() {
    return fetch_get("allitems")
}

export function get_cfg(prject, cfgname) {
    switch (prject) {

        case "NatsStreaming":
            return fetch_get(`otf-config/natsstreaming?cfgName=${cfgname}`)
        case "Nias3":
            return fetch_get(`otf-config/nias3?cfgName=${cfgname}`)
        case "Benthos":
            return fetch_get(`otf-config/benthos?cfgName=${cfgname}`)
        case "Reader":
            return fetch_get(`otf-config/reader?cfgName=${cfgname}`)
        case "Align":
            return fetch_get(`otf-config/align?cfgName=${cfgname}`)
        case "TxtClassifier":
            return fetch_get(`otf-config/textclassifier?cfgName=${cfgname}`)
        case "Level":
            return fetch_get(`otf-config/level?cfgName=${cfgname}`)
        case "Weight":
            return fetch_get(`otf-config/weight?cfgName=${cfgname}`)
        case "Hub":
            return fetch_get(`otf-config/hub?cfgName=${cfgname}`)
    }
}

