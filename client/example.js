const forums = require('./client');

const client = forums.Client('http://localhost:8080');

// Scenario 1: Display available topics.
client.listForum()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available topics:');
        list.forEach((c) => console.log(c));
    })
    .catch((e) => {
        console.log(`Problem listing available topics: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.createUser("Andrew", ['Ukraine', 'Vietnam'])
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create user response:', resp);
        return client.listForum()
            .then((list) => list.forEach((c) => console.log(c)))
    })
    .catch((e) => {
        console.log(`Problem creating a new User: ${e.message}`);
    });