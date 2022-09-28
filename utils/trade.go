package utils

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
) 



type TRADEEXT struct {
	XMLName xml.Name `xml:"TRADEEXT"`
	Text    string   `xml:",chardata"`
	TRADE   []struct {
		Text             string `xml:",chardata"`
		Recordno         string `xml:"recordno,attr"`
		Glosstraderef    string `xml:"glosstraderef,attr"`
		Versiono         string `xml:"versiono,attr"`
		Origin           string `xml:"origin,attr"`
		Tradetype        string `xml:"tradetype,attr"`
		Settlementstatus string `xml:"settlementstatus,attr"`
		Tradestatus      string `xml:"tradestatus,attr"`
		Originversion    string `xml:"originversion,attr"`
		EXTERNALREF      []struct {
			Text       string `xml:",chardata"`
			Extreftype string `xml:"extreftype,attr"`
			Extref     string `xml:"extref,attr"`
		} `xml:"EXTERNALREF"`
		DATE []struct {
			Text        string `xml:",chardata"`
			Datetype    string `xml:"datetype,attr"`
			Datewil     string `xml:"datewil,attr"`
			Time        string `xml:"time,attr"`
			Versionuser string `xml:"versionuser,attr"`
		} `xml:"DATE"`
		DRIVER []struct {
			Text       string `xml:",chardata"`
			Drivertype string `xml:"drivertype,attr"`
			Drivercode string `xml:"drivercode,attr"`
		} `xml:"DRIVER"`
		DURATION struct {
			Text         string `xml:",chardata"`
			Durationtype string `xml:"durationtype,attr"`
			Durationunit string `xml:"durationunit,attr"`
			Duration     string `xml:"duration,attr"`
		} `xml:"DURATION"`
		INSTRUMENT []struct {
			Text          string `xml:",chardata"`
			Instrtype     string `xml:"instrtype,attr"`
			P2000instrref string `xml:"p2000instrref,attr"`
			Instrreftype  string `xml:"instrreftype,attr"`
			Instrref      string `xml:"instrref,attr"`
			Longname      string `xml:"longname,attr"`
			Quantity      string `xml:"quantity,attr"`
			INSTEXT       []struct {
				Text    string `xml:",chardata"`
				Service string `xml:"service,attr"`
				Extref  string `xml:"extref,attr"`
			} `xml:"INSTEXT"`
		} `xml:"INSTRUMENT"`
		JOURNAL []struct {
			Text            string `xml:",chardata"`
			Accountscompany string `xml:"accountscompany,attr"`
			Journaltype     string `xml:"journaltype,attr"`
			Postingtype     string `xml:"postingtype,attr"`
			Journalno       string `xml:"journalno,attr"`
			Procaction      string `xml:"procaction,attr"`
		} `xml:"JOURNAL"`
		LINK []struct {
			Text            string `xml:",chardata"`
			LinkTypeWil     string `xml:"link_type_wil,attr"`
			MainRecordNoWil string `xml:"main_record_no_wil,attr"`
			SubRecordNoWil  string `xml:"sub_record_no_wil,attr"`
			LinkQtyWil      string `xml:"link_qty_wil,attr"`
			LinkStatusWil   string `xml:"link_status_wil,attr"`
		} `xml:"LINK"`
		TRADENARRATIVE []struct {
			Text          string `xml:",chardata"`
			Narrativecode string `xml:"narrativecode,attr"`
			Seqno         string `xml:"seqno,attr"`
			Narrative     string `xml:"narrative,attr"`
		} `xml:"TRADENARRATIVE"`
		PARTY []struct {
			Text         string `xml:",chardata"`
			Tradeparty   string `xml:"tradeparty,attr"`
			Partyref     string `xml:"partyref,attr"`
			Partyreftype string `xml:"partyreftype,attr"`
			Extpartyref  string `xml:"extpartyref,attr"`
			Longname     string `xml:"longname,attr"`
			PARTYDRIVER  []struct {
				Text       string `xml:",chardata"`
				Tradeparty string `xml:"tradeparty,attr"`
				Drivertype string `xml:"drivertype,attr"`
				Drivercode string `xml:"drivercode,attr"`
			} `xml:"PARTYDRIVER"`
		} `xml:"PARTY"`
		RATE []struct {
			Text             string `xml:",chardata"`
			Chargelevytype   string `xml:"chargelevytype,attr"`
			Actualcharge     string `xml:"actualcharge,attr"`
			Amounttype       string `xml:"amounttype,attr"`
			Rateinstrreftype string `xml:"rateinstrreftype,attr"`
			Rateinstrref     string `xml:"rateinstrref,attr"`
			Rateinstrid      string `xml:"rateinstrid,attr"`
			Rateentered      string `xml:"rateentered,attr"`
			Chargerate       string `xml:"chargerate,attr"`
			Multdivind       string `xml:"multdivind,attr"`
		} `xml:"RATE"`
		AMOUNT []struct {
			Text               string `xml:",chardata"`
			ChargeLevyTypeP2k  string `xml:"charge_levy_type_p2k,attr"`
			ChargeLevyInstrP2k string `xml:"charge_levy_instr_p2k,attr"`
			ChargeDiscountWil  string `xml:"charge_discount_wil,attr"`
			ChargeLevyQtyP2k   string `xml:"charge_levy_qty_p2k,attr"`
			ChargeLevyRateP2k  string `xml:"charge_levy_rate_p2k,attr"`
		} `xml:"AMOUNT"`
		SETTLEMENT []struct {
			Text             string `xml:",chardata"`
			Deliverytype     string `xml:"deliverytype,attr"`
			Settleeventinstr string `xml:"settleeventinstr,attr"`
			Settleterms      string `xml:"settleterms,attr"`
			Originalqty      string `xml:"originalqty,attr"`
			Openqty          string `xml:"openqty,attr"`
			Settledate       string `xml:"settledate,attr"`
			Delrecind        string `xml:"delrecind,attr"`
			Settlestatus     string `xml:"settlestatus,attr"`
			Tradestatus      string `xml:"tradestatus,attr"`
			Settlenarrative1 string `xml:"settlenarrative1,attr"`
			Settlenarrative2 string `xml:"settlenarrative2,attr"`
			Settlenarrative3 string `xml:"settlenarrative3,attr"`
			Settlenarrative4 string `xml:"settlenarrative4,attr"`
			Settlenarrative5 string `xml:"settlenarrative5,attr"`
			Settlenarrative6 string `xml:"settlenarrative6,attr"`
			Settlenarrative7 string `xml:"settlenarrative7,attr"`
			Settlenarrative8 string `xml:"settlenarrative8,attr"`
			CompAliasWil     string `xml:"comp_alias_wil,attr"`
			CompAliasDesc    string `xml:"comp_alias_desc,attr"`
			CompDepotTypeWil string `xml:"comp_depot_type_wil,attr"`
			CompDaccWil      string `xml:"comp_dacc_wil,attr"`
			CompServiceWil   string `xml:"comp_service_wil,attr"`
			SecpAliasWil     string `xml:"secp_alias_wil,attr"`
			SecpServiceWil   string `xml:"secp_service_wil,attr"`
		} `xml:"SETTLEMENT"`
		PROCESSING []struct {
			Text            string `xml:",chardata"`
			Procalias       string `xml:"procalias,attr"`
			Procaction      string `xml:"procaction,attr"`
			Duedatetime     string `xml:"duedatetime,attr"`
			Donedatetime    string `xml:"donedatetime,attr"`
			PROCESSINGEVENT []struct {
				Text           string `xml:",chardata"`
				Eventtype      string `xml:"eventtype,attr"`
				Eventdate      string `xml:"eventdate,attr"`
				Eventdateto    string `xml:"eventdateto,attr"`
				Entrydatetime  string `xml:"entrydatetime,attr"`
				Eventcode      string `xml:"eventcode,attr"`
				Exceptiontype  string `xml:"exceptiontype,attr"`
				Expirydate     string `xml:"expirydate,attr"`
				EVENTNARRATIVE struct {
					Text          string `xml:",chardata"`
					Narrativecode string `xml:"narrativecode,attr"`
					Seqno         string `xml:"seqno,attr"`
					Narrative     string `xml:"narrative,attr"`
				} `xml:"EVENTNARRATIVE"`
			} `xml:"PROCESSINGEVENT"`
		} `xml:"PROCESSING"`
		INSTRUCTION struct {
			Text             string `xml:",chardata"`
			Procaction       string `xml:"procaction,attr"`
			Destination      string `xml:"destination,attr"`
			Procstatus       string `xml:"procstatus,attr"`
			Recordidentifier string `xml:"recordidentifier,attr"`
			Recordversion    string `xml:"recordversion,attr"`
			Instformat       string `xml:"instformat,attr"`
			Tradeparty       string `xml:"tradeparty,attr"`
			Partyref         string `xml:"partyref,attr"`
			Deliverytype     string `xml:"deliverytype,attr"`
			Addresscode      string `xml:"addresscode,attr"`
			Servicestatus    string `xml:"servicestatus,attr"`
			Noofcopies       string `xml:"noofcopies,attr"`
			Duedatetime      string `xml:"duedatetime,attr"`
			INSTRUCTIONEVENT []struct {
				Text           string `xml:",chardata"`
				Eventtype      string `xml:"eventtype,attr"`
				Eventdate      string `xml:"eventdate,attr"`
				Eventdateto    string `xml:"eventdateto,attr"`
				Entrydatetime  string `xml:"entrydatetime,attr"`
				Eventcode      string `xml:"eventcode,attr"`
				Exceptiontype  string `xml:"exceptiontype,attr"`
				Expirydate     string `xml:"expirydate,attr"`
				EVENTNARRATIVE struct {
					Text          string `xml:",chardata"`
					Narrativecode string `xml:"narrativecode,attr"`
					Seqno         string `xml:"seqno,attr"`
					Narrative     string `xml:"narrative,attr"`
				} `xml:"EVENTNARRATIVE"`
			} `xml:"INSTRUCTIONEVENT"`
		} `xml:"INSTRUCTION"`
		EVENT []struct {
			Text           string `xml:",chardata"`
			Eventtype      string `xml:"eventtype,attr"`
			Eventdate      string `xml:"eventdate,attr"`
			Eventdateto    string `xml:"eventdateto,attr"`
			Entrydatetime  string `xml:"entrydatetime,attr"`
			Eventcode      string `xml:"eventcode,attr"`
			Exceptiontype  string `xml:"exceptiontype,attr"`
			Expirydate     string `xml:"expirydate,attr"`
			EVENTNARRATIVE struct {
				Text          string `xml:",chardata"`
				Narrativecode string `xml:"narrativecode,attr"`
				Seqno         string `xml:"seqno,attr"`
				Narrative     string `xml:"narrative,attr"`
			} `xml:"EVENTNARRATIVE"`
		} `xml:"EVENT"`
		TRADECODE []struct {
			Text       string `xml:",chardata"`
			Tradeclass string `xml:"tradeclass,attr"`
			Tradecode  string `xml:"tradecode,attr"`
		} `xml:"TRADECODE"`
	} `xml:"TRADE"`
}

