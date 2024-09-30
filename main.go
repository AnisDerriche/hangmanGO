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
!-





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
<<<<<<< HEAD
=======
	// fmt.Println("Good Luck, you have 10 attempts.")
	// MotRandom()
	// MelMot(MotRandom())
	// var i string
	// fmt.Printf("entre une lettre :")
	// fmt.Scan(&i)
	// fmt.Println(lettre_propose(&i))
	// nombre_essai()
	// ffPendu()
	game()
}

var reponse bool = false

func game() {
	/* prends un mot au hasard dans words.txt */
>>>>>>> 8a15dbb2bd70e999bcad032313fd70354abb46ad
	mot, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines []string //met dans une liste les mots present dans words.txt
	scanner := bufio.NewScanner(mot)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	i := rand.Intn(len(lines) - 1)
	word := lines[i] //le mots pris au hasard
	mot.Close()
<<<<<<< HEAD
	//MelMot(word)

	fmt.Println("Good Luck, you have 10 attempts.")

	MelMot(word)
	/*var i string
	fmt.Printf("entre une lettre :")
	fmt.Scan(&i)
	fmt.Println(lettre_propose(&i))
	*/
	Rep()
	//fmt.Println(nombre_essai())
	//fmt.Println(affPendu())
	//nombre_essai()
	//affPendu()
}

var reponse bool = false

/*
	func MotRandom() string {
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
		//fmt.Println(word)
		mot.Close()
		return word
		//MelMot(word)
	}
*/
func MelMot(word string) {
	rune := []rune(word)
=======
	/* fin de la partie du prenage de mot */

	//---------------------------------------------------------------------------------------------------------

	/*cache les lettres du mot*/
	rune1 := []rune(word)
>>>>>>> 8a15dbb2bd70e999bcad032313fd70354abb46ad
	long := len(word)
	letter := len(word)/2 - 1
	if long <= 3 {
		letter = 1
	}
	for i := 1; i <= long-letter; i++ {
		n := rand.Intn(len(word))
		if rune1[n] == 95 /*95(ASCII) == _ */ {
			i--
		}
		rune1[n] = 95
	}
	newword := string(rune1) // le mot est en rune etant donné qu'on l'a modif avec un caractère ASCII
	fmt.Println(newword)
	/* fin du cachage de lettre */

<<<<<<< HEAD
func lettre_propose(n *string) string {
	return ("lettre --> " + *n)
}

func nombre_essai() int {
	nbr_essai := 10
	return nbr_essai
}

func Rep() {
	fmt.Println(word)
	/*word = test
	fo

	lipos := 8

	affPendu()
	lipos += 8*/
}

func affPendu() {
	file, err := os.Open("hangman.txt")
	liposbis := -1
=======
	//---------------------------------------------------------------------------------------------------------
	/*nombre d'essai de base */
	liposbis := 0
>>>>>>> 8a15dbb2bd70e999bcad032313fd70354abb46ad
	lipos := 7
	nbr_essai := 10
	for nbrDeTour := -nbr_essai; nbrDeTour <= nbr_essai; nbrDeTour++ {
		/* demande de la lettre */
		var lettre string
		fmt.Printf("entre une lettre :")
		fmt.Scan(&lettre)
		fmt.Println("lettre --> ", word)
		/* affiche le pendu selon le nombre d'essais restant */
		file, err := os.Open("hangman.txt")
		if nbr_essai >= 0 {
			if err != nil {
				fmt.Println("aze")
			} else {
				scanner := bufio.NewScanner(file)
				lineNumber := 0
				for scanner.Scan() { //cette boucle permet ,selon le nombre d'essai restant dafficher le pendu.
					if lineNumber < liposbis {
						lineNumber++
						continue
					}
					if lineNumber > lipos {
						break
					}
					fmt.Println(scanner.Text())
					lineNumber++
				}
			}
			nbr_essai--
			liposbis += 8
			lipos += 8
			continue
		}
	}
}
