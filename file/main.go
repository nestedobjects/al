package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/rafi993/al/utils"
)

// ListAlias : List all alias in given config file
func ListAlias(fileName string) {
	if utils.FileDoesNotExist(fileName) {
		fmt.Println("You dont have any alias points")
	} else {
		aliases, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Unable to list your alias points")
			return
		}
		fmt.Println(string(aliases))

	}
}

// AddAlias : Adds alias to given config file
func AddAlias(fileName string, args []string) {
	if len(args) > 2 {

		alias := args[1]
		command := args[2]
		options := []string{}

		if len(args) > 3 {
			options = args[3:]
		}

		newEntry := alias + " -> " + command + " " + strings.Join(options[:], " ") + "\n"
		aliases := make(map[string]string)

		if utils.FileDoesNotExist(fileName) {
			utils.AppendAlias(fileName, newEntry)
		} else {

			utils.ReadAlias(fileName, aliases)

			if _, present := aliases[alias]; present {
				fmt.Printf("%v already points to %v", alias, aliases[alias])
			} else {
				utils.AppendAlias(fileName, newEntry)
			}
		}
	} else {
		fmt.Println("use the following syntax\n\nal add alias cmd")
	}

}

// CallAlias : Calls given alias if it exist in the config file
func CallAlias(fileName string, alias string) {

	if utils.FileDoesNotExist(fileName) {
		fmt.Println("You dont have any alias")
	} else {
		aliases := make(map[string]string)

		utils.ReadAlias(fileName, aliases)

		if _, present := aliases[alias]; present {

			command := strings.Split(aliases[alias], " ")

			output, err := exec.Command(command[0], command[1:]...).CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			fmt.Println(string(output))

		} else {
			fmt.Println("Alias point does not exist")
		}
	}

}

// ResetAlias : Removes all exist alias
func ResetAlias(fileName string) {
	utils.RemoveFile(fileName)
}

// RemoveAlias : Remove given alias from config file
func RemoveAlias(fileName string, alias string) {
	aliases := make(map[string]string)
	utils.ReadAlias(fileName, aliases)

	if _, ok := aliases[alias]; ok {
		newAliases := ""
		delete(aliases, alias)

		for k, v := range aliases {
			newAliases = k + " -> " + v + "" + newAliases
		}

		utils.OverWriteFile(fileName, newAliases)
	} else {
		fmt.Println("Alias does not exist")
	}
}
