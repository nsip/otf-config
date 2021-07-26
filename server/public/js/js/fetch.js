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

export function get_allgrp() {
    return fetch_get("allgrp")
}

export function get_allitem() {
    return fetch_get("allitems")
}

export function get_cfg(project, cfgname) {
    return fetch_get(`otf-config/?project=${project}&cfgName=${cfgname}`)
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
    return fetch_post(`otf-config/?project=${project}`, data)
}

export function post_table(name, exepath, data) {
    return fetch_post(`composite?name=${name}&exepath=${exepath}`, data);
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
    return fetch_put(`otf-config/?project=${project}`, data)
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
    return fetch_delete(`otf-config/?project=${project}&cfgName=${cfgname}`)
}
