package main

import (
	"bufio"
	"fmt"
	"github.com/iancoleman/strcase"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fmt.Println(strcase.ToSnake(text))
	fmt.Println(strcase.ToCamel(text))
}
