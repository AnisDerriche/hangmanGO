package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	game()
}

// Structure pour stocker l'état de jeu
type GameState struct {
	Word              string
	MaskedWord        string
	RemainingAttempts int
}

// Fonction pour sauvegarder l'état de jeu dans save.txt
func saveGame(state GameState) error {
	file, err := os.Create("save.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n%s\n%d\n", state.Word, state.MaskedWord, state.RemainingAttempts))
	if err != nil {
		return err
	}
	fmt.Println("Partie sauvegardée dans save.txt.")
	return nil
}

// Fonction pour charger l'état de jeu depuis save.txt
func loadGame() (GameState, error) {
	file, err := os.Open("save.txt")
	if err != nil {
		return GameState{}, err
	}
	defer file.Close()

	var state GameState
	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		state.Word = scanner.Text()
	}
	if scanner.Scan() {
		state.MaskedWord = scanner.Text()
	}
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &state.RemainingAttempts)
	}
	return state, nil
}

// Lit un mot aléatoire dans le fichier words.txt
func lectureWord() (string, error) {
	file, err := os.Open("words.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return "", fmt.Errorf("le fichier words.txt est vide")
	}
	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex], nil
}

// Cache certaines lettres du mot
func motmaque(word string) []rune {
	runeWord := []rune(word)
	longueur := len(word)
	nbLettresMasquees := len(word)/2 - 1
	if longueur <= 3 {
		nbLettresMasquees = 1
	}

	for i := 1; i <= longueur-nbLettresMasquees; i++ {
		index := rand.Intn(len(word))
		if runeWord[index] == '_' {
			i--
			continue
		}
		runeWord[index] = '_'
	}
	return runeWord
}

// Affiche le pendu en fonction du nombre d'essais restants
func displayHangman(nbrEssai int) {
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier hangman.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineStart := (9 - nbrEssai) * 8
	lineEnd := lineStart + 8
	lineNumber := 0

	for scanner.Scan() {
		if lineNumber >= lineStart && lineNumber < lineEnd {
			fmt.Println(scanner.Text())
		}
		lineNumber++
	}
}

// Met à jour le mot masqué avec les lettres devinées
func updateMaskedWord(motmasque []rune, originalWord string, guess string) {
	for i := 0; i < len(originalWord); i++ {
		if string(originalWord[i]) == guess {
			motmasque[i] = rune(guess[0])
		}
	}
}

// Fonction principale du jeu
func game() {
	// Récupération d'un mot aléatoire
	word, err := lectureWord()
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	// Création du mot masqué
	motmasque := motmaque(word)

	// Initialisation des paramètres du jeu
	nbrEssai := 10

	// Boucle principale du jeu
	for nbrEssai > 0 {
		fmt.Println("Mot à deviner :", string(motmasque))
		displayHangman(nbrEssai)

		// Demande à l'utilisateur d'entrer une lettre
		fmt.Print("Entrez une lettre : ")
		var guess string
		fmt.Scan(&guess)

		// Mise à jour du mot masqué si la lettre est correcte
		updateMaskedWord(motmasque, word, guess)

		// Vérification de la victoire
		if string(motmasque) == word {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", word)
			return
		}

		// Réduction du nombre d'essais si la lettre n'est pas trouvée
		if !contienlettre(word, guess) {
			nbrEssai--
			fmt.Println("La lettre que vous avez écris n'est pas présente dans le mot, il vous reste :", nbrEssai, "essai")
		}
	}
	if nbrEssai == 0 {
		displayHangman(nbrEssai)
		fmt.Println("Dommage, vous avez perdu. Le mot était :", word)
	}

}

// Vérifie si le mot contient une lettre donnée
func contienlettre(word, letter string) bool {
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}
