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
	nombre_essai()
	affPendu()
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
	fmt.Println(lines[i])
	mot.Close()
}

func lettre_propose(n *string) string {
	return ("lettre --> " + *n)
}

func nombre_essai() int {
	nbr_essai := 10
	return nbr_essai
}

func affPendu() {
	var pendu []string
	file, err := os.ReadFile("hangman.txt")
	for _, char := range file {
		pendu = append(pendu, string(char))
	}
	if nombre_essai() == 10 {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(pendu[0:78])
		}
	}
	// if nombre_essai() == 9 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[80:157])
	// 	}
	// }
	// if nombre_essai() == 8 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[158:235])
	// 	}
	// }
	// if nombre_essai() == 7 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[237:313])
	// 	}
	// }
	// if nombre_essai() == 6 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[320:391])
	// 	}
	// }
	// if nombre_essai() == 5 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[397:469])
	// 	}
	// }
	// if nombre_essai() == 4 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[475:547])
	// 	}
	// }
	// if nombre_essai() == 3 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[552:625])
	// 	}
	// }
	// if nombre_essai() == 2 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[630:703])
	// 	}
	// }
	// if nombre_essai() == 1 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[708:789])
	// 	}
	// }
	// if nombre_essai() == 0 {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(pendu[787:858])
	// 	}
	// }
	return
}
