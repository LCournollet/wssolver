package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

var array []string

var field = [10][10]rune{
	{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
	{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
	{'e', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
	{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
	{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'e'},
	{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
	{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
	{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
	{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
	{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
}

var Testfield = [10][10]rune{
	{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
	{'z', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
	{'e', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'z'},
	{'p', 'z', 'w', 'z', 'e', 'r', 'e', 'a', 'a', 'f'},
	{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'e'},
	{'s', 'a', 'z', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
	{'p', 'a', 'e', 'z', 'i', 't', 'c', 'd', 't', 'e'},
	{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
	{'f', 'o', 'u', 'r', 'z', 'h', 'e', 't', 't', 'e'},
	{'b', 'a', 'z', 'e', 's', 't', 'h', 'n', 'w', 's'},
}

//======================================================================================================================
//====================================================[↓ FUNCTION ↓]====================================================
//======================================================================================================================

// Impression du tableau dans la console
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

//----------------------------------------------------------------------------------------------------------------------

// Juste pour foutre plein d'espaces
func LOT_OF_SPACE_SHEEESH() {
	for i := 0; i < 3; i++ {
		z01.PrintRune('\n')
	}
}

//----------------------------------------------------------------------------------------------------------------------

// Impression des mots trouvés avec mise en forme
func PrintAllWords(Words []string) {
	z01.PrintRune('[')
	for i := 0; i < len(Words); i++ {
		for j := 0; j < len(Words[i]); j++ {
			z01.PrintRune(rune(Words[i][j]))
		}
		if i != len(Words)-1 {
			z01.PrintRune(' ')
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

//----------------------------------------------------------------------------------------------------------------------
// Permet de voir si on a trouvé un/des mots ou pas
func PrintNoWords(Bool bool) {
	if Bool == false {
		Array := [1]string{"Words Founded: "}
		for i := 0; i < len(Array); i++ {
			for j := 0; j < len(Array[i]); j++ {
				z01.PrintRune(rune(Array[i][j]))
			}
		}
	}
	if Bool == true {
		Array := [1]string{"No words found"}
		for i := 0; i < len(Array); i++ {
			for j := 0; j < len(Array[i]); j++ {
				z01.PrintRune(rune(Array[i][j]))
			}
			z01.PrintRune('\n')
		}
	}
}

//======================================================================================================================
//====================================================[↓ REAL SHIT ↓]===================================================
//======================================================================================================================

// Permet de mettre chaque mot dans un array
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

//----------------------------------------------------------------------------------------------------------------------
// Fonction principale
func WordScrabbleSolver(field [10][10]rune) {
	FinalArray := []string{}
	Temp := []string{}
	words := GetArrayFromWords()

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {

			// On check chaque mot de la liste
			for WordIncrease := 0; WordIncrease < len(words); WordIncrease++ {

				// On vérifie si la lettre est la même que la première de l'un des mots du fichier texte (si non on change de lettre dans le tableau)
				if field[i][j] == rune(words[WordIncrease][0]) {

					// Verification Horizontal
					if j != len(field[i])-1 && field[i][j+1] == rune(words[WordIncrease][1]) {
						Temp = VeryfHorizontal(i, j, words[WordIncrease], field)
						FinalArray = append(FinalArray, Temp...)

						// Verification de la verticale
					} else if i != len(field)-1 && field[i+1][j] == rune(words[WordIncrease][1]) {
						Temp = VeryfVertical(i, j, words[WordIncrease], field)
						FinalArray = append(FinalArray, Temp...)

						// Verification de la diagonale en partant de la gauche vers bas droite
					} else if i != len(field)-1 && j != len(field[i])-1 && field[i+1][j+1] == rune(words[WordIncrease][1]) {
						Temp = VeryfBottomRight(i, j, words[WordIncrease], field)
						FinalArray = append(FinalArray, Temp...)

						// Verification de la diagonale en partant de la gauche vers haut droite
					} else if i != 0 && j != len(field[i])-1 && field[i-1][j+1] == rune(words[WordIncrease][1]) {
						Temp = VeryfTopRight(i, j, words[WordIncrease], field)
						FinalArray = append(FinalArray, Temp...)

					}
				}
			}
		}
	}

	// Condition pour savoir si on a trouvé des mots ou non
	if len(FinalArray) > 1 {
		PrintNoWords(false)
		PrintAllWords(FinalArray)
	} else {
		PrintNoWords(true)
	}
}

//----------------------------------------------------------------------------------------------------------------------

// Cette fonction est utilisable pour toutes les directions possible, il suffit de changer les conditions (i et j) et laquelle de ces condiotions augmentes ou diminues (donc en gros tout ce qu'il y a en dessous c'est du copié collé, masterclass)
func VeryfHorizontal(x int, y int, NewWordsList string, field [10][10]rune) []string {
	array := []string{}
	Count := 0
	LetterIncrease := 0

	for j := y; j < len(field[x]); j++ {

		if field[x][j] == rune(NewWordsList[LetterIncrease]) {
			Count += 1
			// A chaque lettre qui correspond, j'ajoute +1 à "Count", donc si mon Count est égal à la longueur du mot, cela veux dire que j'ai trouvé le mot qui correspond
			if Count == len(NewWordsList) {
				array = append(array, NewWordsList)
				break
			}
			if LetterIncrease != len(NewWordsList)-1 {
				LetterIncrease++
			}
		} else {
			// Ne pas oublier de reset "Count" en cas d'echec
			Count = 0
		}
	}
	return array
}

//----------------------------------------------------------------------------------------------------------------------
// Copy+paste de la fonction précédente
func VeryfVertical(x int, y int, NewWordsList string, field [10][10]rune) []string {
	array := []string{}
	Count := 0
	LetterIncrease := 0

	for i := x; i < len(field); i++ {

		if field[i][y] == rune(NewWordsList[LetterIncrease]) {
			Count += 1

			if Count == len(NewWordsList) {
				array = append(array, NewWordsList)
				break
			}
			if LetterIncrease != len(NewWordsList)-1 {
				LetterIncrease++
			}

		} else {
			Count = 0
		}
	}
	return array
}

//----------------------------------------------------------------------------------------------------------------------
// Pareil que la précédente
func VeryfBottomRight(x int, y int, NewWordsList string, field [10][10]rune) []string {
	array := []string{}
	Count := 0
	LetterIncrease := 0
	i := x
	j := y

	for ; i < len(field); i++ {

		if field[i][j] == rune(NewWordsList[LetterIncrease]) {
			Count++

			if Count == len(NewWordsList) {
				array = append(array, NewWordsList)
				break
			}

			if LetterIncrease != len(NewWordsList)-1 {
				LetterIncrease++
			}

			if j != len(field[i])-1 {
				j++
			} else {
				break
			}

		} else {
			Count = 0
		}
	}
	return array
}

//----------------------------------------------------------------------------------------------------------------------
// Encore pareil toujours similaire à la précédente
func VeryfTopRight(x int, y int, NewWordsList string, field [10][10]rune) []string {
	array := []string{}
	Count := 0
	LetterIncrease := 0
	i := x
	j := y

	for ; j < len(field); j++ {

		if field[i][j] == rune(NewWordsList[LetterIncrease]) {
			Count++

			if Count == len(NewWordsList) {
				array = append(array, NewWordsList)
				break
			}

			if LetterIncrease != len(NewWordsList)-1 {
				LetterIncrease++
			}

			if i != 0 {
				i--
			} else {
				break
			}

		} else {
			Count = 0
		}
	}
	return array
}

//======================================================================================================================
//====================================================[↓ MAIN ↓]=======================================================
//======================================================================================================================
func main() {

	// Changer le "field" par le nom du nouveau tableau pour effectuer des tests
	FIELD := field

	PrintField(FIELD)
	WordScrabbleSolver(FIELD)
}
