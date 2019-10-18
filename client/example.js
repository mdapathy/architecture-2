const forums = require('./forums/client.js');
const user = require('./users/client.js');

const clientForum = forums.ClientForum('http://localhost:8080');
const clientUser = user.ClientUser('http://localhost:8080');

// Scenario 1: Display available topics.
clientForum.listForums()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available topics:');
        list.forEach((c) => console.log(c));
    })
    .catch((e) => {
        console.log(`Problem listing available topics: ${e.message}`);
    });

// Scenario 2: Create new channel.
clientUser.createUser("Maria", ['tv-series'])
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create user response:', resp);
        return clientForum.listForums()
            .then((list) => list.forEach((c) => console.log(c)))
    })
    .catch((e) => {
        console.log(`Problem creating a new User: ${e.message}`);
    });