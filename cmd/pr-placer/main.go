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
		Label: "What type of work is this?",
		Items: []string{
			"Something that end users can experience (e.g. a visual change to an interface, or a new field in an API)",
			"Something internal (documentation, refactoring, linting, CI, performance improvements etc.)",
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %d\n", i)
}
