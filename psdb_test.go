package psdb

import (
	"encoding/json"
	"fmt"
	"testing"
)

// func Test(t *testing.T) {
// 	got := Abs(-1)
// 	if got != 1 {
// 		t.Errorf("Abs(-1) = %d; want 1", got)
// 	}
// }

func Test_checkvalue(t *testing.T) {

}

func Test_searchvalue(t *testing.T) {

}

func Test_inputtomap(t *testing.T) {
	instring := "name=patrick age=122"
	recmap, _ := inputtomap(instring)
	if recmap["name"] != "patrick" {
		t.Fatalf("got value %s for key %s", recmap["name"], "name")
	}
	if recmap["age"] != "122" {
		t.Fatalf("got value %s for key %s", recmap["age"], "age")
	}
	fmt.Println(recmap)
}

func Test_addrecord(t *testing.T) {
	fname := "testdb.txt"
	testdb := new(DataBase)
	testdb.logfile = fname
	testdb.ADD("method=test value=record", "1")
}

func Test_append(t *testing.T) {
	dbappend("dbfile.txt", "this be a test record")
	dbappend("dbfile.txt", "12, test record with index 12")
}

func Test_tojson(t *testing.T) {
	testmap := map[string]string{"key1": "value1", "key2": "value100"}
	res, _ := tojson(testmap)
	fmt.Printf("res to json %v", res)
	var unmarshaled map[string]string
	json.Unmarshal([]byte(res), &unmarshaled)
}

func Test_fromjson(t *testing.T) {
	testmap := map[string]string{"key1": "value1", "key2": "value100"}
	res, _ := tojson(testmap)
	unmarshaled, _ := fromjson(res)
	fmt.Println("\nfrom json", unmarshaled)
}

func Test_searchindex(t *testing.T) {
	want := "good rec"
	ind := "100"
	dbappend("searchfile.txt", ind+","+want)
	res, err := searchindex("searchfile.txt", ind)
	if err != nil {
		t.Fatalf("got error %s for index %s", err, ind)
	}
	if res != want {
		t.Fatalf("want %s, got %s", want, res)
	}
	ind = "10"
	res, err = searchindex("searchfile.txt", ind)
	if err == nil {
		t.Fatalf("raise fail: did not return error for index %s", ind)
	}
}

func Test_large(t *testing.T) {
	entries := []string{"first entry", "second entry", "some text here"}
	indices := []string{"1", "1000", "23"}
	for num, record := range entries {
		dbappend("large_db.txt", indices[num]+","+record)
	}
}
