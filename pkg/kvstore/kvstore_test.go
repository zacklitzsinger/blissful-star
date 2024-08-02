package kvstore

import "testing"

func TestSet(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")

	if kv.Get("some") != "thing" {
		t.Error("Set failed")
	}
}

func TestDelete(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")
	kv.Delete("some")

	if kv.Get("some") != "" {
		t.Error("Delete failed")
	}
}

func TestSetExisting(t *testing.T) {
	kv := CreateKV()
	kv.Set("some", "thing")

	if kv.Get("some") != "thing" {
		t.Error("Set failed")
	}
}

func TestDeleteNonExisting(t *testing.T) {
	kv := CreateKV()

	if err := kv.Delete("some"); err == nil {
		t.Error("DeleteNonExisting failed")
	}
}
