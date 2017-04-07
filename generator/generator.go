package generator

import "fmt"

type generator func() error

var generators = map[string]generator{
	"go":         generateGo,
	"python":     generatePython,
	"ruby":       generateRuby,
	"javascript": generateJS,
	"rust":       generateRust,
}

func Generate(args ...string) error {
	language := args[0]
	return generators[language]()
}

func generateGo() error {
	fmt.Println("TODO: Implement go generator")
	return nil
}

func generatePython() error {
	fmt.Println("TODO: Implement python generator")
	return nil
}

func generateRuby() error {
	fmt.Println("TODO: Implement ruby generator")
	return nil
}

func generateJS() error {
	fmt.Println("TODO: Implement js generator")
	return nil
}

func generateRust() error {
	fmt.Println("TODO: Implement rust generator")
	return nil
}
