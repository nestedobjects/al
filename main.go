package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rafi993/al/file"
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
				file.ListAlias(fileName)
			case "add":
				file.AddAlias(fileName, argsWithoutProg)
			case "rm":
				file.RemoveAlias(fileName, argsWithoutProg[1])
			case "reset":
				file.ResetAlias(fileName)
			default:
				file.CallAlias(fileName, argsWithoutProg[0])
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
