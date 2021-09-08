package maps

import (
	"fmt"
)

type GraphicInterface struct {
	Graphic [][]rune
	Dimensions Dimension
}

func NewGraphicsInterface (dimensions Dimension, backgroundChar rune)(GraphicInterface, error){
	if dimensions.Width <= 0 {
		return GraphicInterface{}, fmt.Errorf("the graphicInterface width must be greater than 0")
	}
	if dimensions.Heigth <= 0 {
		return GraphicInterface{}, fmt.Errorf("the graphicInterface heigth must be greater than 0")
	}
	var newGraphicsInterface GraphicInterface
	newGraphicsInterface.Dimensions = dimensions
	newGraphicsInterface.Graphic = make([][]rune, dimensions.Heigth)
	for rowIndex := range newGraphicsInterface.Graphic {
		newGraphicsInterface.Graphic[rowIndex] = make([]rune, dimensions.Width)
		for columnIndex := range newGraphicsInterface.Graphic[rowIndex] {
			newGraphicsInterface.Graphic[rowIndex][columnIndex] = backgroundChar
		}
	}
	return newGraphicsInterface, nil
}

func (graphicInterface GraphicInterface) Contains (rune rune)(bool){
	for _, row := range graphicInterface.Graphic {
		for _, char := range row {
			if char == rune {
				return true
			}
		}
	}
	return false
}

func (graphicInterface GraphicInterface) Print (){
	for _, row := range graphicInterface.Graphic {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}