-- name: ListParties :many
SELECT *
FROM pty_party
LIMIT 1000;



-- name: InsertParty :one
insert into pty_party (
    pty_partyref, 
    pty_category, 
    pty_longdesc, 
    pty_holiday, 
    pty_country, 
    pty_location, 
    pty_shrtdesc,
    pty_partynam1, 
    pty_partynam2, 
    pty_partynam3, 
    pty_active, 
    pty_verdat)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 )
RETURNING *;
