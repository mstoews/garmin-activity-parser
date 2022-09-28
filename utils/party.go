package utils

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	_ "github.com/lib/pq"
)


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


func insertParty(pty *PARTYEXT) {
	for i := 0; i < len(pty.PARTY); i++ {
		fmt.Printf("insert into pty_party (pty_partyref, pty_category, pty_longdesc, pty_holiday, pty_country, pty_location, pty_shrtdesc, pty_partynam1, pty_partynam2, pty_partynam3, pty_active, pty_verdat) values ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');\n",
			pty.PARTY[i].Partyref,
			pty.PARTY[i].Category,
			pty.PARTY[i].Longdesc,
			pty.PARTY[i].Holiday,
			pty.PARTY[i].Country,
			pty.PARTY[i].Location,
			pty.PARTY[i].Shrtdesc,
			pty.PARTY[i].Partynam1,
			pty.PARTY[i].Partynam2,
			pty.PARTY[i].Partynam3,
			pty.PARTY[i].Active,
			pty.PARTY[i].Verdat)
	}

}

func ProcessParty(filename string) {
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

