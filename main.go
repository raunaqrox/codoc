package main

import (
	"codoc/config"
	"codoc/downloader"
	"codoc/errors"
	"codoc/fs"
	"codoc/messages"
	"codoc/types"
	"codoc/utils"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	// first check if codoc is initialized
	// if the folder exists, config exists and flag is true
	// if all three false keep trying to do all 3 until they match

	// create directory at home folder
	// TODO make it platform agnostic
	fs.CreateDirectoryIfNotExists(filepath.Join(utils.GetHomeFolder(), config.Config["codocFolder"]))
}

func handleArgs(args []string) error {
	switch args[0] {
	case "get":
		if len(args) < 2 || args[1] == "" {
			return errors.ThrowArgumentError("Which documentation to get is not specified")
		}

		// check of the doc already exists
		// check if a more recent version exists
		// else download and store it

		doc, err := downloader.GetDoc(args[1])

		if err == nil {
			fmt.Printf(messages.Messages["downloadingDoc"], args[1])
			// TODO put in separate file
			fmt.Println(messages.Messages["successDocDownload"])
			saveDoc(doc)
			// write to file in correct structure
			return nil
		}

	case "del":
		fmt.Println("tes")
	case "list":
		fmt.Println("te")
	}
	return nil
}

func saveDoc(doc *types.Parsed) {
	res, err := json.Marshal(doc.ParsedDoc.Toc)
	_ = err
	docJsonPath := filepath.Join(utils.GetHomeFolder(), config.Config["codocFolder"], doc.DocInfo.DocName+".json")
	err = fs.WriteFile(docJsonPath, res)
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]
	err := handleArgs(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Config["homeFolder"])
}
