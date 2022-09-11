-- name: ListInstruments :many
SELECT *
FROM instr_instruments
LIMIT 10000;


-- name: GetInstrumentsByRef :one
SELECT *
FROM instr_instruments WHERE instr_instref = $1
ORDER BY instr_instref;





