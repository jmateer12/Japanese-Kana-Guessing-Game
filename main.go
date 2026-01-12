package main

import "fmt"

func main() {
	hiragana := map[string]string{
		"あ": "a",
		"い": "i",
		"う": "u",
		"え": "e",
		"お": "o",
		"か": "ka",
		"き": "ki",
		"く": "ku",
		"け": "ke",
		"こ": "ko",
		"さ": "sa",
		"し": "shi",
		"す": "su",
		"せ": "se",
		"そ": "so",
		"た": "ta",
		"ち": "chi",
		"つ": "tsu",
		"て": "te",
		"と": "to",
		"な": "na",
		"に": "ni",
		"ぬ": "nu",
		"ね": "ne",
		"の": "no",
		"は": "ha",
		"ひ": "hi",
		"ふ": "fu",
		"へ": "he",
		"ほ": "ho",
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
