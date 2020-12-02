const fs = require('fs');
const axios = require('axios');
const readline = require('readline');

const BASE_ADDRESS = 'https://digimon-api.vercel.app/api/digimon/name/';

const start = process.hrtime();

const input = fs.createReadStream('digimon.txt');
const reader = readline.createInterface(input);
const promises = [];

reader.on('line', data => {
    // 'transformResponse' suppresses axios' json parsing, which is unnecessary for this challenge
    const promise = axios.get(BASE_ADDRESS + data, { transformResponse: res => res })
        .then(response => response.data);

    promises.push(promise);
});

reader.on('close', () => {
    Promise.all(promises)
        .then(results => {
            const end = process.hrtime(start);
            const ms = end[0] * 1000 + Math.round(end[1] / 1000000);

            results.forEach(r => console.log(r));
            console.log(`Script executed in ${ms} milliseconds.`);
        });
});
