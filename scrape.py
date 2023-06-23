from recipe_scrapers import scrape_me
import json
import re
all_links = []
used_links = []

with open('foodcom.txt') as file:
    while (line := file.readline().rstrip()):
        scraper = scrape_me(line)
        parsed = scraper.to_json()
        parsed['yields'] = int(re.findall("\d+", parsed['yields'])[0])
        parsed['canonical_url'] = line
        all_links.append(parsed)
        used_links.append(line)
        with open("data/scraped.json", "w") as outfile:
            outfile.write(json.dumps(all_links))
        with open('data/used.txt', mode='wt', encoding='utf-8') as myfile:
            myfile.write('\n'.join(used_links))
        print(line)


