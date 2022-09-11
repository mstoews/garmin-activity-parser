package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

// insert into trd_amount (trd_recordno, trd_charge_levy_type_p2k, trd_charge_levy_instr_p2k, trd_charge_discount_wil,
// trd_charge_levy_qty_p2k, trd_charge_levyrate_p2k)
// values ();

func processTrade(filename string, db *sql.DB) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s", filename)
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var trdExt TRADEEXT

	xml.Unmarshal(byteValue, &trdExt)
	insertTrades(&trdExt, db)
}

func insertTrades(trd *TRADEEXT, db *sql.DB) {
	for index := 0; index < len(trd.TRADE); index++ {
		execTrdTrade(trd, db, index)
		execTrdExternalRef(trd, db, index)
		execTrdSettlement(trd, db, index)
	}
}

func execTrdTrade(trd *TRADEEXT, db *sql.DB, i int) {
	insertStatement := fmt.Sprintf("insert into trd_trade ( trd_recordno, trd_glosstraderef, trd_versiono, trd_origin, trd_tradetype, trd_settlementstatus, trd_tradestatus, trd_originversion) values ( %s,%s,%s,'%s','%s','%s','%s',%s);",
		trd.TRADE[i].Recordno,
		trd.TRADE[i].Glosstraderef,
		trd.TRADE[i].Versiono,
		trd.TRADE[i].Origin,
		trd.TRADE[i].Tradetype,
		trd.TRADE[i].Settlementstatus,
		trd.TRADE[i].Tradestatus,
		trd.TRADE[i].Originversion)
	_, err := db.Exec(insertStatement)
	if err != nil {
		fmt.Printf("error %s\n", err)
	}
}

func execTrdExternalRef(trd *TRADEEXT, db *sql.DB, index int) {
	for j := 0; j < len(trd.TRADE[index].EXTERNALREF); j++ {
		insertStatement := fmt.Sprintf("insert into trd_external_ref (trd_recordno, ext_ref_extreftype, ext_ref_extref) values (%s,'%s','%s');",
			trd.TRADE[index].Recordno,
			trd.TRADE[index].EXTERNALREF[j].Extreftype,
			trd.TRADE[index].EXTERNALREF[j].Extref)
		_, err := db.Exec(insertStatement)
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
}

func execTrdSettlement(trd *TRADEEXT, db *sql.DB, index int) {
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
		_, err := db.Exec(insertStatement)
		if err != nil {
			fmt.Printf("error %s\n", err)
		}
	}
}

func execInstruments(inst *INSTEXT, db *sql.DB, i int) {
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
