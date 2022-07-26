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
	{'p', 'r', 'e', 's', 't', 'a', 't', 'a', 'i', 'z'},
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

							NewWordsList := NewArray(field[i][j], field[i][j+1], words)
							VeryfHorizontal(i, j, NewWordsList, field)
						}
					}
				}
			}
		}
	}
}

func VeryfHorizontal(x int, y int, NewWordsList []string, field [1][10]rune) []string {
	fmt.Println("-----------------------RESET-----------------------------")
	array := []string{}
	Count := 0
	WordIncrease := 0
	LetterIncrease := 0
	fmt.Println("Y: ", y)
	k := y
	for j := y; j < len(field[x]); j++ {
		Word := NewWordsList[WordIncrease]
		if field[x][j] == rune(NewWordsList[WordIncrease][LetterIncrease]) {
			fmt.Println("hey")
			fmt.Println("J BUT YOU KNOW: ", j)
			fmt.Println(string(field[x][j]), "//", string(NewWordsList[WordIncrease][LetterIncrease]))
			Count += 1
			if Count == len(NewWordsList[WordIncrease]) {
				array = append(array, NewWordsList[WordIncrease])
				fmt.Println("FOUNDED THE WORD: ", array)
				return array
			}
			if LetterIncrease != len(Word)-1 {
				LetterIncrease++
			}
		} else if WordIncrease != len(NewWordsList)-1 {
			Count = 0
			j = k
			fmt.Println("J: ", j)
			WordIncrease++
		}
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
	// k := -1
	// for j := 0; j < 1; j++ {
	// 	fmt.Println("Hello")
	// 	j = k
	// 	fmt.Println(j)
	// }
	// PrintField(field)
	WordScrabbleSolver(Test_Field)
}
