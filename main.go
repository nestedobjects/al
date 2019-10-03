package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rafi993/al/utils"
)

func main() {
	argsWithoutProg := os.Args[1:]

	home, err := os.UserHomeDir()
	fileName := filepath.Join(home, "al.txt")

	if err == nil {
		if len(argsWithoutProg) >= 1 {
			command := argsWithoutProg[0]
			switch command {
			case "list":
				utils.ListAlias(fileName)
			case "add":
				utils.AddAlias(fileName, argsWithoutProg)
			case "rm":
				fmt.Println("Removing alias")
			case "reset":
				fmt.Println("Clear all alias")
			default:
				utils.CallAlias(fileName, argsWithoutProg[0])
			}

		} else {
			fmt.Println(`
	al ls - List all alias\n
	al add alias cmd - Add new alias\n
	al rm alias -  Remove alias
	al reset - remove all alias and start fresh 
		`)
		}
	} else {
		fmt.Println("Unable to get your home directory")
	}

}
