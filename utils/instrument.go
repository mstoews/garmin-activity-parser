package utils

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)



type INSTEXT struct {
	XMLName    xml.Name `xml:"INSTEXT"`
	Text       string   `xml:",chardata"`
	INSTRUMENT []struct {
		Text       string `xml:",chardata"`
		Instref    string `xml:"instref,attr"`
		Instgrp    string `xml:"instgrp,attr"`
		Longdesc   string `xml:"longdesc,attr"`
		Denminst   string `xml:"denminst,attr"`
		Pricedps   string `xml:"pricedps,attr"`
		Divisor    string `xml:"divisor,attr"`
		Multiplier string `xml:"multiplier,attr"`
		Pricetype  string `xml:"pricetype,attr"`
		Tick       string `xml:"tick,attr"`
		Accrued    string `xml:"accrued,attr"`
		Minqty     string `xml:"minqty,attr"`
		Qtydps     string `xml:"qtydps,attr"`
		Shrtdesc   string `xml:"shrtdesc,attr"`
		Market     string `xml:"market,attr"`
		Settinst   string `xml:"settinst,attr"`
		Book       string `xml:"book,attr"`
		Maxmovt    string `xml:"maxmovt,attr"`
		Minmovt    string `xml:"minmovt,attr"`
		Active     string `xml:"active,attr"`
		Verdat     string `xml:"verdat,attr"`
		INSTFLAG   []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			Code  string `xml:"code,attr"`
		} `xml:"INSTFLAG"`
		INSTCLASS []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			Code  string `xml:"code,attr"`
		} `xml:"INSTCLASS"`
		DENOM struct {
			Text    string `xml:",chardata"`
			Denmqty string `xml:"denmqty,attr"`
		} `xml:"DENOM"`
		INSTNARR []struct {
			Text      string `xml:",chardata"`
			Narrcode  string `xml:"narrcode,attr"`
			SeqNoWil  string `xml:"seq_no_wil,attr"`
			Narrative string `xml:"narrative,attr"`
		} `xml:"INSTNARR"`
		INSTEXT []struct {
			Text    string `xml:",chardata"`
			Service string `xml:"service,attr"`
			Extref  string `xml:"extref,attr"`
		} `xml:"INSTEXT"`
		INSTAMT []struct {
			Text     string `xml:",chardata"`
			Amnt     string `xml:"amnt,attr"`
			Amntqty  string `xml:"amntqty,attr"`
			Amntinst string `xml:"amntinst,attr"`
		} `xml:"INSTAMT"`
		INSTPTY []struct {
			Text  string `xml:",chardata"`
			Itype string `xml:"itype,attr"`
			Iref  string `xml:"iref,attr"`
			Iqty  string `xml:"iqty,attr"`
		} `xml:"INSTPTY"`
		INSTASSC []struct {
			Text     string `xml:",chardata"`
			Assctype string `xml:"assctype,attr"`
			Asscinst string `xml:"asscinst,attr"`
		} `xml:"INSTASSC"`
		ATTACHED struct {
			Text      string `xml:",chardata"`
			Assoc     string `xml:"assoc,attr"`
			Attach    string `xml:"attach,attr"`
			Attachamt string `xml:"attachamt,attr"`
			Hostamt   string `xml:"hostamt,attr"`
			Sepdate   string `xml:"sepdate,attr"`
		} `xml:"ATTACHED"`
		EQUITY struct {
			Text    string `xml:",chardata"`
			Registr string `xml:"registr,attr"`
		} `xml:"EQUITY"`
		ISSUE struct {
			Text           string `xml:",chardata"`
			Issuer         string `xml:"issuer,attr"`
			Issuetype      string `xml:"issuetype,attr"`
			Issueprice     string `xml:"issueprice,attr"`
			Issuepricetype string `xml:"issuepricetype,attr"`
			Issueqty       string `xml:"issueqty,attr"`
			Issuedate      string `xml:"issuedate,attr"`
			Leadmanager    string `xml:"leadmanager,attr"`
			Guarantor      string `xml:"guarantor,attr"`
			Guarantype     string `xml:"guarantype,attr"`
		} `xml:"ISSUE"`
	} `xml:"INSTRUMENT"`
}


func ProcessInstruments(filename string, db *sql.DB) {
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
		execTrdInstruments(instr, db, i)
	}
}

func execTrdInstruments(inst *INSTEXT, db *sql.DB, i int) {
	insertStatement := fmt.Sprintf(`insert into instr_instruments (instr_instref, instr_instgrp, instr_longdesc, instr_denminst, instr_pricedps,
                               instr_divisor, instr_multiplier, instr_pricetype, instr_tick, instr_accrued,
                               instr_minqty, instr_qtydps, instr_shrtdesc, instr_market, instr_settinst, instr_book,instr_maxmovt, instr_minmovt, instr_active, instr_verdat) 
								values ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');`,
		inst.INSTRUMENT[i].Instref,
		inst.INSTRUMENT[i].Instgrp,
		inst.INSTRUMENT[i].Longdesc,
		inst.INSTRUMENT[i].Denminst,
		inst.INSTRUMENT[i].Pricedps,
		inst.INSTRUMENT[i].Divisor,
		inst.INSTRUMENT[i].Multiplier,
		inst.INSTRUMENT[i].Pricetype,
		inst.INSTRUMENT[i].Tick,
		inst.INSTRUMENT[i].Accrued,
		inst.INSTRUMENT[i].Minqty,
		inst.INSTRUMENT[i].Qtydps,
		inst.INSTRUMENT[i].Shrtdesc,
		inst.INSTRUMENT[i].Market,
		inst.INSTRUMENT[i].Settinst,
		inst.INSTRUMENT[i].Book,
		inst.INSTRUMENT[i].Maxmovt,
		inst.INSTRUMENT[i].Minmovt,
		inst.INSTRUMENT[i].Active,
		inst.INSTRUMENT[i].Verdat)

	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}
