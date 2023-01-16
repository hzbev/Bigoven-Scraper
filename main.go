package main

import (
	"fmt"
	"scraper/helper"
)

// switch to yummly

func main() {
	parsed := helper.GetData("https://www.bigoven.com/recipe/frosted-strawberry-salad/502652")
	fmt.Println(parsed.Description)
}
