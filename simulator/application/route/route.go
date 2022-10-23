package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

var LATITUDE = 0
var LONGITUDE = 1

type Route struct {
	ID        string     `json: "routeId"`
	ClientID  string     `json: "clientId"`
	Positions []Position `json: "positions"`
}

type Position struct {
	Latitude  float64
	Longitude float64
}

type PartialRoutePosition struct {
	ID        string    `json: "routeId"`
	ClientID  string    `json: clientId`
	Positions []float64 `json: positions`
	Finished  bool      `json: finished`
}

func (route *Route) LoadPositions() error {

	if route.ID == "" {
		return errors.New("Route ID is not informed")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		position := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(position[LATITUDE], 64)

		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(position[LONGITUDE], 64)

		if err != nil {
			return nil
		}

		route.Positions = append(route.Positions, Position{
			Latitude:  lat,
			Longitude: long,
		})

	}
	return nil

}

func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for index, pos := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Positions = []float64{pos.Latitude, pos.Longitude}
		route.Finished = false
		if total-1 == index {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))

	}
	return result, nil
}
