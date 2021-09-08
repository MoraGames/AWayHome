package maps

import "fmt"

var StandardDimensions = map[string]Dimension {
	"extra-small": {15,10},
	"small": {20,15},
	"medium": {25,20},
	"large": {30,25},
	"extra-large": {40,35},
}

type Dimension struct {
	Width int
	Heigth int
}

func NewDimension (width, heigth int)(Dimension, error){
	if width <= 0 {
		return Dimension{}, fmt.Errorf("width must be greater than 0")
	}
	if heigth <= 0 {
		return Dimension{}, fmt.Errorf("heigth must be greater than 0")
	}
	return Dimension{width, heigth}, nil
}