package pokeapi

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
