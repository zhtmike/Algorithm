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
	arr0, arr1 := fileio.ReadJobListFromText(os.Args[1])
	fmt.Println(arr0)
	fmt.Println(len(arr1))
}
