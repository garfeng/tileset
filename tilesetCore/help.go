package main

import "fmt"

func printHelp() {
	fmt.Printf(
		"tilesetCore.exe %s=<input> %s=<output> [%s=<cpu|gpu>] [%s=<true|false>]\n",
		INPUT, OUTPUT, CORE, MODIFY_HUI)
	fmt.Println("run \"tilesetCore.exe -h|--help\" for more information")
}

func printHelpInDetail() {
	printHelp()
	fmt.Println("\n-----------------\n")
	fmt.Println(INPUT, ": input, png file or directory")
	fmt.Println(OUTPUT, ": output, png file or directory")
	fmt.Println(CORE, ": processor, must be cpu or gpu, cpu default")
	fmt.Println(MODIFY_HUI, ": whether modify the Hue and light, set to true when you convert the xp tileset, false default")
}
