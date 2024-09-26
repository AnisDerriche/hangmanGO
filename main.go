package main

/*

Mots random dans fichier words.txt

10 essais
Lettres révélées : len(word) / 2 - 1

si lettre non presente : message erreur et nbr essais restant
si lettre présente, affiche toutes les lettres présentent dans le mot
le programme continue tant que le mot n'a pas été trouvé ou que le nombre d'essaie n'est pas egale à 0

hangman position :
10 positions différentes
	hauteur de 7 (8 avec retour a la ligne)

exemple :
$> ./hangman words.txt
Good Luck, you have 10 attempts.
_ _ _ _ O

Choose: E
_ E _ _ O

Choose: A
Not present in the word, 9 attempts remaining






=========
...



func :
- choisir le mot de manière random
- afficher le mot et cachant des lettres random
- lettre proposé
- nombre restant d'essais et affichage du pendu

*/

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// var i string
	// print(reponse)
	// fmt.Printf("entre une lettre :")
	// fmt.Scan(&i)
	// fmt.Println(lettre_propose(&i))
	//fmt.Println(nombre_essai())
	//fmt.Println(affPendu())
	test()
}

var reponse bool = false

func MotRandom() {
	mot, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines []string
	scanner := bufio.NewScanner(mot)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	i := rand.Intn(len(lines) - 1)
	word := lines[i]
	fmt.Println(word)
	mot.Close()
	MelMot(word)
}

func MelMot(word string) {
	rune := []rune(word)
	long := len(word)
	letter := len(word)/2 - 1
	if long <= 3 {
		letter = 1
	}
	for i := 1; i <= long-letter; i++ {
		n := rand.Intn(len(word))
		if rune[n] == 95 {
			i--
		}
		rune[n] = 95
	}
	nword := string(rune)
	fmt.Println(nword)
}

func lettre_propose(n *string) string {
	return ("lettre --> " + *n)
}

func nombre_essai() int {
	nbr_essai := 10
	return nbr_essai
}

func affPendu() int {
	var pendu []string
	nombre_essai()
	if nombre_essai() == 10 {
		file, err := os.ReadFile("hangman.txt")
		for _, char := range file {
			pendu = append(pendu, string(char))
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(pendu)
		}
	}
	return 3
}

func test() {
	pos, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines []string
	scanner := bufio.NewScanner(mot)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	i := rand.Intn(len(lines) - 1)
	word := lines[i]
	fmt.Println(word)
	mot.Close()
	MelMot(word)
}
