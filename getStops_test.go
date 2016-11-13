package main

import (
	"fmt"
	"testing"
)

func TestGetStopsFromLineID(t *testing.T) {
	res, err := GetStopsFromLineID("0004")

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
		fmt.Print(res)
	}
}

func TestGetLinesByStopID(t *testing.T) {
	res, err := GetLinesByStopID("1214")

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
		fmt.Printf("%s\n", res)
	}
}

func TestGetStopAutocomplete(t *testing.T) {
	res, err := getStopAutocomplete("Olympe")

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
		fmt.Print(res)
	}
}
