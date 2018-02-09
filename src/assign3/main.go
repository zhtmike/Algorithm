package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"utils/fileio"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file.")
	}
	numcpu := flag.Int("n", runtime.NumCPU(), "Number of processes")
	flag.Parse()
	runtime.GOMAXPROCS(*numcpu)
	arr := fileio.ReadAdjListFromText(os.Args[1])
	ans := GetMinCut(arr)
	fmt.Println(ans)
}
