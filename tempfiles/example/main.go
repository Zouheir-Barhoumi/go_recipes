package main

import "github.com/Zouheir-Barhoumi/go_recipes/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
