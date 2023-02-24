package main

import (
	"fmt"
	"strings"

	"github.com/hobbymarks/go-difflib/difflib"
)

func main() {
	a := "This is a File"
	b := "this book, not File"
	seqm := difflib.NewMatcher(strings.Split(a, " "), strings.Split(b, " "))
	for _, opc := range seqm.GetOpCodes() {
		switch opc.Tag {
		case 'r':
			fmt.Println("r:", opc)
		case 'd':
			fmt.Println("d:", opc)
		case 'i':
			fmt.Println("i:", opc)
		case 'e':
			fmt.Println("e:", opc)
		}
	}

}
