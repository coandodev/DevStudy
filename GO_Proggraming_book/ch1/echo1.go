// Echo1 prints its command-line arguments.package GO_Proggraming_book

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[0]
		sep = " "
	}
	fmt.Println(s)
}
