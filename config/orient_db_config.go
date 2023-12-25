package config

import orient "github.com/quux00/ogonori"

type OrientDBService interface {
	GetConnect() orient.DBSession
}

type OrientDBConfig struct {
}

func (o *OrientDBConfig) GetConnect() orient.DBSession {
	var db orient.DBConnection
	open, err := db.Open("demo", "document", "root", "admin")
	if err != nil {
		return nil
	}
	return open
}
