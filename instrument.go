package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

func processInstruments(filename string, db *sql.DB) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s", filename)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var instExt INSTEXT
	xml.Unmarshal(byteValue, &instExt)
	insertInstruments(&instExt, db)
}

func insertInstruments(instr *INSTEXT, db *sql.DB) {
	for i := 0; i < len(instr.INSTRUMENT); i++ {
		execInstruments(instr, db, i)
	}
}
