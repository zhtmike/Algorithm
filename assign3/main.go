package main

import (
	"flag"
	"fmt"
	"runtime"

	"../utils"
	"./src"
)

func main() {
	numcpu := flag.Int("cpu", runtime.NumCPU(), "Number of CPU")
	flag.Parse()
	runtime.GOMAXPROCS(*numcpu)
	arr := utils.ReadAdjListFromText("data/kargerMinCut.txt")
	ans := assign3.GetMinCut(arr)
	fmt.Println(ans)
}
