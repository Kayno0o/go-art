package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"

	"kayn.ooo/go-art/app"
)

func drawRandomSchema() {
	shapes := []app.Shape{}
	rectangle := app.Vector{
		X: 0,
		Y: 0,
	}

	shapesNb := rand.Intn(5) + 1
	for i := 0; i < shapesNb; i++ {
		vectorsNb := rand.Intn(5) + 2

		vectors := []app.Vector{}
		for v := 0; v < vectorsNb; v++ {
			vector := app.Vector{
				X: float64(rand.Intn(10) * 10),
				Y: float64(rand.Intn(10) * 10),
			}
			vectors = append(vectors, vector)
			if vector.X > rectangle.X {
				rectangle.X = vector.X
			}
			if vector.Y > rectangle.Y {
				rectangle.Y = vector.Y
			}
		}

		shapes = append(shapes, app.Shape{Vectors: vectors})
	}

	rectangle = app.Vector{
		X: rectangle.X - 10,
		Y: rectangle.Y - 10,
	}

	canvas := app.Canvas{
		Shapes:        shapes,
		RectangleSize: rectangle,
		MultVector: app.Vector{
			X: 1,
			Y: 1,
		},
		CanvasSize: app.Vector{
			X: rectangle.X * 10,
			Y: rectangle.Y * 10,
		},
		LineWidth: 3,
	}

	canvas.Init()
	canvas.Draw()
	canvas.Save("src/random")

	jsonData, err := json.MarshalIndent(canvas, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile("src/random.json", jsonData, 0766)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No .json file provided")
		return
	}

	args := os.Args[1:]

	firstArg := args[0]
	args = args[1:]

	switch firstArg {
	case "draw":
		for _, arg := range args {
			// If the argument ends with ".json", read the file
			if strings.HasSuffix(arg, ".json") {
				filename := arg
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
	case "pi":
		fmt.Print(math.Pi)
	}
}
