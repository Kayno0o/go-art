package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Schema struct {
	Shapes        [][]Vector `json:"shapes"`
	RectangleSize Vector     `json:"rectangleSize"`
	CanvasSize    Vector     `json:"canvasSize"`
	MultVector    Vector     `json:"multVector"`
	LineWidth     float64    `json:"lineWidth"`
}

func drawLines(dc *gg.Context, vectors []Vector) {
	if len(vectors) < 2 {
		return // If there are less than two vectors, no line can be drawn
	}
	dc.MoveTo(vectors[0].X, vectors[0].Y)
	for i := 1; i < len(vectors); i++ {
		dc.LineTo(vectors[i].X, vectors[i].Y)
	}
	dc.SetRGB(0, 0, 0)
	dc.Stroke()
}

func drawShapes(dc *gg.Context, shapes [][]Vector) {
	for i := 0; i < len(shapes); i++ {
		drawLines(dc, shapes[i])
	}
}

func moveVectors(vectors []Vector, displacement Vector) {
	for i := range vectors {
		vectors[i].X += displacement.X
		vectors[i].Y += displacement.Y
	}
}

func moveShapes(shapes [][]Vector, displacement Vector) {
	for i := 0; i < len(shapes); i++ {
		moveVectors(shapes[i], displacement)
	}
}

func multVectors(vectors []Vector, factor Vector) {
	for i := range vectors {
		vectors[i].X *= factor.X
		vectors[i].Y *= factor.Y
	}
}

func multShapes(shapes [][]Vector, factor Vector) {
	for i := 0; i < len(shapes); i++ {
		multVectors(shapes[i], factor)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No .json file provided")
		return
	}

	schema := Schema{}
	filename := ""

	// Iterate over the command-line arguments
	for _, arg := range os.Args[1:] {
		// If the argument ends with ".json", read the file
		if strings.HasSuffix(arg, ".json") {
			filename = arg
			// Read the file
			jsonData, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Unmarshal the JSON data into the shapes variable
			err = json.Unmarshal(jsonData, &schema)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Stop after the first .json file
			break
		}
	}

	rectangleSize := schema.RectangleSize
	canvasSize := schema.CanvasSize
	shapes := schema.Shapes
	multVector := schema.MultVector
	lineWidth := float64(1)

	if schema.LineWidth != 0 {
		lineWidth = schema.LineWidth
	}

	dc := gg.NewContext(int(canvasSize.X), int(canvasSize.Y))

	multShapes(shapes, multVector)
	rectangleSize.X *= multVector.X
	rectangleSize.Y *= multVector.Y

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetLineWidth(lineWidth)

	start := Vector{
		X: -rectangleSize.X,
		Y: -rectangleSize.Y,
	}
	end := Vector{
		X: canvasSize.X + rectangleSize.X,
		Y: canvasSize.Y + rectangleSize.Y,
	}

	for y := int(-start.Y); y < int(end.Y); y += int(rectangleSize.Y) {
		for x := int(-start.X); x < int(end.X); x += int(rectangleSize.X) {
			drawShapes(dc, shapes)
			moveShapes(shapes, *&Vector{X: rectangleSize.X, Y: 0})
		}
		moveShapes(shapes, *&Vector{X: -end.X - start.X, Y: rectangleSize.Y})
	}

	dc.SavePNG(filename + ".png")
}
