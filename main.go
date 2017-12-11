package main

import (
	"codoc/config"
	"codoc/downloader"
	"codoc/errors"
	"codoc/messages"
	"fmt"
	"os"
)

func handleArgs(args []string) error {
	switch args[0] {
	case "get":
		if len(args) < 2 || args[1] == "" {
			return errors.ThrowArgumentError("Which documentation to get is not specified")
		}

		err := downloader.GetDoc(args[1])

		if err == nil {
			fmt.Printf(messages.Messages["downloadingDoc"], args[1])
			// put in separate file
			fmt.Println(messages.Messages["successDocDownload"])
			return nil
		}

	case "del":
		fmt.Println("tes")
	case "list":
		fmt.Println("te")
	}
	return nil
}

func main() {
	args := os.Args[1:]
	err := handleArgs(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Config["homeFolder"])
}
