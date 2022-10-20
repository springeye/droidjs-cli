package client

import "fmt"

func SetupClientListAction(args []string, options map[string]string) int {
	fmt.Printf("%v\n", args)
	fmt.Printf("%v\n", options)
	return 0
}
