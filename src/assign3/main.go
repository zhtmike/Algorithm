package main

import (
	"fmt"
	"os"
	"runtime"
	"utils/fileio"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file.")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	arr := fileio.ReadAdjListFromText(os.Args[1])
	ans := GetMinCut(arr, runtime.NumCPU())
	fmt.Println(ans)
}
