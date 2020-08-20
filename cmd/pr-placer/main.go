package main

import (
	"github.com/manifoldco/promptui"
	"github.com/tcnksm/go-input"
	"log"
	"os"
)

type option struct {
	key   string
	value string
}

func selectFromOptions(question string, options []option) string {
	optionValues := make([]string, len(options))
	for i, opt := range options {
		optionValues[i] = opt.value
	}

	prompt := promptui.Select{
		Label: question,
		Items: optionValues,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return options[i].key
}

func getChangeType() string {
	changeTypes := []option{
		option{
			key:   "external",
			value: "Users (bugs, features, performance)",
		},
		option{
			key:   "internal",
			value: "Just us (docs, ci, code style)",
		},
	}

	return selectFromOptions("Who can notice the change?", changeTypes)
}

func getExternalChangeType() string {
	changeTypes := []option{
		option{
			key:   "feat",
			value: "A new feature",
		},
		option{
			key:   "fix",
			value: "A bug fix",
		},
		option{
			key:   "perf",
			value: "A performance improvement",
		},
		option{
			key:   "docs",
			value: "External, documentation only changes (e.g. API docs)",
		},
	}

	return selectFromOptions("What sort of change is it?", changeTypes)
}

func getInternalChangeType() string {
	changeTypes := []option{
		option{
			key:   "ci",
			value: "Changes purely to our CI configuration files and scripts",
		},
		option{
			key:   "build",
			value: "Changes that affect how we compile or execute code",
		},
		option{
			key:   "docs",
			value: "Documentation only changes",
		},
		option{
			key:   "refactor",
			value: "Refactoring code without changing what it does",
		},
		option{
			key:   "style",
			value: "Code style changes (e.g. new linter rule)",
		},
		option{
			key:   "test",
			value: "Adding missing tests or correcting existing tests",
		},
	}

	return selectFromOptions("What sort of change is it?", changeTypes)
}

func getTitle() string {
	ui := &input.UI{}
	query := "What's the one-sentence title of the change?"
	name, err := ui.Ask(query, &input.Options{
		HideOrder: true,
		Loop:      true,
		Required:  true,
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return name
}

func main() {
	changeType := getChangeType()

	var conventionalType string

	if changeType == "external" {
		conventionalType = getExternalChangeType()
	} else {
		conventionalType = getInternalChangeType()
	}

	log.Printf("%s\n", conventionalType)

	getTitle()
}
