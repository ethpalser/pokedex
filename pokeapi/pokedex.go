package pokeapi

import "fmt"

type Pokedex struct {
	data map[string]Pokemon
}

func (p *Pokedex) Add(pkmn Pokemon) {
	p.data[pkmn.Name] = pkmn
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	dat, ok := p.data[name]
	if !ok {
		return Pokemon{}, false
	}
	return dat, true
}

func (p *Pokedex) CommandPokedex(args ...string) error {
	println("Your Pokedex:")
	for key := range p.data {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}
