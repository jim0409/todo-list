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
	Close() error

	CreateNotes(string, string) error
	ReadAllNotes() (map[uint]interface{}, error)
	UpdateNotes(string, string, string) error
	DeleteNote(string) error

	ReadNoteByPage(int, int) (map[uint]interface{}, error)
	CountPage(int64) (int64, error)

	// auth -
	// Check Policy sub, pub, act
	CheckPolicy(string, string, string) (bool, error)

	// account -
	LoginUser(string, string) (map[string]interface{}, error)
	CreateUser(string, string, string, string) error
	GetUserInfo(string) (map[string]interface{}, error)

	UpdateUserInfos(map[string]interface{}) error

	// need to verify auth again ...
	UpdateUserPassword(string, string) error

	// admin action
	ChangeUserStatus(string, uint8) error
	UpdateUserRole(string, string) error
	DeleteUser(string) error
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
		c.DbLogMode,
		c.DbMaxConnect,
		c.DbIdleConnect,
	)

	return mysqldb.StartMySQLDB()
}
