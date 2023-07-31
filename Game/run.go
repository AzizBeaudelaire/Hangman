package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartGame : Start a new game
func StartGame(WordsSlice []string) GameData {
	fmt.Println("\n")
	fmt.Println(ToAsciiArt("HANGMAN"))
	fmt.Println()
	fmt.Println("\033[32m", "Good Luck, you have 10 attempts.\n\n", "\033[0m") // green - color reset
	WordToFind := TakeRandomWord(WordsSlice)
	Data := NewGame(WordToFind)
	Data.ToFind = WordToFind
	Data.Tries = InitialLetters(Data).Tries
	Data.Index = RevealInitialLetters(Data).Index
	Data = RevealInitialLetters(Data)
	return Data
}

func Run(arg string) {
	fmt.Println(arg)
	var Data GameData
	if arg == "words.txt" {
		file := "Extra/words.txt"
		WordsSlice := GetFile(file)
		if TestFile(WordsSlice) == false {
			return
		}
		Data = StartGame(WordsSlice)
	} else if arg == "--startWithsave.txt" {
		file := "Saves/save.txt"
		Data = StartWithFlag(file)
	} else {
		fmt.Println("\033[31m", "Please enter words.txt to start a new game or --startWith save.txt to continue your last game.", "\033[0m")
		return
	}

	fmt.Println(ToAsciiArt(Data.Word))

	reader := bufio.NewReader(os.Stdin)
	for Data.Attempts > 0 {
		fmt.Print("Choose:")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(strings.ToUpper(guess))
		for IntputTesting(guess, Data) == false {
			fmt.Print("Choose:")
			guess, _ = reader.ReadString('\n')
			guess = strings.TrimSpace(strings.ToUpper(guess))
		}
		Data.Tries = append(Data.Tries, guess)
		if len(guess) < 2 {
			Data = FindLetter(Data)
			if Data.LetterCheck == false {
				Data.Attempts--
				fmt.Print("\033[35m", "Not present in the word ", "\033[36m")
				fmt.Printf("%v attempts remaining\n", Data.Attempts)
				fmt.Print("\033[0m") // color reset
				fmt.Println(OpenJose()[9-Data.Attempts])
				fmt.Println(ToAsciiArt(Data.Word))
			} else {
				Data.Word = RevealLetters(Data)
				fmt.Println("\033[32m", "Good guess !", "\033[0m")
				fmt.Println(ToAsciiArt(Data.Word))
				if WordGuessed(Data) {
					fmt.Println("\033[92m", "Congrats ! You find the word.", "\033[0m")
					break
				}
			}
		} else { // if we guess a word instead of a letter or if we want to save and quit the game
			if guess == "STOP" {
				StopAndSaveGame(Data)
				fmt.Println("\033[32m", "Your game is saved. Enter --startWith save.txt in argument to restart your last save", "\033[0m")
				return
			}
			if GuessingWord(guess, Data) == false {
				Data.Attempts = Data.Attempts - 2
				if Data.Attempts > 0 {
					fmt.Print("\033[36m")
					fmt.Printf("You have %v attempts left\n", Data.Attempts)
					fmt.Print("\033[0m")
					fmt.Println(OpenJose()[9-Data.Attempts])
					fmt.Println(ToAsciiArt(Data.Word))
				} else {
					fmt.Println(OpenJose()[9])
				}
			} else {
				Data.Word = guess
				fmt.Println(ToAsciiArt(Data.Word))
				fmt.Println("\033[92m", "Congrats ! You find the word.", "\033[0m")
				break
			}
		}
	}
	if Data.Attempts <= 0 {
		fmt.Println("\033[31m", "Sorry, you loose!, word was:", strings.ToUpper(Data.ToFind), "\033[0m")
	}
}
