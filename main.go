package main

import (
	"encoding/json"
	"fmt"
	"scraper/helper"
	"strings"

	"github.com/imroc/req/v3"
)

func main() {
	parsedData := helper.BigOvenApi{}
	res, _ := req.Get("https://www.bigoven.com/recipe/frosted-strawberry-salad/502652")
	res_string, _ := res.ToString()
	first_index := strings.Index(res_string, `<script type="application/ld+json">`) + 45
	second_index := strings.Index(res_string[first_index:], `</script>`) - 1
	json.Unmarshal([]byte(res_string[first_index:first_index+second_index]), &parsedData)
	fmt.Println(parsedData.CookTime)
	// fmt.Println(second_index)

}
