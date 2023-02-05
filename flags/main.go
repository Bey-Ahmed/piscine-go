package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsNumb := len(os.Args)

	if argsNumb == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("\t This flag inserts the string into the string passed as argument.")

		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
	} else if argsNumb > 1 {
		toInsert := ""
		givenString := ""
		toOrder := false
		argument := ""
		for i := 1; i < argsNumb; i++ {
			argument = os.Args[i]
			if len(argument) > 8 && argument[:8] == "--insert" {
				toInsert = argument[9:]
			} else if len(argument) > 2 && argument[:2] == "-i" {
				toInsert = argument[3:]
			} else if (len(argument) == 7 && argument == "--order") || (len(argument) == 2 && argument == "-o") {
				toOrder = true
			} else if len(argument) > 0 {
				givenString += argument
			}
		}

		givenString += toInsert
		stringSize := len(givenString)

		if toOrder {
			store := []rune(givenString)
			for i := 0; i < stringSize-1; i++ {
				for j := i + 1; j < stringSize; j++ {
					if store[i] > store[j] {
						store[i], store[j] = store[j], store[i]
					}
				}
			}
			givenString = string(store)
		}

		for k := 0; k < stringSize; k++ {
			z01.PrintRune(rune(givenString[k]))
		}
		z01.PrintRune('\n')
	}
}
