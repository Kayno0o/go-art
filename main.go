package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"kayn.ooo/go-art/app"
)

func drawRandomSchema() {

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No .json file provided")
		return
	}

	filename := ""

	// Iterate over the command-line arguments
	for _, arg := range os.Args[1:] {
		// If the argument ends with ".json", read the file
		if strings.HasSuffix(arg, ".json") {
			filename = arg
			// Read the file
			jsonData, err := os.ReadFile(arg)
			if err != nil {
				fmt.Println(err)
				return
			}

			canvas := app.Canvas{}

			// Unmarshal the JSON data into the shapes variable
			err = json.Unmarshal(jsonData, &canvas)
			if err != nil {
				fmt.Println(err)
				return
			}

			canvas.Init()
			canvas.Draw()
			canvas.Save(filename)
			continue
		}

		if arg == "random" {
			drawRandomSchema()
		}
	}
}
