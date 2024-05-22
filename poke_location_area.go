package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	Previous any            `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func (cfg *config) commandLocationsFwd() error {
	var next string
	if cfg.Next != "" {
		next = cfg.Next
	} else {
		next = "https://pokeapi.co/api/v2/location-area/"
	}

	return cfg.commandLocations(next)
}

func (cfg *config) commandLocationsBck() error {
	var prev string
	if cfg.Previous != "" {
		prev = cfg.Previous
	} else {
		return nil
	}

	return cfg.commandLocations(prev)
}

func (cfg *config) commandLocations(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, ioErr := io.ReadAll(resp.Body)
	if ioErr != nil {
		println("An io error occurred")
		return ioErr
	}

	locAreas := LocationAreas{}
	decErr := json.Unmarshal(data, &locAreas)
	if decErr != nil {
		println("A json unmarshal error occurred")
		return decErr
	}

	cfg.Next = locAreas.Next
	if locAreas.Previous == nil {
		cfg.Previous = ""
	} else {
		cfg.Previous = locAreas.Previous.(string)
	}
	for _, loc := range locAreas.Results {
		println(loc.Name)
	}
	return nil
}
