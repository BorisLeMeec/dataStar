package main

import (
	"fmt"
	"testing"
)

func TestPopulateDB(t *testing.T) {
	err := createTownFromJSON("jsonData.json")

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
	}
}
