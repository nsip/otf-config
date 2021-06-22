let url = '127.0.0.1:1555/otf-config/reader?cfgName=spa_prescribed_config';

function test(url) {
    return 'why'
}

Vue.createApp({
    data() {
        return {
            message: test(url)
        };
    }
}).mount("#app");

// GET /otf-config/reader