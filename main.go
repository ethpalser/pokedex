package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethpalser/pokedex/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args ...string) error
}

func main() {
	config := pokeapi.NewConfig()
	commands := commands(&config)
	reader := bufio.NewScanner(os.Stdin)

	for {
		print("PokeDex > ")
		hasTokens := reader.Scan()
		if !hasTokens {
			return
		}

		token := reader.Text()
		split := strings.Split(token, " ")
		if len(split) < 1 {
			continue
		}

		command := split[0]
		args := split[1:]

		c, ok := commands[command]
		if !ok {
			fmt.Printf("'%s' is not a valid command\n", token)
			continue
		}

		err := c.callback(args...)

		if err != nil {
			if err.Error() == "exit" {
				return
			} else if errors.Is(err, &pokeapi.CommandError{}) {
				fmt.Printf("Failed to %v %v\n", c.name, err.Error())
			} else {
				print(err.Error())
				return
			}
		}
	}
}

func commands(cfg *pokeapi.Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays pokemon locations",
			callback:    cfg.CommandLocationsFwd,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous pokemon locations",
			callback:    cfg.CommandLocationsBck,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon at a location",
			callback:    cfg.CommandExplore,
		},
	}
}

func commandHelp(args ...string) error {
	// argument can be nil as its method's are not needed
	c := commands(nil)
	print("Usage:\n\n")
	for _, command := range c {
		println(fmt.Sprintf("%s: %s", command.name, command.description))
	}
	return nil
}

func commandExit(args ...string) error {
	return errors.New("exit")
}
