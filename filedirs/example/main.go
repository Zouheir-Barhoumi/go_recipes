package main

import "github.com/Zouheir-Barhoumi/go_recipes/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}
	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}
