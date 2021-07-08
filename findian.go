package findian

import "fmt"

func Findian() bool {
	var input string
	fmt.Printf("Enter a word:\n")
	fmt.Scan(&input)
	lenInput := len(input)
	if lenInput < 3 {
		return false
	}
	isValidFirstCharacter := input[0] == 105 || input[0] == (105-32)
	isValidEndCharacter := (input[lenInput-1] == 110) || (input[lenInput-1] == (110 - 32))
	isValid := isValidFirstCharacter && isValidEndCharacter
	for i := 1; i < lenInput-1; i++ {
		if input[i] == 97 || input[i] == (97-32) {
			if isValid {
				return true

			}
		}
	}

	return false
}
