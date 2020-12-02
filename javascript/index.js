const fs = require('fs');
const axios = require('axios');
const readline = require('readline');

const BASE_ADDRESS = 'https://digimon-api.vercel.app/api/digimon/name/';

const start = process.hrtime();

const input = fs.createReadStream('digimon.txt');
const reader = readline.createInterface(input);
const promises = [];

reader.on('line', data => {
    const promise = axios.get(BASE_ADDRESS + data, { transformResponse: res => res })
        .then(response => console.log(response.data));

    promises.push(promise);
});

reader.on('close', () => {
    Promise.all(promises)
        .then(() => {
            const end = process.hrtime(start);
            console.log(`Script executed in ${end} seconds.`);
        });
});
