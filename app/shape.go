package app

type Shape struct {
	Vectors []Vector `json:"vectors"`
}

func (s *Shape) Move(displacement Vector) {
	for i := range s.Vectors {
		s.Vectors[i].Move(displacement)
	}
}

func (s *Shape) Mult(factor Vector) {
	for i := range s.Vectors {
		s.Vectors[i].Mult(factor)
	}
}
