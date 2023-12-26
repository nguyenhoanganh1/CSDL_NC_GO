package config

import (
	"gopkg.in/istreamdata/orientgo.v2"
	_ "gopkg.in/istreamdata/orientgo.v2/obinary"
	_ "net/http"
	_ "net/http/pprof"
)

var (
	address                          = "localhost:2424"
	databaseName                     = "demo"
	documentType orient.DatabaseType = "document"
	username                         = "root"
	password                         = "admin"
)

type OrientDBService interface {
	GetConnect() orient.DBSession
}

type OrientDBConfig struct {
}

func (o *OrientDBConfig) GetConnect() *orient.Database {
	client, err := orient.Dial(address)

	if err != nil {
		panic(err)
	}

	admin, err := client.Auth(username, password)
	if err != nil {
		panic(err)
	}

	ok, err := admin.DatabaseExists(databaseName, orient.Persistent)
	if err != nil {
		panic(err)
	}

	if !ok {
		// There are 2 options
		// 1. orient.GraphDB - graph database
		// 2. orient.DocumentDB - document database
		err = admin.CreateDatabase(databaseName, orient.GraphDB, orient.Persistent)
		if err != nil {
			panic(err)
		}
	}

	database, err := client.Open(databaseName, documentType, username, password)
	if err != nil {
		panic(err)
	}

	defer func(database *orient.Database) {
		err := database.Close()
		if err != nil {

		}
	}(database)

	return database
}
