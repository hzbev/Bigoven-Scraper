const fs = require('fs');
const he = require('he');
const axios = require('axios');
let pattern = /https:\/\/foodnetwork.co.uk\/recipes\/\S+/;
let links = JSON.parse(fs.readFileSync("a.json", "utf8"));
// fs.writeFileSync('a.json',JSON.stringify(links))
let ALL_LINKS = []
let word = "snack";

const timer = ms => new Promise(res => setTimeout(res, ms))




async function start() {
    for (let x = 1; x <= 68; x++) {
        let res = await axios({
            method: 'get',
            url: `https://foodnetwork.co.uk/collections/${word}-recipes?recipes=${x}`,
            headers: { 
                'authority': 'foodnetwork.co.uk', 
                'accept': 'text/html, application/xhtml+xml', 
                'accept-language': 'en-GB,en;q=0.9,sk-GB;q=0.8,sk;q=0.7,de-DE;q=0.6,de;q=0.5,en-US;q=0.4,cs;q=0.3', 
                'content-type': 'application/json', 
                'cookie': 'OptanonAlertBoxClosed=2023-01-24T15:03:47.626Z; eupubconsent-v2=CPmH4EzPmH4EzAcABBENC0CsAP_AAH_AACiQIDtf_X__b2_j-_5_f_t0eY1P9_7__-0zjhfdl-8N3f_X_L8X52M7vF36pq4KuR4ku3LBIQdlHOHcTUmw6okVryOsbk2cr7NKJ7PEmnMbOydYGH9_n13T-ZKY7__vf_7zvneAA75qAGAEAIASQggAgAAAAAIAAAMABAAAABIAAAAAgAgciASYatxAF2JY4E20YRQIgRhWEh1AoAKKAYWiAwgdXBTsrgJ9YQIAEAoAjAiBDgCjBgEAAAEASERASBHggEABEAgABAAqEQgAI2AQUAFgYBAAKAaFiBFAEIEhBkQERSmBAQAAFAIAQAAAAAAiAAAAAAAAAgAhAAIAMgAA.f_gAD_gAAAAA; permutive-id=7fc210e7-5aec-4708-9583-193af0761ab3; __utmc=42912946; __utmz=42912946.1674572628.1.1.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __qca=P0-11849531-1674572627901; __utma=42912946.822777415.1674572628.1674572628.1674575395.2; OptanonConsent=isIABGlobal=false&datestamp=Tue+Jan+24+2023+16%3A57%3A17+GMT%2B0100+(Central+European+Standard+Time)&version=6.20.0&hosts=&consentId=644ae618-8508-42bf-980b-390b27d988f0&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CBG1550%3A1%2CC0002%3A1%2CC0003%3A1%2CBG1551%3A1%2CC0004%3A1%2CBG1552%3A1&geolocation=SK%3BZI&AwaitingReconsent=false', 
                'referer': 'https://foodnetwork.co.uk/collections/lunch-recipes?recipes=32', 
                'sec-ch-ua': '"Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"', 
                'sec-ch-ua-mobile': '?0', 
                'sec-ch-ua-platform': '"macOS"', 
                'sec-fetch-dest': 'empty', 
                'sec-fetch-mode': 'cors', 
                'sec-fetch-site': 'same-origin', 
                'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36', 
                'x-inertia': 'true', 
                'x-inertia-version': '2d7bdd94cbba517752336bcf82da0e92', 
                'x-requested-with': 'XMLHttpRequest'
              }
        })
        let data = res.data.props.resource.all_recipes.paginate.resource.data;
        for (let i = 0; i < data.length; i++) {
            ALL_LINKS.push(data[i].url)
        }
        console.log("len", ALL_LINKS.length, "index", x)
        fs.writeFileSync(`${word}.txt`, ALL_LINKS.join('\n'))
        await timer(5000)
    }
}

start()






