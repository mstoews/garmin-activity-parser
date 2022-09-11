-- name: ListTradeAmount :many
SELECT *
FROM trd_amount;

-- name: ListTrades :many
SELECT *
FROM trd_trade
WHERE trd_versiono = $1
ORDER BY trd_recordno
LIMIT $2 OFFSET $3;


-- name: GetTrade :one
SELECT *
FROM trd_trade
WHERE trd_recordno = $1
ORDER BY trd_recordno
LIMIT 1;


-- name: ListAllTrades :many
SELECT *
FROM trd_trade
ORDER BY trd_recordno
LIMIT 1000;

-- name: ListTradeEvent :many
SELECT *
FROM trd_event
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeEventNarrative :many
SELECT *
FROM trd_event_narrative
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeExternalRef :many
SELECT *
FROM trd_external_ref
WHERE trd_recordno = $1
ORDER BY trd_recordno;


-- name: ListTradeInstExt :many
SELECT *
FROM trd_inst_ext
WHERE trd_recordno = $1
ORDER BY trd_recordno;


-- name: ListTradeInstruction :many
SELECT *
FROM trd_instruction
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeInstructionEffect :many
SELECT *
FROM trd_instruction_effect
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeInstrument :many
SELECT *
FROM trd_instrument
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeJournal :many
SELECT *
FROM trd_journal
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeLink :many
SELECT *
FROM trd_link
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeParty :many
SELECT *
FROM trd_party
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradePartyDriver :many
SELECT *
FROM trd_party_driver
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeProcessing :many
SELECT *
FROM trd_processing
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeProcessingEvent :many
SELECT *
FROM trd_processing_event
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeRate :many
SELECT *
FROM trd_rate
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeRefDate :many
SELECT *
FROM trd_ref_date
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeSettlement :many
SELECT *
FROM trd_settlement
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeCode :many
SELECT *
FROM trd_trade_code
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeNarrative :many
SELECT *
FROM trd_trade_narrative
WHERE trd_recordno = $1
ORDER BY trd_recordno;

-- name: ListTradeDriver :many
SELECT *
FROM trd_driver
WHERE trd_recordno = $1
ORDER BY trd_recordno;


