package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Priyanka488/ZenScript/lexer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("zen >> ")
		if scanner.Scan() {
			input := scanner.Text()
			generatedTokens, err := lexer.Run(input, "stdin")
			if err != nil {
				fmt.Println(err)
			} else {
				for _, token := range generatedTokens {
					fmt.Printf("Type: %s Value: %s\n", token.Ttype, token.Value)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
