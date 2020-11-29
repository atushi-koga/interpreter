package main

import (
	"fmt"
	"github.com/atushi-koga/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	appUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", appUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
