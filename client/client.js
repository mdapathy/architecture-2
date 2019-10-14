const http = require('./http.js');

const Client = baseUrl => {
    const client = http.Client(baseUrl);

    return {
        listForum: () => client.get('/forums'),
        createUser: (id, strings) => client.get('/users', {id, strings})
    }
};

module.exports = {Client};