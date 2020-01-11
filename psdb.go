package psdb

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func append(filename string, record string) error {
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

func search(filename string, index string) (result string, err error) {
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

//DB needs to 1) add 2) delete 3) look up
//DB should work via CL psdb FROM $NAME GET $INDEX
