module github.com/calvarado2004/go-testing

go 1.20

require (
	github.com/alexedwards/scs/v2 v2.5.0
	github.com/calvarado2004/go-testing/pkg/db v0.0.0-20230303055409-c68a48658271
	github.com/go-chi/chi/v5 v5.0.8
	github.com/jackc/pgconn v1.14.0
	github.com/jackc/pgx/v4 v4.18.1
)

require (
	github.com/calvarado2004/go-testing/go-webapp/webapp/pkg/data v0.0.0-20230303055409-c68a48658271 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)

replace (
	github.com/calvarado2004/go-testing/pkg/data => ./pkg/data
	github.com/calvarado2004/go-testing/pkg/db => ./pkg/db
)
