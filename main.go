package main

import "fmt"

func main() {
	/*
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
			"ま": "ma",
			"み": "mi",
			"む": "mu",
			"め": "me",
			"も": "mo",
			"や": "ya",
			"ゆ": "yu",
			"よ": "yo",
			"ら": "ra",
			"り": "ri",
			"る": "ru",
			"れ": "re",
			"ろ": "ro",
			"わ": "wa",
			"ゐ": "wi",
			"を": "wo",
			"ん": "n",
		}
	*/

	katakana := map[string]string{
		"ア": "a",
		"イ": "i",
		"ウ": "u",
		"エ": "e",
		"オ": "o",
		"カ": "ka",
		"キ": "ki",
		"ク": "ku",
		"ケ": "ke",
		"コ": "ko",
		"サ": "sa",
	}

	answersCorrect := 0

	for character := range katakana {
		guess := ""

		fmt.Printf("What does this character represent %s: ", character)
		fmt.Scanln(&guess)

		if guess == katakana[character] {
			answersCorrect += 1
			fmt.Printf("\033[32mAnswer Correct!!!\033[0m\n")
		} else {
			fmt.Printf("\033[31mAnswer Incorrect, answer should be %s\033[0m\n", katakana[character])
		}
	}

	percent := (float64(answersCorrect) / float64(len(katakana))) * 100

	fmt.Printf("Answers correct: %d/%d\n", answersCorrect, len(katakana))
	fmt.Printf("Percent correct: %.1f%%\n", percent)

}
