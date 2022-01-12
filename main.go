package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("a", "none", "idk")
	config := flag.String("c", "Test", "idk")

	flag.Parse()

	/*fmt.Println("action: ", *action)
	fmt.Println("config: ", *config)*/

	if *action == "fill" {
		FillForm(ReadConfig(*config))
		fmt.Println("filled form")

	} else if *action == "gen" {
		GenConfig(ReadConfigGen())
		fmt.Println("created Config")

	} else {
		PrintGuide()
		return
	}
}

func PrintGuide() {
	fmt.Println("")
	fmt.Println("flags:")
	fmt.Println("-a=gen")
	fmt.Println("-a=fill")
	fmt.Println("-c={ConfigName}")
	fmt.Println("")
	fmt.Println("default config is Test")
	fmt.Println("")
}
