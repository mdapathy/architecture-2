const http = require('./http.js');

const Client = baseUrl => {
    const client = http.Client(baseUrl);

    return {
        listForum: () => client.get('/forums'),
        createUser: (name, strings) => client.post('/user', {"name": name, "interests": strings})
    }
};

module.exports = {Client};