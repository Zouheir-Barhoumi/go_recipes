package main

import (
	"github.com/Zouheir-Barhoumi/go_recipes/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		panic(err)
	}
	// if err := csvformat.WriteCSVOutput(); err != nil {
	// 	panic(err)
	// }
	// buffer, err := csvformat.WriteCSVBuffer()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Buffer = ", buffer.String())
}