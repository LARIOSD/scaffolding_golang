package drivers

import (
	"context"
	"genesis/pos/infreastructure/db/postgres"
	"github.com/jackc/pgx/v5"
	"log"
)

type ConfigConnectionPgx struct {
	UrlToConnect string
}

func (configConnection *ConfigConnectionPgx) GetDatabaseConnectionPgx() *postgres.ConfigDatabaseConnectionPgxDriver {

	connection, err := pgx.Connect(context.Background(), configConnection.UrlToConnect)
	if err != nil {
		log.Fatal(err)
	}
	return &postgres.ConfigDatabaseConnectionPgxDriver{PgxConn: connection}
}
