package ascii

import (
	"io/ioutil"
	"strings"
)

// fonction pour ouvrir le fichier et tableau transformer a 2 entrée pour les lettres asciis
func Ascii(txt, namefile string) string {
	convert := ""
	file, _ := ioutil.ReadFile(namefile)
	file2 := strings.Split(string(file), "\n")
	arg := ChangArgs(txt)
	tab := SplitTab(FindTab(arg))

	// tab a 2 entrées
	for _, l := range tab { //execute printletter pour chaque ligne de mon tab
		convert += PrintLetterAscii(file2, l)
	}
	return convert
}

//permet d'afficher les lettres ascii à partir de la ligne calculé
func PrintLetterAscii(file2 []string, tab []int) string {
	convert := ""
	for i := -1; i < 7; i++ {
		for _, l := range tab {
			convert += string(file2[l+i])
			// fmt.Printf(file2[l+i]) //afficher tout les lignes de la lettre
		}
		convert += "\n"
		// fmt.Printf("\n")
	}
	return convert
}

//permet de calculer la ligne du début de la lettre
////permet de calculer la ligne du commencement de l'ascci-art à partir de la valeur ascii du caractère - '\n'
// prend la valeur 0 dans ce tableau

//c'est pour remplir ton tab avec
//les valeurs des premières lignes des caractères que tu veu afficher
func FindTab(arg []byte) []int {
	var tab []int
	tab = make([]int, len(arg)) // créer un tab vide de la taille voulu ici arg
	for i, l := range arg {     // meme pour tab
		if l == '\n' {
			tab[i] = 0
		} else {
			tab[i] = int(l-' ')*9 + 1 // determine la 1ere ligne de la lettre
		}
	}
	return tab
}

//tableau divisé en 2 entrées, pour toutes les valeurs 0
//alors c'est juste par rapport au sol de ligne. En gros si j'ai "hello\nyou"
//il faut d'abord affiché "hello" donc la première ligne du tableau et ensuite faire le "you"

func SplitTab(tab []int) [][]int {
	var board [][]int
	count := 1
	n := 0
	for _, l := range tab { //compte le nombre de 0 (\n)
		if l == 0 {
			count++
		}
	}
	board = make([][]int, count)
	for _, l := range tab {
		if l == 0 { // qd changer de ligne
			n++
		} else {
			board[n] = append(board[n], l) // rajouter des caracteres dans la ligne carac 1 apres l'autre
		}
	}
	return board
}

//changer 2 args en 1 seul pour que ce soit plus facile
func ChangArgs(txt string) []byte {
	arg := ""

	for i, l := range txt {
		if i < len(txt) && l == '\r' && txt[i+1] == '\n' {
			continue
		}
		if i > 0 && l == '\n' && txt[i-1] == '\r' {
			arg = arg + string('\n')
		} else {
			arg = arg + string(l)
		}
	}
	return []byte(arg)
}
