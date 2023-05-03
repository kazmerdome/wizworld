package wizardworldapi

type ListElixirsRequest struct {
	Ingredient string
}

type ElixirResponse struct {
	Id              string               `json:"id"`
	Name            string               `json:"name"`
	Effect          string               `json:"effect"`
	SideEffects     string               `json:"sideEffects"`
	Difficulty      string               `json:"difficulty"`
	Ingredients     []IngredientResponse `json:"ingredients"`
	Characteristics string               `json:"characteristics"`
}

type IngredientResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
