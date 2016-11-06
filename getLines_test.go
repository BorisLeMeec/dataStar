package main

import (
	"fmt"
	"testing"
)

func TestGetLines(t *testing.T) {
	res, err := GetAllLines()

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
		fmt.Printf("%s\n", res)
	}
}
