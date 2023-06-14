package main

import (
	"fmt"
	"os"
	"os/user"
	"terminal"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Shelgi programming language!\n",
		user.Username)
	terminal.Start(os.Stdin, os.Stdout)
}
