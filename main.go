package main

import "fmt"

func main() {
	hiragana := map[string]string{
		"あ": "a",
		"い": "i",
		"う": "u",
	}

	answersCorrect := 0

	for character := range hiragana {
		guess := ""

		fmt.Printf("What does this character represent %s: ", character)
		fmt.Scanln(&guess)

		if guess == hiragana[character] {
			answersCorrect += 1
		}
	}

	fmt.Println("Answers correct:", answersCorrect)

}
