package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func fileDoesNotExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}

func readAlias(fileName string, aliases map[string]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Unable to read your existing alias")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading your existing alias", err)
			return
		}
		currentLine := strings.Split(line, " -> ")
		aliases[currentLine[0]] = currentLine[1]
	}
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to create config file")
		return
	}
}

func appendAlias(fileName string, newEntry string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to read your existing alias")
		return
	}
	if _, err = file.WriteString(newEntry); err != nil {
		fmt.Println("Unable to add new alias")
	}
}

func overWriteFile(fileName string, newAliases string) {
	ResetAlias(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to read your existing alias")
		return
	}
	if _, err = file.WriteString(newAliases); err != nil {
		fmt.Println("Unable remove alias")
	}
}

func ListAlias(fileName string) {
	if fileDoesNotExist(fileName) {
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

		if fileDoesNotExist(fileName) {
			createFile(fileName)
			appendAlias(fileName, newEntry)
		} else {

			readAlias(fileName, aliases)

			if _, present := aliases[alias]; present {
				fmt.Printf("%v already points to %v", alias, aliases[alias])
			} else {
				appendAlias(fileName, newEntry)
			}
		}
	} else {
		fmt.Println("use the following syntax\n\nal add alias cmd")
	}

}

func CallAlias(fileName string, alias string) {

	if fileDoesNotExist(fileName) {
		fmt.Println("You dont have any alias")
	} else {
		aliases := make(map[string]string)

		readAlias(fileName, aliases)

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

func ResetAlias(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Unable to clear your existing alias")
		return
	}
}

func RemoveAlias(fileName string, alias string) {
	aliases := make(map[string]string)
	readAlias(fileName, aliases)

	if _, ok := aliases[alias]; ok {
		newAliases := ""
		delete(aliases, alias)

		for k, v := range aliases {
			newAliases = k + " -> " + v + "" + newAliases
		}

		overWriteFile(fileName, newAliases)
	} else {
		fmt.Println("Alias does not exist")
	}
}
