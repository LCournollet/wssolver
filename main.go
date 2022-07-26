package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

var array []string

var field = [10][10]rune{
	{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
	{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
	{'ê', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
	{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
	{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'é'},
	{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
	{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
	{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
	{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
	{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
}

var Test_Field = [1][10]rune{
	{'c', 'a', 'f', 'e', 'p', 'r', 'o', 'u', 't', 'z'},
}

var Test_Field_2 = [4][1]rune{
	{'c'},
	{'a'},
	{'f'},
	{'é'},
}

//======================================================================================================================
//====================================================[↓ FUNCTION ↓]====================================================
//======================================================================================================================

func PrintField(field [10][10]rune) {
	LOT_OF_SPACE_SHEEESH()

	for i := 0; i < 10; i++ {
		z01.PrintRune('[')

		for j := 0; j < 10; j++ {
			z01.PrintRune(field[i][j])

			if j != 9 {
				z01.PrintRune(' ')
			}

		}

		z01.PrintRune(']')
		z01.PrintRune('\n')
	}

	LOT_OF_SPACE_SHEEESH()
}

func LOT_OF_SPACE_SHEEESH() {
	for i := 0; i < 3; i++ {
		z01.PrintRune('\n')
	}
}

//======================================================================================================================
//====================================================[↓ REAL SHIT ↓]===================================================
//======================================================================================================================
func GetArrayFromWords() []string {
	readFile, _ := os.Open("words.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}

	readFile.Close()
	return array
}

func NewArray(FirstLetter rune, SecondLetter rune, Words []string) []string {
	WordIncrease := 0
	NewArray := []string{}

	for i := 0; i < len(Words); i++ {

		if FirstLetter == rune(Words[i][0]) && SecondLetter == rune(Words[i][1]) {
			NewArray = append(NewArray, Words[WordIncrease])
		}

		if WordIncrease != len(Words)-1 {
			WordIncrease++
		}
	}
	return NewArray
}

func WordScrabbleSolver(field [1][10]rune) {
	words := GetArrayFromWords()
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			for WordIncrease := 0; WordIncrease < len(words); WordIncrease++ {
				if field[i][j] == rune(words[WordIncrease][0]) {
					if j != len(field[i])-1 {
						if field[i][j+1] == rune(words[WordIncrease][1]) {
							// NewWordsList := NewArray(field[i][j], field[i][j+1], words)
							VeryfHorizontal(i, j, words[WordIncrease], field)
						}
					}
				}
			}
		}
	}
}

func VeryfHorizontal(x int, y int, NewWordsList string, field [1][10]rune) []string {
	array := []string{}
	Count := 0
	// WordIncrease := 0
	LetterIncrease := 0
	for j := y; j < len(field[x]); j++ {
		if field[x][j] == rune(NewWordsList[LetterIncrease]) {
			Count += 1
			if Count == len(NewWordsList) {
				array = append(array, NewWordsList)
				fmt.Println("FOUNDED THE WORD: ", array)
				return array
			}
			if LetterIncrease != len(NewWordsList)-1 {
				LetterIncrease++
			}
		}
		// } else if WordIncrease != len(NewWordsList)-1 {
		// 	Count = 0
		// 	j = y
		// 	WordIncrease++
		// }
	}
	return array
}

func VeryfVertical(field [10][10]rune) {

}

func VeryfBottomRight(field [10][10]rune) {

}
func VeryfTopRight(field [10][10]rune) {

}

//======================================================================================================================
//====================================================[↓ MAIN ↓]=======================================================
//======================================================================================================================
func main() {
	// y := 0
	// for j := y; j < 5; j++ {
	// 	fmt.Println(j)
	// 	// k := y
	// 	fmt.Println("Hello")
	// 	if j == 4 {
	// 		j = 0
	// 	}
	// }
	// PrintField(field)
	WordScrabbleSolver(Test_Field)
}
