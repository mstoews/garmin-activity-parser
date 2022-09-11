package main

import (
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "secret", "the database password")
	port     *int = flag.Int("port", 32342, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", "root", "the database user")
	database      = flag.String("database", "postgres", "database name")
	source        = flag.String("source", "trade", "Source type (inst, party, trade)")
	filename      = flag.String("filename", "./xml/trades.xml", "XML file name")
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
type PARTYEXT struct {
	XMLName xml.Name `xml:"PARTYEXT"`
	Text    string   `xml:",chardata"`
	PARTY   []struct {
		Text      string `xml:",chardata"`
		Partyref  string `xml:"partyref,attr"`
		Category  string `xml:"category,attr"`
		Longdesc  string `xml:"longdesc,attr"`
		Holiday   string `xml:"holiday,attr"`
		Country   string `xml:"country,attr"`
		Location  string `xml:"location,attr"`
		Shrtdesc  string `xml:"shrtdesc,attr"`
		Partynam1 string `xml:"partynam1,attr"`
		Partynam2 string `xml:"partynam2,attr"`
		Partynam3 string `xml:"partynam3,attr"`
		Active    string `xml:"active,attr"`
		Verdat    string `xml:"verdat,attr"`
		PARTYASSC []struct {
			Text      string `xml:",chardata"`
			Assctype  string `xml:"assctype,attr"`
			Asscparty string `xml:"asscparty,attr"`
		} `xml:"PARTYASSC"`
		PARTYCLASS []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			Code  string `xml:"code,attr"`
		} `xml:"PARTYCLASS"`
		PARTYEXT []struct {
			Text    string `xml:",chardata"`
			Service string `xml:"service,attr"`
			Extref  string `xml:"extref,attr"`
		} `xml:"PARTYEXT"`
		PARTYFLG []struct {
			Text  string `xml:",chardata"`
			Class string `xml:"class,attr"`
			Code  string `xml:"code,attr"`
		} `xml:"PARTYFLG"`
		PARTYNARR []struct {
			Text      string `xml:",chardata"`
			Narrcode  string `xml:"narrcode,attr"`
			SeqNoWil  string `xml:"seq_no_wil,attr"`
			Narrative string `xml:"narrative,attr"`
		} `xml:"PARTYNARR"`
		PARTYADDR []struct {
			Text     string `xml:",chardata"`
			Addrcode string `xml:"addrcode,attr"`
			Contact  string `xml:"contact,attr"`
			Title    string `xml:"title,attr"`
			Addrlin1 string `xml:"addrlin1,attr"`
			Addrlin2 string `xml:"addrlin2,attr"`
			Addrlin3 string `xml:"addrlin3,attr"`
			Addrlin4 string `xml:"addrlin4,attr"`
			Addrlin5 string `xml:"addrlin5,attr"`
			Addrlin6 string `xml:"addrlin6,attr"`
			Postcode string `xml:"postcode,attr"`
			Dialcode string `xml:"dialcode,attr"`
			Phone    string `xml:"phone,attr"`
			Telexcou string `xml:"telexcou,attr"`
			Telex    string `xml:"telex,attr"`
			Ansback  string `xml:"ansback,attr"`
			Fax      string `xml:"fax,attr"`
			Email    string `xml:"email,attr"`
		} `xml:"PARTYADDR"`
		PARTYRATE struct {
			Text     string `xml:",chardata"`
			Ratetype string `xml:"ratetype,attr"`
			Rate     string `xml:"rate,attr"`
		} `xml:"PARTYRATE"`
		PARTYGL struct {
			Text     string `xml:",chardata"`
			Acctype  string `xml:"acctype,attr"`
			Acctcode string `xml:"acctcode,attr"`
		} `xml:"PARTYGL"`
	} `xml:"PARTY"`
}

func main() {
	start := time.Now()
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
		fmt.Printf(" source:%s\n", *source)
		fmt.Printf(" filename:%s\n", *filename)
	}

	fmt.Printf("Source:  %s\n", *source)
	fmt.Printf("File name:  %s\n", *filename)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *server, *port, *user, *password, *database)

	fmt.Printf("Connecting to database %s\n", psqlconn)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// Open our xmlFile
	switch *source {
	case "inst":
		processInstruments(*filename, db)
	case "party":
		processParty(*filename)
	case "trade":
		processTrade(*filename, db)
	default:
		panic("Invalid source")
	}
	log.Printf("main, execution time %s\n", time.Since(start))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func processParty(filename string) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)

		fmt.Printf("Successfully Opened %s", filename)
		// defer the closing of our xmlFile so that we can parse it later on
		defer xmlFile.Close()
	}
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var partyExt PARTYEXT
	xml.Unmarshal(byteValue, &partyExt)
	insertParty(&partyExt)
	// insertExternalRef(&partyExt)
}
