package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/jmcvetta/neoism"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level} %{id}%{color:reset} %{message}`,
)

func createTownFromJSON(fileName string) error {
	backend1Leveled := logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	backend1Leveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backend1Leveled, logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), format))

	if _, err := DB.CreateUniqueConstraint("BusLine", "name"); err != nil {
		return err
	}
	if _, err := DB.CreateUniqueConstraint("BusStop", "id"); err != nil {
		return err
	}
	town, err := getTown(fileName)
	if err != nil {
		return err
	}
	for _, busLine := range town.BusLines {
		log.Noticef("Start registering line %s\n", busLine.Name)
		if err := createBusLine(busLine); err != nil {
			return err
		}
		if err := createBusStops(busLine); err != nil {
			return err
		}
	}
	return nil
}

func getTown(fileName string) (*jsonData, error) {
	var out jsonData
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(content, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func createBusLine(busLine *busLine) error {
	query := `MERGE (b:BusLine {name: {name}, id: {id}, color: {color}, textColor: {textColor}})
	ON CREATE
	SET b:BusLine`
	params := neoism.Props{
		"name":      busLine.Name,
		"id":        busLine.ID,
		"color":     busLine.Color,
		"textColor": busLine.TextColor,
	}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: params,
	}
	if err := DB.Cypher(&cq); err != nil {
		return err
	}
	log.Noticef("Line %s created\n", busLine.Name)
	return nil
}

func createBusStops(busLine *busLine) error {
	for dir := 0; dir < len(busLine.Path); dir++ {
		if busLine.Path[dir] == nil {
			continue
		}
		log.Debugf("Start path %s (%d (on %d))\n", busLine.Path[dir].Dir, dir, len(busLine.Path))
		for idStop := 0; idStop < len(busLine.Path[dir].Stops); idStop++ {
			stop := busLine.Path[dir].Stops[idStop]
			if err := createStop(stop); err != nil {
				return err
			}
			if idStop+1 < len(busLine.Path[dir].Stops) {
				stopNext := busLine.Path[dir].Stops[idStop+1]
				if err := createStop(stopNext); err != nil {
					return err
				}
				if err := createRelationBetweenTwoStops(stop, stopNext, busLine, dir); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func createStop(stop *stop) error {
	query := `MERGE (b:BusStop {name: {name}, idStop: {id}, lat: {lat}, lon: {lon}})
	ON CREATE
	SET b:BusStop`
	params := neoism.Props{
		"name": stop.Name,
		"id":   stop.ID,
		"lat":  stop.Lat,
		"lon":  stop.Lon,
	}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: params,
	}
	if err := DB.Cypher(&cq); err != nil {
		return err
	}
	log.Infof("Stop %s created\n", stop.Name)
	return nil
}

func createRelationBetweenTwoStops(startStop *stop, endStop *stop, busLine *busLine, dir int) error {
	query := "MATCH (a:BusStop), (b:BusStop) WHERE a.idStop = {idStopA} AND b.idStop = {idStopB} CREATE (a)-[p:PATH {busLineID: {busLineID}, direction: {dir}}]->(b)"
	params := neoism.Props{
		"idStopA":   startStop.ID,
		"idStopB":   endStop.ID,
		"busLineID": busLine.ID,
		"dir":       dir,
	}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: params,
	}
	if err := DB.Cypher(&cq); err != nil {
		return err
	}
	log.Infof("Relationship between %s and %s of for line %s\n", startStop.Name, endStop.Name, busLine.Name)
	return nil
}
