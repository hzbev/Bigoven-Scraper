package helper

import (
	"encoding/json"
	"strings"

	"github.com/imroc/req/v3"
)

type BigOvenApi struct {
	Context string `json:"@context"`
	Type    string `json:"@type"`
	Name    string `json:"name"`
	Author  struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
	DatePublished  string   `json:"datePublished"`
	Image          []string `json:"image"`
	Description    string   `json:"description"`
	PrepTime       string   `json:"prepTime"`
	CookTime       string   `json:"cookTime"`
	TotalTime      string   `json:"totalTime"`
	RecipeCategory string   `json:"recipeCategory"`
	RecipeCuisine  string   `json:"recipeCuisine"`
	RecipeYield    string   `json:"recipeYield"`
	Keywords       string   `json:"keywords"`
	Nutrition      struct {
		Type                string `json:"@type"`
		Calories            string `json:"calories"`
		FatContent          string `json:"fatContent"`
		CarbohydrateContent string `json:"carbohydrateContent"`
		CholesterolContent  string `json:"cholesterolContent"`
		FiberContent        string `json:"fiberContent"`
		ProteinContent      string `json:"proteinContent"`
		SaturatedFatContent string `json:"saturatedFatContent"`
		ServingSize         string `json:"servingSize"`
		SodiumContent       string `json:"sodiumContent"`
		SugarContent        string `json:"sugarContent"`
		TransFatContent     string `json:"transFatContent"`
	} `json:"nutrition"`
	RecipeIngredient   []string `json:"recipeIngredient"`
	RecipeInstructions []string `json:"recipeInstructions"`
	Video              []struct {
		Type         string   `json:"@type"`
		ContentURL   string   `json:"contentUrl"`
		ThumbnailURL []string `json:"thumbnailUrl"`
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		UploadDate   string   `json:"uploadDate"`
	} `json:"video"`
}

func GetData(link string) BigOvenApi {
	parsedData := BigOvenApi{}
	res, _ := req.Get(link)
	res_string, _ := res.ToString()
	first_index := strings.Index(res_string, `<script type="application/ld+json">`) + 45
	second_index := strings.Index(res_string[first_index:], `</script>`) - 1
	json.Unmarshal([]byte(res_string[first_index:first_index+second_index]), &parsedData)
	return parsedData
}

func ParseSearch(link string) []string {
	var allLinks []string
	res, _ := req.Get(link)
	res_string, _ := res.ToString()

	for i := 0; i < 24; i++ {
		first_index := strings.Index(res_string, `data-url="https://www.bigoven.com/recipe/`) + 10
		second_index := strings.Index(res_string[first_index:], `"><div class="`)
		allLinks = append(allLinks, res_string[first_index:first_index+second_index])
		res_string = res_string[first_index+second_index:]
	}
	return allLinks
}
