package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scripts []Script

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Execute a available feature. Developer: Dominik48N")
	fmt.Printf("Scripts: %v\n", scripts)
	fmt.Println("--------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("scripts", text) == 0 {
			fmt.Printf("Scripts: %v\n", scripts)
		}
	}
}

type Script interface {
}