// insert into trd_amount (trd_recordno, trd_charge_levy_type_p2k, trd_charge_levy_instr_p2k, trd_charge_discount_wil,
// trd_charge_levy_qty_p2k, trd_charge_levyrate_p2k)
// values ();

func ProcessTrade(filename string, db *sql.DB) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", filename)
	
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var trdExt TRADEEXT

	xml.Unmarshal(byteValue, &trdExt)
	insertTrades(&trdExt, db)
	
}


func insertTrades(trd *TRADEEXT, db *sql.DB) {
	for index := 0; index < len(trd.TRADE); index++ {
		trdTrade(trd, db, index)
		trdExternalRef(trd, db, index)
		trdSettlement(trd, db, index)
		trdEvent(trd,  db, index)
		trdDriver(trd,  db, index)
		trdEventNarrative(trd, db, index)
	}
}

func trdTrade(trd *TRADEEXT, db *sql.DB, i int) {
	insertStatement := fmt.Sprintf("insert into trd_trade ( trd_recordno, trd_glosstraderef, trd_versiono, trd_origin, trd_tradetype, trd_settlementstatus, trd_tradestatus, trd_originversion) values ( %s,%s,%s,'%s','%s','%s','%s',%s);",
		trd.TRADE[i].Recordno,
		trd.TRADE[i].Glosstraderef,
		trd.TRADE[i].Versiono,
		trd.TRADE[i].Origin,
		trd.TRADE[i].Tradetype,
		trd.TRADE[i].Settlementstatus,
		trd.TRADE[i].Tradestatus,
		trd.TRADE[i].Originversion)
	
		// fmt.Println(insertStatement)
	
		Exec(insertStatement, db)
}

