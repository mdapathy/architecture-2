const forums = require('./client');

const client = forums.Client('http://localhost:8080');

// Scenario 1: Display available channels.
client.listForum()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available channels:');
        list.forEach((c) => console.log(c.name));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.createUser(1234, 'Ukraine')
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create user response:', resp);
        return client.listForums()
            .then((list) => list.map((c) => c.name).join(', '))
            .then((str) => {
                console.log(`Forum channels: ${str}`);
            })
    })
    .catch((e) => {
        console.log(`Problem creating a new User: ${e.message}`);
    });