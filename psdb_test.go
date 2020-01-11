package psdb

import "testing"

// func Test(t *testing.T) {
// 	got := Abs(-1)
// 	if got != 1 {
// 		t.Errorf("Abs(-1) = %d; want 1", got)
// 	}
// }

func Test_append(t *testing.T) {
	append("dbfile.txt", "this be a test record")
	append("dbfile.txt", "12, test record with index 12")
}

func Test_search(t *testing.T) {
	want := "good rec"
	ind := "100"
	append("searchfile.txt", ind+","+want)
	res, err := search("searchfile.txt", ind)
	if err != nil {
		t.Fatalf("got error %s for index %s", err, ind)
	}
	if res != want {
		t.Fatalf("want %s, got %s", want, res)
	}
	ind = "10"
	res, err = search("searchfile.txt", ind)
	if err == nil {
		t.Fatalf("raise fail: did not return error for index %s", ind)
	}
}

func Test_large(t *testing.T) {
	entries := []string{"first entry", "second entry", "some text here"}
	indices := []string{"1", "1000", "23"}
	for num, record := range entries {
		append("large_db.txt", indices[num]+","+record)
	}
}
