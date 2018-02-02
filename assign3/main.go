package main

import (
	"flag"
	"fmt"
	"runtime"

	"../utils"
	"./src"
)

func main() {
	numcpu := flag.Int("n", runtime.NumCPU(), "Number of processes")
	flag.Parse()
	runtime.GOMAXPROCS(*numcpu)
	arr := utils.ReadAdjListFromText("data/kargerMinCut.txt")
	ans := assign3.GetMinCut(arr)
	fmt.Println(ans)
}
