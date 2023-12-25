package config

import orient "github.com/quux00/ogonori"

type OrientDBService interface {
	GetConnect() orient.DBSession
}

type OrientDBConfig struct {
	host         string              `default:"localhost"`
	port         string              `default:"2424"`
	databaseName string              `default:"demo"`
	documentType orient.DatabaseType `default:"document"`
	username     string              `default:"root"`
	password     string              `default:"root"`
}

func (o *OrientDBConfig) GetConnect() orient.DBSession {
	var db orient.DBConnection
	open, err := db.Open(o.databaseName, o.documentType, o.username, o.password)
	if err != nil {
		return nil
	}
	return open
}
