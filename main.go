package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("goshell % ")
		text, err := reader.ReadString('\n')
		if err != nil {
			printError("error occured when reading input", err)
		}
		
		text = strings.Replace(text, "\n", "", 1)
		if strings.Compare("exit", text) == 0 {
			break
		}
	}
}

func printError(message string, err error) {
	fmt.Println(message)
	fmt.Println(err.Error())
}