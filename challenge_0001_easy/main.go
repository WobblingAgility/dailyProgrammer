package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name := reader.ReadString('\n')
	fmt.Print("How old are you? ")
	age := reader.ReadString('\n')
	fmt.Print("And what is your reddit username? ")
	un := reader.ReadString('\n')
	fmt.Println(name, age, un)
}
