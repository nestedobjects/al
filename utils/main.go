package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// RemoveFile : Deletes given file
func RemoveFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Unable to clear your existing alias")
		return
	}
}

// FileDoesNotExist : Check if file does not exist
func FileDoesNotExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}

// ReadAlias : Reads all alias in given config file and populates in a map
func ReadAlias(fileName string, aliases map[string]string) {
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

// AppendAlias : Adds new alias to given config file
func AppendAlias(fileName string, newEntry string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("Unable to read your existing alias")
		return
	}
	if _, err = file.WriteString(newEntry); err != nil {
		fmt.Println("Unable to add new alias")
	}
}

// OverWriteFile : Overwrites the existing config file with new content
func OverWriteFile(fileName string, newAliases string) {
	RemoveFile(fileName)
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
