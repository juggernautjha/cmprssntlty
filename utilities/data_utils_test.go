// BEGIN: 3f8a9d5c4b7c
package utilities

import (
	"reflect"
	"testing"
)

func TestGet_and_count(t *testing.T) {
	db := Get_datablock([]string{"a", "b", "c", "a", "b", "d"})
	expectedList := []string{"a", "b", "c", "d"}
	expectedMap := map[string]int{"a": 2, "b": 2, "c": 1, "d": 1}
	list, m := db.Get_and_count()
	if !reflect.DeepEqual(list, expectedList) {
		t.Errorf("Get_and_count() returned list %v, expected %v", list, expectedList)
	}
	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("Get_and_count() returned map %v, expected %v", m, expectedMap)
	}
}

// END: 3f8a9d5c4b7c