-- name: CreateTrade :one
INSERT INTO trd_trade (
        trd_recordno,
        trd_glosstraderef,
        trd_versiono,
        trd_origin,
        trd_tradetype,
        trd_settlementstatus,
        trd_tradestatus,
        trd_originversion
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: CreateTrdAmount :one
insert into trd_amount (trd_recordno,
                        trd_charge_levy_type_p2k, 
                        trd_charge_levy_instr_p2k, 
                        trd_charge_discount_wil,
                        trd_charge_levy_qty_p2k, 
                        trd_charge_levyrate_p2k)
values ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateTrdDriver :one
insert into trd_driver (trd_recordno, 
                        trd_drivertype, 
                        trd_drivercode)
values ($1, $2, $3)
RETURNING *;

-- name: CreateTrdEvent :one
insert into trd_event (trd_recordno, 
                       trd_eventtype, 
                       trd_eventdate, 
                       trd_eventdateto, 
                       trd_entrydatetime, 
                       trd_eventcode,
                       trd_exceptiontype, 
                       trd_expirydate)
values ($1,$2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: CreateTrdEventNarrative :one
insert into trd_event_narrative (trd_recordno, 
                                 trd_narrative_code, 
                                 trd_seqno, 
                                 trd_narrative)
values ($1,$2, $3, $4)
RETURNING *;


-- name: CreateTrdExternalRef :one
insert into trd_external_ref (trd_recordno,
                              ext_ref_extreftype, 
                              ext_ref_extref)
values ($1,$2, $3)
RETURNING *;

-- name: CreateInstExt :one
insert into trd_inst_ext (trd_recordno,
                          trd_service, 
                          trd_extref)
values ($1,$2, $3)
RETURNING *;

insert into trd_instruction (trd_recordno, 
                             trd_procaction, 
                             trd_destination, 
                             trd_procstatus, 
                             trd_recordidentifier,
                             trd_recordversion, 
                             trd_instformat, 
                             trd_tradeparty, 
                             trd_partyref, 
                             trd_deliverytype,
                             trd_addresscode, 
                             trd_servicestatus, 
                             trd_noofcopies, 
                             trd_duedatetime)
values ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING *;


-- name: CreateInstructionEffect :one
insert into trd_instruction_effect (trd_recordno, 
                                    trd_eventtype, 
                                    trd_eventdate, 
                                    trd_eventdateto, 
                                    trd_entrydatetime,
                                    trd_eventcode, 
                                    trd_exceptiontype, 
                                    trd_expirydate)
values ($1,$2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: CreateInstrument :one
insert into trd_instrument (trd_recordno, 
                            trd_instrtype, 
                            trd_p2000instrref, 
                            trd_instrreftype, 
                            trd_instrref,
                            trd_longname, 
                            trd_quantity)
values ($1,$2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: CreateParty :one
insert into trd_party (trd_recordno, 
                       trd_trade_party, 
                       trd_partyref, 
                       trd_partyref_type_text, 
                       trd_ext_partyref,
                       trd_longname)
values ($1,$2, $3, $4, $5, $6)
RETURNING *;

-- name: CreatePartyDriver :one
insert into trd_party_driver (trd_recordno, 
                              trd_trade_party, 
                              trd_driver_type, 
                              trd_driver_code)
values ($1,$2, $3, $4)
RETURNING *;


-- name: CreateProcessing :one
insert into trd_processing (trd_recordno,
                            trd_proc_alias, 
                            trd_proc_action, 
                            trd_due_datetime, 
                            trd_done_datetime)
values ($1,$2, $3, $4, $5)
RETURNING *;

-- name: CreateProcessingEvent :one
insert into trd_processing_event (trd_recordno, 
                                  trd_eventtype, 
                                  trd_eventdate, 
                                  trd_eventdateto, 
                                  trd_entrydatetime,
                                  trd_eventcode, 
                                  trd_exceptiontype, 
                                  trd_expirydate)
values ($1,$2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: CreateRate :one
insert into trd_rate (trd_recordno, 
                      trd_charge_levy_type, 
                      trd_actual_charge, 
                      trd_amount_type, 
                      trd_rate_instrref_type,
                      trd_rate_instrref, 
                      trd_rate_instrid, 
                      trd_rate_entered, 
                      trd_charge_rate, 
                      trd_mult_divind)
values ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;


-- name: CreateRefDate :one
insert into trd_ref_date (trd_recordno, 
                          datetype, 
                          datewil, 
                          refdatetime, 
                          dateversionuser)
values ($1,$2, $3, $4, $5)
RETURNING *;

-- name: CreateSettlement :one
insert into trd_settlement (trd_recordno, 
                            trd_deliverytype, 
                            trd_settleeventinstr, 
                            trd_settleterms, 
                            trd_originalqty,
                            trd_openqty, 
                            trd_settledate, 
                            trd_delrecind, 
                            trd_settlestatus, 
                            trd_tradestatus,
                            trd_settlenarrative1, 
                            trd_settlenarrative2, 
                            trd_settlenarrative3, 
                            trd_settlenarrative4,
                            trd_settlenarrative5, 
                            trd_settlenarrative6, 
                            trd_settlenarrative7, 
                            trd_settlenarrative8,
                            trd_dompaliaswil, 
                            trd_dompaliasdesc, 
                            trd_dompdepottypewil, 
                            trd_dompdaccwil,
                            trd_dompservicewil, 
                            trd_secpaliaswil, 
                            trd_secpservicewil)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,$21, $22, $23, $24, $25)
RETURNING *;

-- name: CreateTradeOne :one
insert into trd_trade_code (trd_recordno, 
                            trd_ext, 
                            trd_radeclass, 
                            trd_radecode)
values ($1, $2, $3, $4)
RETURNING *;

-- name: CreateTradeNarrative :one
insert into trd_trade_narrative (trd_recordno, 
                                 trd_narrative_code, 
                                 trd_seqno, 
                                 trd_narrative)
values ($1, $2, $3, $4)
RETURNING *;


-- name: CreateJournal :one
insert into trd_journal (trd_recordno, 
                         trd_accounts_company, 
                         trd_journal_type, 
                         trd_posting_type, 
                         trd_journal_no,
                         trd_procaction)
values ($1,$2, $3, $4, $5, $6)
RETURNING *;

-- name: CreateLink :one
insert into trd_link (trd_recordno, 
                      trd_link_type_wil, 
                      trd_main_record_no_wil, 
                      trd_sub_recordno_wil, 
                      trd_link_qty_wil,
                      trd_link_status_wil)
values ($1,$2, $3, $4, $5, $6)
RETURNING *;

