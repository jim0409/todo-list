package models

import (
	"todo-list/config"
	"todo-list/models/mysqldb"
)

/*
use facade design pattern to new db object
1. initDB
2. closeDB
*/

func RetriveMySqlDbAccessModel() MySqlImplement {
	return mysqldb.RetriveMySQLDBAccessObj()
}

type MySqlImplement interface {
	ExecSql(string) error
	Close() error

	CreateNotes(string, string) error
	ReadAllNotes() (map[string]string, error)
	UpdateNotes(string, string) error
	DeleteNotes(string) error
}

// InitDb init db
func InitDb(mysqlconf *config.DbConf) error {
	return initMySql(mysqlconf)
}

// Close close db
func Close() error {
	return RetriveMySqlDbAccessModel().Close()
}

func initMySql(c *config.DbConf) error {
	mysqldb.LoadMySQLDBConfig(
		c.DbName,
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbLogEnable,
		c.DbMaxConnect,
		c.DbIdleConnect,
	)

	return mysqldb.StartMySQLDB()
}