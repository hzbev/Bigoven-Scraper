const fs = require('fs');
const he = require('he');
const axios = require('axios');
let pattern = /https:\/\/foodnetwork.co.uk\/recipes\/\S+/;
let links = JSON.parse(fs.readFileSync("a.json", "utf8"));
// fs.writeFileSync('a.json',JSON.stringify(links))
let ALL_LINKS = []
let word = "foodcom";

const timer = ms => new Promise(res => setTimeout(res, ms))





    // for (let x = 1; x <= 4500; x++) {
// 
async function start() {
    for (let x = 1; x <= 2; x++) {
        let res = await axios({method: 'get',
        url: `https://api.food.com/services/mobile/fdc/search/sectionfront?pn=${x}&recordType=Recipe&sortBy=mostPopular&collectionId=17`,
      })
        let data = res.data.response.results;
        for (let i = 0; i < data.length; i++) {
            ALL_LINKS.push(data[i].record_url)
        }
        console.log("len", ALL_LINKS.length, "index", x)
        fs.writeFileSync(`${word}.txt`, ALL_LINKS.join('\n'))
        await timer(500)
    }
}

start()






