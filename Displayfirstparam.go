package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	str := os.Args                //l'argument de la ligne de commande est transformé en tableau de string séparé par des quotes
	for _, char := range str[1] { //pour n'importe quelle position des valeurs char de notre string en position 1 dans le tableau str
		//la ligne du haut a transfomé notre string 1, premier paramètre, en runes appelées char que notre z01 peut imprimer
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
