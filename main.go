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
<<<<<<< HEAD

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	MotRandom()
}

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
	fmt.Println(lines[i])
	mot.Close()
=======
import (
	"fmt"
)

var reponse bool = false

func main() {
	var i string
	print(reponse)
	fmt.Printf("entre une lettre :")
	fmt.Scan(&i)
	fmt.Println(lettre_propose(&i))
	fmt.Println(nombre_essai())
}

func lettre_propose(n *string) string {
	return ("lettre --> " + *n)
}

func nombre_essai() int {
	nbr_essai := 10
	return nbr_essai
>>>>>>> e38ea66628e54f5c2cc27a80d14123a3d70647cf
}