func trdExternalRef(trd *TRADEEXT, db *sql.DB, index int) {
	var refType = "{"
	var ref = "{"
	for j := 0; j < len(trd.TRADE[index].EXTERNALREF); j++ {	
		    if (j>0) {
				refType += ","
			}
			refType += trd.TRADE[index].EXTERNALREF[j].Extreftype 
			if( j>0) {
				ref += ","
			}
			ref += trd.TRADE[index].EXTERNALREF[j].Extref 
	}

	refType += "}"
	ref += "}"

	insertStatement := fmt.Sprintf("insert into trd_external_reference (trd_recordno, ext_ref_extreftype, ext_ref_extref) values (%s,'%s','%s');",
	trd.TRADE[index].Recordno, refType, ref)
	
    // fmt.Println(insertStatement)

	Exec(insertStatement, db)
	
}

func trdSettlement(trd *TRADEEXT, db *sql.DB, index int) {
	var i = index
	for j := 0; j < len(trd.TRADE[i].SETTLEMENT); j++ {
		insertStatement := fmt.Sprintf("insert into trd_settlement (trd_recordno, trd_deliverytype, trd_settleeventinstr, trd_settleterms, trd_originalqty,"+
			"trd_openqty, trd_settledate, trd_delrecind, trd_settlestatus, trd_tradestatus, "+
			"trd_settlenarrative1, trd_settlenarrative2, trd_settlenarrative3, trd_settlenarrative4,"+
			"trd_settlenarrative5, trd_settlenarrative6, trd_settlenarrative7, trd_settlenarrative8, "+
			"trd_dompaliaswil, trd_dompaliasdesc, trd_dompdepottypewil, trd_dompdaccwil,"+
			"trd_dompservicewil, trd_secpaliaswil, trd_secpservicewil)"+
			"values (%s,'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');\n",
			trd.TRADE[i].Recordno,
			trd.TRADE[i].SETTLEMENT[j].Deliverytype,
			trd.TRADE[i].SETTLEMENT[j].Settleeventinstr,
			trd.TRADE[i].SETTLEMENT[j].Settleterms,
			trd.TRADE[i].SETTLEMENT[j].Originalqty,
			trd.TRADE[i].SETTLEMENT[j].Openqty,
			trd.TRADE[i].SETTLEMENT[j].Settledate,
			trd.TRADE[i].SETTLEMENT[j].Delrecind,
			trd.TRADE[i].SETTLEMENT[j].Settlestatus,
			trd.TRADE[i].SETTLEMENT[j].Tradestatus,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative1,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative2,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative3,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative4,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative5,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative6,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative7,
			trd.TRADE[i].SETTLEMENT[j].Settlenarrative8,
			trd.TRADE[i].SETTLEMENT[j].CompAliasWil,
			trd.TRADE[i].SETTLEMENT[j].CompAliasDesc,
			trd.TRADE[i].SETTLEMENT[j].CompDepotTypeWil,
			trd.TRADE[i].SETTLEMENT[j].CompDaccWil,
			trd.TRADE[i].SETTLEMENT[j].CompServiceWil,
			trd.TRADE[i].SETTLEMENT[j].SecpAliasWil,
			trd.TRADE[i].SETTLEMENT[j].SecpServiceWil)
			Exec(insertStatement, db)
	}
}

