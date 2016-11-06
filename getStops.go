package main

import (
	"encoding/json"

	"github.com/jmcvetta/neoism"
)

//Stop represent a stop bus
type Stop struct {
	IDStop    string `json:"IDStop"`
	Name      string `json:"Name"`
	Direction int    `json:"Dir"`
}

//GetStopsFromLineID return all stops for a lineId
func GetStopsFromLineID(idLine string) (res []Stop, err error) {
	query := "MATCH (b:BusStop)-[p:PATH]->() WHERE p.busLineID = {busLineID} RETURN b.idStop AS IDStop, b.name AS Name, p.direction AS Dir"
	params := neoism.Props{
		"busLineID": idLine,
	}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: params,
		Result:     &res,
	}
	if err = DB.Cypher(&cq); err != nil {
		return nil, err
	}
	return res, nil
}

func getStopAutocomplete(name string) (res []byte, err error) {
	resNeo4j := []struct {
		IDStop    string `json:"IDStop"`
		Name      string `json:"Name"`
		Direction int    `json:"Dir"`
	}{}
	query := "MATCH (b:BusStop)-[p:PATH]->(n) WHERE b.name =~ '(?ui).*{name}.*' return DISTINCT b.idStop AS IDStop, b.name AS Name, p.direction AS Dir"
	params := neoism.Props{
		"name": name,
	}
	cq := neoism.CypherQuery{
		Statement:  query,
		Parameters: params,
		Result:     &resNeo4j,
	}
	if err = DB.Cypher(&cq); err != nil {
		return nil, err
	}
	res, err = json.Marshal(resNeo4j)
	if err != nil {
		return nil, err
	}
	return res, nil
}
