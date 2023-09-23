package cli

type cliCommand struct {
	name        string
	description string
	callback    func(*CliConfig, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    mapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists the Pokemon in your current location area. You don't have a starting location, so make sure you 'travel' to a location area first.",
			callback:    explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a pokemon and add it to your pokedex",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View information about a caught pokemon",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all your caught pokemon",
			callback:    pokedex,
		},
		"travel": {
			name:        "travel <area_name>",
			description: "Travel to a specific location area",
			callback:    travel,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
	}
}
