package maps

import (
	"fmt"
)

type Map struct {
	name string
	GraphicInterface GraphicInterface
}

func NewMap (name string, graphicInterface GraphicInterface)(Map, error){
	if graphicInterface.Dimensions.Width <= 0 {
		return Map{}, fmt.Errorf("the map width must be greater than 0")
	}
	if graphicInterface.Dimensions.Heigth <= 0 {
		return Map{}, fmt.Errorf("the map heigth must be greater than 0")
	}
	if graphicInterface.Contains(0) == true {
		return Map{}, fmt.Errorf("the graphicInterface of the map cannot contain nil runes ('')")
	}
	return Map{name, graphicInterface}, nil
}