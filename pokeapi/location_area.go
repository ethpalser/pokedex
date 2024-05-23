package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (cfg *Config) CommandLocationsFwd(args ...string) error {
	var next string
	if cfg.Next != "" {
		next = cfg.Next
	} else {
		next = "https://pokeapi.co/api/v2/location-area/"
	}

	return cfg.commandLocations(next)
}

func (cfg *Config) CommandLocationsBck(args ...string) error {
	var prev string
	if cfg.Previous != "" {
		prev = cfg.Previous
	} else {
		return nil
	}

	return cfg.commandLocations(prev)
}

func (cfg *Config) commandLocations(url string) error {
	data, getErr := cfg.ApiGet(url)
	if getErr != nil {
		return getErr
	}

	locAreas := LocationAreas{}
	decErr := json.Unmarshal(data, &locAreas)
	if decErr != nil {
		println("Error unmarshaling json")
		return decErr
	}

	cfg.Next = locAreas.Next
	if locAreas.Previous == nil {
		cfg.Previous = ""
	} else {
		cfg.Previous = *locAreas.Previous
	}
	for _, loc := range locAreas.Results {
		println(loc.Name)
	}
	return nil
}

func (cfg *Config) CommandExplore(args ...string) error {
	if len(args) < 1 || len(args) > 1 {
		return &CommandError{message: "An area on a map must be provided"}
	}
	area := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + area

	data, getErr := cfg.ApiGet(url)
	if getErr != nil {
		return getErr
	}
	loc := LocationArea{}
	jsonErr := json.Unmarshal(data, &loc)
	if jsonErr != nil {
		println("Error unmarshaling json")
		return jsonErr
	}

	fmt.Printf("Exploring %v...\n", area)
	if len(loc.PokemonEncounters) > 0 {
		println("Found Pokemon:")
		for _, pkmn := range loc.PokemonEncounters {
			println(" - " + pkmn.Pokemon.Name)
		}
	} else {
		println("No Pokemon found")
	}
	return nil
}
