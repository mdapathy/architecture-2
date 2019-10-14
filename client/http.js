const request = require('request');

const Client = (baseUrl) => {
    const respHandler = resp => {
        if(resp.ok)
            return resp.json();
        throw new Error(`Unexpected server responce ${resp.status} ${resp.statusText}`);
    }

    return {
        get: path => {
            return new Promise((res, rej) => {
                request(`${baseUrl}${path}`, {json: true}, (err, res, body) => {
                    if(err) {
                        rej(err);
                        return;
                    }
                    res(body);
                })
            })
        },
        post: async (path, data) => {
            return new Promise((res, rej) => {
                request(baseUrl+path, {json: true, method: 'POST', body: data}, (err, res, body) => {
                    if(err) {
                        rej(err);
                        return;
                    }
                    res(body);
                });
            });
        }
    };
};

module.exports = {Client};