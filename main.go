package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     string
	Previous string
}

func main() {
	config := config{}
	commands := config.commands()
	reader := bufio.NewScanner(os.Stdin)

	for {
		print("PokeDex > ")
		hasTokens := reader.Scan()
		if !hasTokens {
			return
		}

		token := reader.Text()
		lower := strings.ToLower(token)
		c, ok := commands[lower]
		if !ok {
			fmt.Printf("'%s' is not a valid command\n", token)
			continue
		}

		err := c.callback()

		if err != nil {
			if err.Error() == "exit" {
				return
			} else {
				print(err.Error())
				return
			}
		}
		println()
	}
}

func (cfg *config) commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cfg.commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cfg.commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays pokemon locations",
			callback:    cfg.commandLocationsFwd,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous pokemon locations",
			callback:    cfg.commandLocationsBck,
		},
	}
}

func (cfg *config) commandHelp() error {
	c := cfg.commands()
	print("Usage:\n\n")
	for _, command := range c {
		println(fmt.Sprintf("%s: %s", command.name, command.description))
	}
	return nil
}

func (cfg *config) commandExit() error {
	return errors.New("exit")
}
