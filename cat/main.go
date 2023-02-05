package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printError(erreur string) {
	erreur = "ERROR: " + erreur
	for _, char := range erreur {
		z01.PrintRune(rune(char))
	}
	z01.PrintRune('\n')
}

func main() {
	size := len(os.Args)

	if size < 2 {
		line, errStdin := io.ReadAll(os.Stdin)
		if errStdin != nil {
			printError(errStdin.Error())
			os.Exit(1)
		}
		printStr(string(line))
		return
	}

	for i := 1; i < size; i++ {
		filename := os.Args[i]
		file, errfile := os.Open(filename)
		if errfile != nil {
			printError(errfile.Error())
			os.Exit(1)
			return
		}
		info, _ := file.Stat()
		input := make([]byte, info.Size())
		file.Read(input)
		printStr(string(input))
		file.Close()
	}
}
