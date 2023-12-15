package app

import "github.com/fogleman/gg"

type Canvas struct {
	Context       *gg.Context
	Shapes        []Shape `json:"shapes"`
	RectangleSize Vector  `json:"rectangleSize"`
	CanvasSize    Vector  `json:"canvasSize"`
	MultVector    Vector  `json:"multVector"`
	LineWidth     float64 `json:"lineWidth"`
}

func (c *Canvas) Init() {
	c.Context = gg.NewContext(int(c.CanvasSize.X), int(c.CanvasSize.Y))
}

func (c *Canvas) Draw() {
	lineWidth := float64(1)

	if c.LineWidth != 0 {
		lineWidth = c.LineWidth
	}

	c.multShapes(c.MultVector)
	c.RectangleSize.X *= c.MultVector.X
	c.RectangleSize.Y *= c.MultVector.Y

	c.Context.SetRGB(1, 1, 1)
	c.Context.Clear()
	c.Context.SetLineWidth(lineWidth)

	start := Vector{
		X: -c.RectangleSize.X,
		Y: -c.RectangleSize.Y,
	}
	end := Vector{
		X: c.CanvasSize.X + c.RectangleSize.X,
		Y: c.CanvasSize.Y + c.RectangleSize.Y,
	}

	for y := int(-start.Y); y < int(end.Y); y += int(c.RectangleSize.Y) {
		for x := int(-start.X); x < int(end.X); x += int(c.RectangleSize.X) {
			c.drawShapes()
			c.moveShapes(Vector{X: c.RectangleSize.X, Y: 0})
		}
		c.moveShapes(Vector{X: -end.X - start.X, Y: c.RectangleSize.Y})
	}

}

func (c *Canvas) Save(filename string) {
	if filename != "" {
		c.Context.SavePNG(filename + ".png")
	}
}

func (c *Canvas) drawShape(s Shape) {
	c.Context.MoveTo(s.Vectors[0].X, s.Vectors[0].Y)
	for i := 1; i < len(s.Vectors); i++ {
		c.Context.LineTo(s.Vectors[i].X, s.Vectors[i].Y)
	}
	c.Context.SetRGB(0, 0, 0)
	c.Context.Stroke()
}

func (c *Canvas) drawShapes() {
	for i := 0; i < len(c.Shapes); i++ {
		c.drawShape(c.Shapes[i])
	}
}

func (c *Canvas) moveShapes(factor Vector) {
	for i := 0; i < len(c.Shapes); i++ {
		c.Shapes[i].Move(factor)
	}
}

func (c *Canvas) multShapes(factor Vector) {
	for i := 0; i < len(c.Shapes); i++ {
		c.Shapes[i].Mult(factor)
	}
}
