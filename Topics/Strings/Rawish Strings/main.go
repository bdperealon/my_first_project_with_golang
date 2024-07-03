package main

import "fmt"

func help() string {
	// Remember the difference between raw and normal strings?
	debug := `\n A newline character.` +
		`\t A tab character.` +
		`\" Double quotation marks.` +
		`\\ A backslash.`

	// DO NOT remove the 'return debug' statement!
	return debug
}

// DO NOT MODIFY the main function below!
func main() {
	fmt.Println(help())
}
