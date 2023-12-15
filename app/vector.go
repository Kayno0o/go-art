package app

type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v *Vector) Move(displacement Vector) {
	v.X += displacement.X
	v.Y += displacement.Y
}

func (v *Vector) Mult(displacement Vector) {
	v.X *= displacement.X
	v.Y *= displacement.Y
}
