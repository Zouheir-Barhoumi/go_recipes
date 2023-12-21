package main

import "github.com/Zouheir-Barhoumi/go_recipes/templates"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
