package infrastructure

import (
	"genesis/pos/infreastructure/db/mongo"
	"genesis/pos/infreastructure/db/postgres"
)

type DatabaseConnectionMongoInterface interface {
	GetConnectionMongo() *mongo.ConfigDatabaseConnectionMongoDriver
}
type DatabaseConnectionPgxInterface interface {
	GetDatabaseConnectionPgx() *postgres.ConfigDatabaseConnectionPgxDriver
}
