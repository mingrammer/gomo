package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mingrammer/gomo/manager"
)

const usage = "Usage: gomo [command] [args]"

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println(usage)
		os.Exit(-1)
	}

	mgr := manager.New()

	mgr.AddCommand(initCommand)
	mgr.AddCommand(newCommand)
	mgr.AddCommand(listCommand)
	mgr.AddCommand(deleteCommand)

	args := flag.Args()

	if err := mgr.Execute(args); err != nil {
		fmt.Println(err)
		fmt.Print(mgr.Usage())
	}
}
