package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	db := Connect()
	db.Create("some", "thing")
}
