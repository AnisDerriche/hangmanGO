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
	fmt.Println("Good Luck, you have 10 attempts.")
	MotRandom()
	MelMot(MotRandom())
	var i string
	fmt.Printf("entre une lettre :")
	fmt.Scan(&i)
	fmt.Println(lettre_propose(&i))
	//nombre_essai()
	//affPendu()
}

var reponse bool = false

func MotRandom() string {
	/* prends un mot au hasard dans words.txt */
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
	return word
}

func MelMot(word string) {
	/*cache les lettres du mot*/
	mot := MotRandom() //reprise du mot pioché au hasard dans la fonction MotRandom
	rune := []rune(mot)
	long := len(mot)
	letter := len(mot)/2 - 1
	if long <= 3 {
		letter = 1
	}
	for i := 1; i <= long-letter; i++ {
		n := rand.Intn(len(mot))
		if rune[n] == 95 /*95(ASCII) == _ */ {
			i--
		}
		rune[n] = 95
	}
	newword := string(rune) // le mot est en rune etant donné qu'on l'a modif avec un caractère ASCII
	fmt.Println(newword)
}

func lettre_propose(n *string) string {
	/* fonction en plus qui permet de voir la lettre mis mais peut être supp et placer ailleur*/
	return ("lettre --> " + *n)
}

func nombre_essai() int {
	/* affiche le nombre d'essai restant (de base a 10)*/
	nbr_essai := 10
	return nbr_essai
}

func Rep() {
	lipos := 8

	affPendu()
	lipos += 8
}

func affPendu() {
	/* affiche le pendu selon le nombre d'essais restant */
	file, err := os.Open("hangman.txt")
	liposbis := -1
	lipos := 7
	if nombre_essai() == 10 {
		if err != nil {
			fmt.Println(err)
		} else {
			scanner := bufio.NewScanner(file)
			lineNumber := 5
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
	}
	return
}