func Exec(query string, db *sql.DB) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func trdEvent(trd *TRADEEXT,  db *sql.DB, index int) string {
	var insertStatement string
	for j := 0; j < len(trd.TRADE[index].EXTERNALREF); j++ {
		insertStatement = fmt.Sprintf(`insert into trd_event (
                       trd_recordno, 
                       trd_eventtype, 
                       trd_eventdate, 
                       trd_eventdateto, 
                       trd_entrydatetime, 
                       trd_eventcode, 
                       trd_exceptiontype, 
                       trd_expirydate) values (%s,'%s','%s','%s','%s','%s','%s','%s');`,
			trd.TRADE[index].Recordno,
			trd.TRADE[index].EVENT[j].Eventtype,
			trd.TRADE[index].EVENT[j].Eventdate,
			trd.TRADE[index].EVENT[j].Eventdateto,
			trd.TRADE[index].EVENT[j].Entrydatetime,
			trd.TRADE[index].EVENT[j].Eventcode,
			trd.TRADE[index].EVENT[j].Exceptiontype,
			trd.TRADE[index].EVENT[j].Expirydate)
	}
	return insertStatement
}

func trdDriver(trd *TRADEEXT,  db *sql.DB, index int) {
	var insertStatement string
	for j := 0; j < len(trd.TRADE[index].DRIVER); j++ {
		fmt.Sprintf(`insert into trd_driver (trd_recordno, trd_drivertype, trd_drivercode)  values (%s,'%s','%s');`,
			trd.TRADE[index].Recordno,
			trd.TRADE[index].DRIVER[j].Drivertype,
			trd.TRADE[index].DRIVER[j].Drivercode)
	}
	Exec(insertStatement, db)
}

