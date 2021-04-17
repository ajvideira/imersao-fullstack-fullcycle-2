package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)


type Position struct {
	Latitude 	float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Route struct {
	ID        string			`json:"routeId"`
	ClientID  string			`json:"clientId"`
	Positions []Position	`json:"positions"`
}

type PartialRoutePosition struct {
	ID 				string 		`json:"routeId"`
	ClientID 	string 		`json:"clientId"`
	Position 	[]float64 `json:"position"`
	Finished 	bool 			`json:"finished"`
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("route's ID not informed")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	for (scanner.Scan()) {
		data := strings.Split(scanner.Text(), ",")
		
		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		route.Positions = append(route.Positions, Position{
			Latitude: latitude,
			Longitude: longitude,
		})
	}

	return nil
}

func (route *Route) ExportJsonPositions() ([]string, error) {
	var jsonPositions []string
	
	total := len(route.Positions)

	for key, position := range route.Positions {
		partialRoutePosition := PartialRoutePosition{
			ID: route.ID,
			ClientID: route.ClientID,
			Position: []float64{position.Latitude, position.Longitude},
			Finished: key == total - 1,
		}

		jsonPartialRoutePosition, err := json.Marshal(partialRoutePosition)
		if err != nil {
			return nil, err
		}
		jsonPositions = append(jsonPositions, string(jsonPartialRoutePosition))
	}

	return jsonPositions, nil
}