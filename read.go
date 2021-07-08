package filereader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func first20chars(s string) string {
	runes := []rune(s)
	return string(runes[:])
}

func Reader() []Person {
	var fileName string
	personSlice := make([]Person, 0)
	var personObj Person

	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)

	fileHandle, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		if len(s[0]) > 20 {
			s[0] = first20chars(s[0])
		}
		if len(s[1]) > 20 {
			s[1] = first20chars(s[1])
		}

		personObj.fName, personObj.lName = s[0], s[1]
		personSlice = append(personSlice, personObj)
	}
	// This is not for here
	for _, v := range personSlice {
		fmt.Println(v.fName, v.lName)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return personSlice
}
