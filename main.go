package main

import (
	"flag"
	"fmt"
	"github.com/Aur0ra/core"
	"os"
)

func main() {
	options := core.Options{}

	flag.StringVar(&options.Domain, "D", "", "Detect Main-Domain")
	flag.StringVar(&options.DictPath, "d", "names.txt", "Path to dict-> Default ./names.txt")
	flag.StringVar(&options.FingerPath, "f", "fingers.json", "Path to fingers(.json)->Default ./fingers.json")
	flag.IntVar(&options.Thread, "t", 5, "Number of concurrent threads->Default 5")
	flag.IntVar(&options.Timeout, "timeout", 2, "Timeout for request-> Default 60")

	flag.Parse()

	flag.Usage = func() {
		fmt.Println("Usage of Subtaker")
		flag.PrintDefaults()
	}

	if flag.NFlag() == 0 || options.Domain == "0" {
		flag.Usage()
		os.Exit(1)
	}

	core.Drive(&options)

}
