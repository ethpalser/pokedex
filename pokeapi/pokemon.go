package pokeapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
}

func (cfg *Config) CommandCatch(args ...string) error {
	if len(args) < 1 || len(args) > 1 {
		return &CommandError{message: "A Pokemon name must be provided"}
	}
	pokemon := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	data, getErr := cfg.ApiGet(url)
	if getErr != nil {
		return getErr
	}
	pkmn := Pokemon{}
	jsonErr := json.Unmarshal(data, &pkmn)
	if jsonErr != nil {
		println("Error unmarshaling json")
		return jsonErr
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pkmn.Name)
	// A pokemon worth 0 exp is guaranteed, and chance can never be 0 but can reach an infinitesimally small number
	exp := float32(pkmn.BaseExperience)
	chance := float32(1.0 - exp/(exp+100.0))
	if rand.Float32() <= chance {
		fmt.Printf("%v was caught!\n", pkmn.Name)
		cfg.Pokedex.Add(pkmn)
	} else {
		fmt.Printf("%v escaped!\n", pkmn.Name)
	}
	return nil
}
