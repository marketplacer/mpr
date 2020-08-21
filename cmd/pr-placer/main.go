package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
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
	title := ""
	prompt := &survey.Input{
		Message: "What's the one-sentence title of the change?",
	}
	survey.AskOne(prompt, &title, survey.WithValidator(survey.Required))

	return title
}

func getDescription() string {
	text := ""
	prompt := &survey.Multiline{
		Message: "Now some detail. What's the change do? Why is it being done?",
	}
	survey.AskOne(prompt, &text)

	return text
}

func checkQaRequired() bool {
	prompt := promptui.Select{
		Label: "Does this work require QA?",
		Items: []string{
			"No",
			"Yes",
		},
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return i == 1
}

func getReproductionSteps() []string {
	fmt.Println("What are the steps to see this working? Submit a blank step to finish.")

	var steps []string
	i := 0

	for {
		step := ""
		prompt := &survey.Input{
			Message: fmt.Sprintf("Step %d", i+1),
		}
		survey.AskOne(prompt, &step)
		if step == "" {
			break
		}

		i++
		steps = append(steps, step)
	}

	return steps
}

func getEnvironmentURL() string {
	url := ""
	prompt := &survey.Input{
		Message: "URL where the change can be tested",
	}
	survey.AskOne(prompt, &url)

	return url
}

func getResolvedTickets() []string {
	fmt.Println("Paste any ticket URLs that this resolves (FULL URLs)")

	var urls []string
	i := 0

	for {
		url := ""
		prompt := &survey.Input{
			Message: "Ticket URL",
		}
		survey.AskOne(prompt, &url)
		if url == "" {
			break
		}

		i++
		urls = append(urls, url)
	}

	return urls
}

func formatPr(conventionalType string, title string, description string, reproductionSteps []string, url string, ticketUrls []string) {
	fmt.Println("Okay, go to the page where you enter your PR details, then press enter here")
	fmt.Scanln()

	fmt.Printf("Here's your PR title. Copy it over, then press enter here\n\n")
	fmt.Printf("%s: %s\n\n", conventionalType, title)
	fmt.Scanln()

	fmt.Printf("Here's your PR body. Copy it over, add any extra detail you'd like over in GitHub\n\n")
	fmt.Printf("%s\n\n", description)

	if len(reproductionSteps) > 0 {

		fmt.Printf("**Steps to test this is working**\n\n")
		for i, step := range reproductionSteps {
			fmt.Printf("%d. %s\n", i+1, step)
		}

		fmt.Println()
	}

	fmt.Printf("Deployed to: %s\n", url)

	if len(ticketUrls) > 0 {
		for _, ticketURL := range ticketUrls {
			fmt.Printf("Resolves %s\n", ticketURL)
		}

		fmt.Println()
	}
}

func main() {
	changeType := getChangeType()

	var conventionalType string

	if changeType == "external" {
		conventionalType = getExternalChangeType()
	} else {
		conventionalType = getInternalChangeType()
	}

	title := getTitle()
	description := getDescription()

	qaRequired := true
	if conventionalType != "feat" && conventionalType != "fix" {
		qaRequired = checkQaRequired()
	}

	reproductionSteps := []string{}
	var url string

	if qaRequired {
		reproductionSteps = getReproductionSteps()
		url = getEnvironmentURL()
	}

	ticketUrls := getResolvedTickets()

	log.Printf("Type: %s\n", conventionalType)
	log.Printf("%v", reproductionSteps)
	log.Printf("%v", url)
	log.Printf("%v", ticketUrls)

	formatPr(conventionalType, title, description, reproductionSteps, url, ticketUrls)
}
