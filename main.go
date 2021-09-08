package awayhome

import (
	"github.com/MoraGames/awayhome/internal/pkg/maps"
	"log"
)

func main (){
	graphicInterface, err := maps.NewGraphicsInterface(maps.StandardDimensions["extra-large"], '#')
	if err != nil {
		log.Panicln(err)
	}

	FirstMap, err := maps.NewMap("First Map", graphicInterface)
	if err != nil {
		log.Panicln(err)
	}
	FirstMap.GraphicInterface.Print()
}