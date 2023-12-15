package utils

import (
	"math"

	"kayn.ooo/go-art/app"
)

func GetPositionFromDegree(v app.Vector, degrees float64, distance int) *app.Vector {
	return &app.Vector{
		X: (math.Ceil(float64(v.X) + float64(distance)*math.Sin((float64(degrees)*math.Pi)/180))),
		Y: (math.Ceil(float64(v.Y) + float64(distance)*math.Cos((float64(degrees)*math.Pi)/180))),
	}
}
