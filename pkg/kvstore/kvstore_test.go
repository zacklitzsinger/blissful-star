package kvstore

import "testing"

func TestSet(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")

	if kv.Get("some") != "thing" {
		t.Error("Set failed")
	}
}

func TestUnset(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")
	kv.Unset("some")

	if kv.Get("some") != "" {
		t.Error("Delete failed")
	}
}

func TestSetExisting(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")
	kv.Set("some", "other")

	if kv.Get("some") != "other" {
		t.Error("Set failed")
	}
}

func TestDeleteNonExisting(t *testing.T) {
	kv := CreateKV()
	kv.Unset("some")
}
