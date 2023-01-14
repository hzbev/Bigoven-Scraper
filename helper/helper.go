package helper

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
