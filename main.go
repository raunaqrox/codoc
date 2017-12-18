package main

import (
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
	if exists, _ := utils.FolderExists(utils.GetCodocFolder()); !exists {
		fs.CreateDirectoryIfNotExists(utils.GetCodocFolder())
	}

}

func installDoc(docName string) {
	doc, err := downloader.GetDoc(docName)

	if err == nil {
		fmt.Printf(messages.Messages["downloadingDoc"], docName)
		// TODO put in separate file
		fmt.Println(messages.Messages["successDocDownload"])
		if doc != nil {
			// saveDoc(doc)
		}
	}
}

func handleArgs(args []string) error {
	if !(len(args) == 0) {
		switch args[0] {
		case "get":
			if len(args) < 2 || args[1] == "" {
				return errors.ThrowArgumentError("Which documentation to get is not specified")
			}

			// check of the doc already exists
			// check if a more recent version exists
			// else download and store it
			if utils.IsDocInstalled(args[1]) {
				fmt.Println("Node.js is already installed!")
				return nil // TODO create codoc error
			} else if utils.IsAllowedDoc(args[1]) {
				installDoc(args[1])
			} else {
				fmt.Println(args[1], " not allowed")
			}

		case "del":
			fmt.Println("tes")
		case "list":
			docs := utils.GetDocList(utils.GetCodocFolder())
			if len(docs) == 0 {
				fmt.Print("No doc installed")
				return nil // TODO create a codoc error and return that here
			}

			for _, doc := range docs {
				fmt.Println(doc)
			}
		default:
			// TODO: turn this into a single fn, which get's the allowed
			// docs from the installed docs folder for now
			if utils.IsAllowedDoc(args[0]) && utils.IsDocInstalled(args[0]) {
				// list the table of contents of nodejs
				docOutput, err := utils.GetDoc(args[0])
				_ = docOutput
				if err != nil {
					panic(err)
				}
				// for _, toc := range docOutput.Toc {
				// 	fmt.Printf("%s\n", toc.Name)
				// }

			}
		}

	}
	return nil
}

func saveDoc(doc *types.Parsed) {
	res, err := json.Marshal(doc.ParsedDoc.Toc)
	_ = err
	docJsonPath := filepath.Join(utils.GetCodocFolder(), doc.DocInfo.DocName+".json")
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
}
