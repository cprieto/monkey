package main

import (
	"fmt"
	"github.com/cprieto/monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s to the Monkey REPL!\n", user.Username)
	fmt.Println("Feel free to type in commands.")
	repl.Start(os.Stdin, os.Stdout)
}
