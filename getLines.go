package main

import "github.com/jmcvetta/neoism"

//Line is used to represent a line
type Line struct {
	Name      string `json:"Name"`
	ID        string `json:"ID"`
	TextColor string `json:"TextColor"`
	Color     string `json:"Color"`
}

//Lines is used to represent multiple Line
type Lines []Line

//GetAllLines return all lines
func GetAllLines() (res Lines, err error) {
	query := "MATCH (n:BusLine) RETURN n.name AS Name, n.id AS ID, n.textColor AS TextColor, n.color AS Color"
	cq := neoism.CypherQuery{
		Statement: query,
		Result:    &res,
	}
	if err = DB.Cypher(&cq); err != nil {
		return nil, err
	}
	return res, nil
}

//GetLinesByStopID return all the lines passing at stopID stopBus
func GetLinesByStopID(stopID string) (res Lines, err error) {
	query := "MATCH (b:BusStop)-[p:PATH]->(a:BusStop), (r:BusLine) WHERE b.idStop = {stopID} AND r.id = p.busLineID RETURN r.id AS ID, r.textColor AS TextColor, r.color AS Color"
	params := neoism.Props{
		"stopID": stopID,
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
