const http = require('../../common/http');

const ClientForum = (baseUrl) => {

    const client = http.ClientForum(baseUrl);

    return {
        listForums: () => client.get('/forums'),
    }

};

module.exports = { ClientForum };
