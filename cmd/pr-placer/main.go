package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type work struct {
	Description string
	Key         string
}

func main() {
	prompt := promptui.Select{
		Label: "Who can notice the change?",
		Items: []string{
			"Users (bugs, features, performance)",
			"Just us (docs, ci, code style)",
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %d\n", i)
}
