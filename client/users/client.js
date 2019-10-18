const http = require('../common/http');

const ClientUser = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        createUser: (name, strings) => client.post('/user', {"name": name, "interests": strings})
    }

};

module.exports = { ClientUser };
