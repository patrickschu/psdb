package main

import (
	"fmt"
	"os"
	"strings"
)

// user needs to put input like FROM $NAME SELECT * WHERE $KEY=$VALUE

func main() {
	// parses user input
	userinput := os.Args[1:]
	if len(userinput) < 1 {
		panic("Not enough arguments to psdb")
	}
	if userinput[0] == "CREATE" {
		db := DataBase{logfile: userinput[1]}
		fmt.Printf("Created DB '%s'\n", db.logfile)
	}
	if userinput[0] == "ADD" {
		//ADD testdb name=patrick
		db := DataBase{logfile: userinput[1]}
		// oops default index and reconstructing string
		querystring := strings.Join(userinput[2:], " ")
		db.ADD(querystring, "1")
	}
	if userinput[0] == "SELECT" {
		//SELECT testdb WHERE x=y
		db := DataBase{logfile: userinput[1]}
		querystring := strings.Join(userinput[3:], " ")
		//fmt.Printf("Received query string %s\n", querystring)
		res, _ := db.SELECT(querystring)
		fmt.Printf("Found %d record\n%s\n", len(res), res)
	}

}
