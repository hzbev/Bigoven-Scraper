package main

import (
	"fmt"
	"scraper/helper"
)

// switch to yummly

func main() {
	parsedLinks := helper.ParseSearch("https://www.bigoven.com/recipes/salad/dessert-and-gelatin/page/5")

	for i := 0; i < len(parsedLinks); i++ {

		parsed := helper.GetData(parsedLinks[i])
		fmt.Println(parsed.RecipeInstructions, parsedLinks[i])
		fmt.Println()

	}

	// fmt.Println(parsed.Image)
}
