package makejson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func makeJSON() {

	hashmap := make(map[string]string)

	var name string
	var adress string

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	input.Scan()
	name = input.Text()

	fmt.Print("Enter your adress: ")
	input.Scan()
	adress = input.Text()

	hashmap["name"] = name
	hashmap["adress"] = adress

	jsonObject, err := json.Marshal(hashmap)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(jsonObject))

}
