package main

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create()

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
		t.Error("Test failed")
	} else {
	}
}
