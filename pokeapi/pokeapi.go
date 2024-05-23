package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/ethpalser/pokedex/internal/cache"
)

type Config struct {
	Next     string      // Next URL of the resource last fetched
	Previous string      // Previous URL of the resource last fetched
	Cache    cache.Cache // Cache for api calls
	Pokedex  *Pokedex    // Pointer to Pokedex database
}

type CommandError struct {
	message string
}

func (err *CommandError) Error() string {
	return err.message
}

func (err *CommandError) Is(e error) bool {
	_, ok := e.(*CommandError)
	return ok
}

func NewConfig() Config {
	return Config{
		Next:     "",
		Previous: "",
		Cache:    cache.NewCache(30 * time.Minute),
		Pokedex: &Pokedex{
			data: map[string]Pokemon{},
		},
	}
}

func (c *Config) ApiGet(url string) ([]byte, error) {
	// Check if data exists in Cache
	dat, ok := c.Cache.Get(url)
	if ok {
		return dat, nil
	}
	// Perform an API request, as data is not in Cache
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, &CommandError{}
	}

	// Read all data from the io.ReadCloser into a byte array
	body, ioErr := io.ReadAll(resp.Body)
	if ioErr != nil {
		println("An io error occurred")
		return nil, ioErr
	}
	// Save data in Cache
	c.Cache.Add(url, body)
	return body, nil
}
