package main

import (
	"fmt"
	"os"

	"utils/fileio"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file.")
	}
	arr := fileio.ReadIntArrayFromText(os.Args[1])
	fmt.Println(InverseCount(arr))
}
