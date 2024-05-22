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

func main() {
	commands := commands()
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
	}
}

func commands() map[string]cliCommand {
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
	}
}

func commandHelp() error {
	c := commands()
	print("Usage:\n\n")
	for _, command := range c {
		println(fmt.Sprintf("%s: %s", command.name, command.description))
	}
	println()
	return nil
}

func commandExit() error {
	return errors.New("exit")
}
