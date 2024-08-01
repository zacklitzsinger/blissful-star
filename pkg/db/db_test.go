// test file for db.go

package db

import "testing"

func TestCreate(t *testing.T) {
	db := Connect()
	db.Create("some", "thing")

	if db.Get("some") != "thing" {
		t.Error("Create failed")
	}
}

func TestUpdate(t *testing.T) {
	db := Connect()
	db.Create("some", "thing")
	db.Update("some", "nothing")

	if db.Get("some") != "nothing" {
		t.Error("Update failed")
	}
}

func TestDelete(t *testing.T) {
	db := Connect()
	db.Create("some", "thing")
	db.Delete("some")

	if db.Get("some") != "" {
		t.Error("Delete failed")
	}
}

func TestCreateExisting(t *testing.T) {
	db := Connect()
	db.Create("some", "thing")

	if err := db.Create("some", "thing"); err == nil {
		t.Error("CreateExisting failed")
	}
}

func TestUpdateNonExisting(t *testing.T) {
	db := Connect()

	if err := db.Update("some", "thing"); err == nil {
		t.Error("UpdateNonExisting failed")
	}
}

func TestDeleteNonExisting(t *testing.T) {
	db := Connect()

	if err := db.Delete("some"); err == nil {
		t.Error("DeleteNonExisting failed")
	}
}
