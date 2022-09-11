package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

// insert into trd_amount (trd_recordno, trd_charge_levy_type_p2k, trd_charge_levy_instr_p2k, trd_charge_discount_wil,
// trd_charge_levy_qty_p2k, trd_charge_levyrate_p2k)
// values ();

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
