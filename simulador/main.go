package main

import (
	"fmt"

	"github.com/raul-marques/imersaofsfc2-simulator/application/route"
)

func main() {
	routeVar := route.Route{ID: "1", ClientID: "1"}
	routeVar.LoadPositions()
	stringjson, _ := routeVar.ExportJsonPositions()
	fmt.Println(stringjson[1])
}
