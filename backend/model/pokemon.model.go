package model

type Pokemon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Sprite string `json:"sprite"`
}

type PokemonResponse struct {
	Count   int       `json:"count"` // Total number of items
	Results []Pokemon `json:"results"`
}

type PokemonDetailAPI struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
}

type PokemonDetail struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Sprite string `json:"sprite"`
	Types  []struct {
		Name string `json:"name"`
	} `json:"types"`
	Moves []struct {
		Name string `json:"name"`
	} `json:"moves"`
}
