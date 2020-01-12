package psdb

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type DataBase struct {
	//DataBase is the basic struct for DataBase
	logfile string
}

func inputtomap(userinput string) (map[string]string, error) {
	//convert the string received from CL to map for DB
	// this should really have some error checks
	recordmap := map[string]string{}
	for _, kvpair := range strings.Split(userinput, " ") {
		kvsplit := strings.SplitN(kvpair, "=", 2)
		key := kvsplit[0]
		value := kvsplit[1]
		recordmap[key] = value
	}
	return recordmap, nil
}

func (db *DataBase) ADD(record string, ind string) error {
	//add a record to the DB
	//user input name=patrick home=away
	//convert inputtstring to map
	recordmap, _ := inputtomap(record)
	recordjson, _ := tojson(recordmap)
	//convert map to to JSON for write, might be cumbersome
	dbrecord := ind + "," + recordjson
	err := dbappend(db.logfile, dbrecord)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) SELECT(query string) (matches []string, err error) {
	//get a record from db
	//user input SELECT FROM testdb * WHERE x=y
	//convert inputstring to map
	results := []string{}
	querymap, _ := inputtomap(query)
	// for each key in querymap check matching records
	for key, val := range querymap {
		matches, err := searchvalue(db.logfile, []string{key, val})
		if err == nil {
			results = append(results, matches...)
		}
	}
	return matches, nil
}

func checkvalue(kvpair []string, line string) (match string, err error) {
	//line will look like `index, {}`
	//convert from JSON string to map
	if len(kvpair) != 2 {
		return "", fmt.Errorf("kvpair needs to be len 2, is %v", len(kvpair))
	}
	recordmap := strings.SplitN(line, ",", 2)[1]
	linemap, _ := fromjson(recordmap)
	if linemap[kvpair[0]] == kvpair[1] {
		return line, nil
	}
	return "", fmt.Errorf("No match %v", kvpair)
}

func searchvalue(filename string, kvpair []string) (results []string, err error) {
	// searchvalue returns for name=patrick return all records where this true
	// checkvalue does the line by line checking
	matches := []string{}
	fil, err := os.Open(filename)
	if err != nil {
		return matches, fmt.Errorf("Cannot open file %s", filename)
	}
	defer fil.Close()

	scan := bufio.NewScanner(fil)
	for scan.Scan() {
		res, err := checkvalue(kvpair, scan.Text())
		if err == nil {
			matches = append(matches, res)
		}
	}
	return matches, fmt.Errorf("Condition %s not matched", kvpair)
}

// CREATE testdb
// testdb ADD name=patrick age=122
// FROM testdb SELECT name WHERE age=122

func dbappend(filename string, record string) error {
	// append appends a `record` to the file at `filename`
	fil, openerr := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if openerr != nil {
		return fmt.Errorf("Cannot open file %s", filename)
	}
	_, writeerr := fil.Write([]byte(record + "\n"))
	if writeerr != nil {
		return fmt.Errorf("Cannot write to file %s", filename)
	}
	closeerr := fil.Close()
	if closeerr != nil {
		return fmt.Errorf("Cannot close file %s", filename)
	}
	return nil
}

func checkline(index string, line string) (result string, err error) {
	//checkline returns `line` if it corresponds to `index`
	if strings.HasPrefix(line, index+",") {
		return line, nil
	}
	return "", errors.New("Index not matched")
}

func searchindex(filename string, index string) (result string, err error) {
	// search for `index` in `filename`
	fil, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Cannot open file %s", filename)
	}
	defer fil.Close()

	scan := bufio.NewScanner(fil)
	for scan.Scan() {
		res, err := checkline(index, scan.Text())
		if err == nil {
			return strings.SplitN(res, ",", 2)[1], err
		}
	}
	return "", fmt.Errorf("Index %s not matched", index)
}

func tojson(record map[string]string) (string, error) {
	// convert map to json str
	jsonrecord, err := json.Marshal(record)
	if err != nil {
		return "", err
	}
	return string(jsonrecord), nil
}

func fromjson(jsonobject string) (map[string]string, error) {
	//convert json string to map
	var unmarshaled map[string]string
	err := json.Unmarshal([]byte(jsonobject), &unmarshaled)
	if err != nil {
		return map[string]string{}, err
	}
	return unmarshaled, nil
}

//DB needs to 1) add 2) delete 3) look up
//DB should work via CL psdb FROM $NAME GET $INDEX
//DB needs conditionals psdb FROM $NAME GET * WHERE $CONDITION
