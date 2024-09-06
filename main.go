package main

import (
	"fmt"
	"os"
	"os/user"
	"rhydb/repl"
)

func main(){
	user,err:=user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sup %s Welcome to rhydb\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}