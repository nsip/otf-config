// port 1555 is set in config.toml

const HOST_PORT = `http://127.0.0.1:1555/`

function fetch_get(path) {

    let url = HOST_PORT + path;

    const rest = fetch(url)
        .then(resp => resp.json())
        .then(data => {
            // console.log(data);
            return data;
        });

    //   const data = async () => {
    //     const a = await prom;
    //     console.log(a);
    //     return a;
    //   };

    return rest;
}

export function get_dispense() {
    return fetch_get("dispense")
}

export function get_allgrp() {
    return fetch_get("allgrp")
}

export function get_allitem() {
    return fetch_get("allitems")
}

export function get_cfg(project, cfgname) {
    switch (project) {
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

////////////////////////////////////////////////////////////////////////////////

function fetch_post(path, data) {

    let url = HOST_PORT + path;

    // console.log(data);

    const rest = fetch(url, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', },
        body: JSON.stringify(data),
    })
        .then(resp => resp.json())
        .then(data => {
            console.log(data);
            return data
        })
        .catch(error => console.error('Error:', error));

    return rest;
}

export function post_cfg(project, data) {
    switch (project) {
        case "NatsStreaming":
            return fetch_post(`otf-config/natsstreaming`, data)
        case "Nias3":
            return fetch_post(`otf-config/nias3`, data)
        case "Benthos":
            return fetch_post(`otf-config/benthos`, data)
        case "Reader":
            return fetch_post(`otf-config/reader`, data)
        case "Align":
            return fetch_post(`otf-config/align`, data)
        case "TxtClassifier":
            return fetch_post(`otf-config/textclassifier`, data)
        case "Level":
            return fetch_post(`otf-config/level`, data)
        case "Weight":
            return fetch_post(`otf-config/weight`, data)
        case "Hub":
            return fetch_post(`otf-config/hub`, data)
    }
}

export function post_table(name, data) {
    return fetch_post(`composite?name=${name}`, data);
}

////////////////////////////////////////////////////////////////////////////////

function fetch_put(path, data) {

    let url = HOST_PORT + path;

    // console.log(data);

    const rest = fetch(url, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json', },
        body: JSON.stringify(data),
    })
        .then(resp => resp.json())
        .then(data => {
            console.log(data);
            return data
        })
        .catch(error => console.error('Error:', error));

    return rest;
}

export function put_cfg(project, data) {

    // console.log(data)

    switch (project) {
        case "NatsStreaming":
            return fetch_put(`otf-config/natsstreaming`, data)
        case "Nias3":
            return fetch_put(`otf-config/nias3`, data)
        case "Benthos":
            return fetch_put(`otf-config/benthos`, data)
        case "Reader":
            return fetch_put(`otf-config/reader`, data)
        case "Align":
            return fetch_put(`otf-config/align`, data)
        case "TxtClassifier":
            return fetch_put(`otf-config/textclassifier`, data)
        case "Level":
            return fetch_put(`otf-config/level`, data)
        case "Weight":
            return fetch_put(`otf-config/weight`, data)
        case "Hub":
            return fetch_put(`otf-config/hub`, data)
    }
}

////////////////////////////////////////////////////////////////////////////////

function fetch_delete(path) {

    let url = HOST_PORT + path;
    // console.log(url);

    const rest = fetch(url, {
        method: 'DELETE',
    })
        .then(resp => resp.json())
        .then(data => {
            console.log(data);
            return data
        })
        .catch(error => console.error('Error:', error));

    return rest;
}

export function delete_cfg(project, cfgname) {

    // console.log("in delete_cfg");

    switch (project) {
        case "NatsStreaming":
            return fetch_delete(`otf-config/natsstreaming?cfgName=${cfgname}`)
        case "Nias3":
            return fetch_delete(`otf-config/nias3?cfgName=${cfgname}`)
        case "Benthos":
            return fetch_delete(`otf-config/benthos?cfgName=${cfgname}`)
        case "Reader":
            return fetch_delete(`otf-config/reader?cfgName=${cfgname}`)
        case "Align":
            return fetch_delete(`otf-config/align?cfgName=${cfgname}`)
        case "TxtClassifier":
            return fetch_delete(`otf-config/textclassifier?cfgName=${cfgname}`)
        case "Level":
            return fetch_delete(`otf-config/level?cfgName=${cfgname}`)
        case "Weight":
            return fetch_delete(`otf-config/weight?cfgName=${cfgname}`)
        case "Hub":
            return fetch_delete(`otf-config/hub?cfgName=${cfgname}`)
    }
}