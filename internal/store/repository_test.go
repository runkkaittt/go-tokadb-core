package store

import (
	"testing"
)

func TestSetGetDelete(t *testing.T) {
	db := NewBucket("bucket_test")

	db.Set("testKey1", "testVal1")
	val1, ok := db.Get("testKey1")
	if val1 != "testVal1" && ok != true {
		t.Error("Respone from db.Get is unexpected value 1")
	}

	db.Set("testKey1", "testVal2")
	val2, ok := db.Get("testKey1")
	if val2 != "testVal2" && ok != true {
		t.Error("Respone from db.Get is unexpected value 2")
	}

	db.Delete("testKey1")
	val3, ok := db.Get("testKey1")
	if val3 != nil && ok != false {
		t.Error("Respone from db.Get is unexpected value 3")
	}
}