func trdEventNarrative(trd *TRADEEXT, db *sql.DB, index int ) {
	var insertStatement string
	for j := 0; j < len(trd.TRADE[index].EVENT); j++ {
		insertStatement = fmt.Sprintf(`insert into trd_event_narrative (trd_recordno, trd_narrative_code, trd_seqno, trd_narrative) values (%s,'%s','%s','%s');`,
			trd.TRADE[index].Recordno,
			trd.TRADE[index].EVENT[j].EVENTNARRATIVE.Narrativecode,
			trd.TRADE[index].EVENT[j].EVENTNARRATIVE.Narrative,
			trd.TRADE[index].EVENT[j].EVENTNARRATIVE.Seqno)
	}
	 Exec(insertStatement, db)
}

func trdInstExt(trd *TRADEEXT,  db *sql.DB, index int) {
	var insertStatement string
	for j := 0; j < len(trd.TRADE[index].INSTRUCTION.INSTRUCTIONEVENT); j++ {
		insertStatement = fmt.Sprintf(`insert into trd_inst_ext (trd_recordno, trd_service, trd_extref) values (%s,'%s','%s');`,
			trd.TRADE[index].INSTRUCTION.INSTRUCTIONEVENT)
	}
	Exec(insertStatement, db)
}

func trdInstruction(trd *TRADEEXT, db *sql.DB, index int) {
	var insertStatement string
	for j := 0; j < len(trd.TRADE[index].INSTRUCTION.INSTRUCTIONEVENT); j++ {
		insertStatement =
	insert into trd_instruction_effect (trd_recordno, trd_eventtype, trd_eventdate, trd_eventdateto, trd_entrydatetime,
										trd_eventcode, trd_exceptiontype, trd_expirydate)
	values ();
}

func instEvent(inst *INSTEXT, db *sql.DB, i int) {
	insert into trd_event (trd_recordno, trd_eventtype, trd_eventdate, trd_eventdateto, trd_entrydatetime, trd_eventcode,
                       trd_exceptiontype, trd_expirydate)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}


func trdInstrument(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_journal (trd_recordno, trd_accounts_company, trd_journal_type, trd_posting_type, trd_journal_no,
                         trd_procaction)
	values ();	
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func trdJournal(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_journal (trd_recordno, trd_accounts_company, trd_journal_type, trd_posting_type, trd_journal_no,
						 trd_procaction)
	values ();	
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func trdInstruction(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_instruction_effect (trd_recordno, trd_eventtype, trd_eventdate, trd_eventdateto, trd_entrydatetime,
										trd_eventcode, trd_exceptiontype, trd_expirydate)
	values ();
}

func trdParty(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_party (trd_recordno, trd_trade_party, trd_partyref, trd_partyref_type_text, trd_ext_partyref,
						   trd_longname)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func func trdParty(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_party (trd_recordno, trd_trade_party, trd_partyref, trd_partyref_type_text, trd_ext_partyref,
						   trd_longname)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}


func trdDriver(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_driver (trd_recordno, trd_drivertype, trd_drivercode)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}


func  trdProcessing(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_processing (trd_recordno, trd_procaction, trd_destination, trd_procstatus, trd_recordidentifier,
								trd_recordversion, trd_instformat, trd_tradeparty, trd_partyref, trd_deliverytype,
								trd_addresscode, trd_servicestatus, trd_noofcopies, trd_duedatetime)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func  trdProcessingEvent(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_processing_event (trd_recordno, trd_eventtype, trd_eventdate, trd_eventdateto, trd_entrydatetime,
		trd_eventcode, trd_exceptiontype, trd_expirydate)
		values ();
	
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}

func trdPartyDriver(trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_party_driver (trd_recordno, trd_trade_party, trd_driver_type, trd_driver_code)
	values ();

	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}	 
}


func trdRate (trd *TRADEEXT, db *sql.DB, index int) {
	insert into trd_rate (trd_recordno, trd_charge_levy_type, trd_actual_charge, trd_amount_type, trd_rate_instrref_type,
		trd_rate_instrref, trd_rate_instrid, trd_rate_entered, trd_charge_rate, trd_mult_divind)
	values ();
	insertStatement := fmt.Sprintf(``,
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s", err)
	}
}



