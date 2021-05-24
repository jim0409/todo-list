package mysqldb

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	mysqlDbConfig *MySQLConfig
	mysqlDb       MySQLDBAccessObject
	once          sync.Once
)

type MySQLConfig struct {
	DBName           string
	DBHost           string
	DBPort           string
	DBUsr            string
	DBPassword       string
	DBLogEnable      bool
	DBMaxConnection  int
	DBIdleConnection int
	DBUri            string
}

func LoadMySQLDBConfig(dbName, dbHost, dbPort, dbUsr, dbPassword string, dbLogEnable bool, dbMaxConnection, dbIdleConnection int) {
	mysqlDbConfig = &MySQLConfig{
		DBName:           dbName,
		DBHost:           dbHost,
		DBPort:           dbPort,
		DBUsr:            dbUsr,
		DBPassword:       dbPassword,
		DBLogEnable:      dbLogEnable,
		DBMaxConnection:  dbMaxConnection,
		DBIdleConnection: dbIdleConnection,
		// DBUri: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUri: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", // 會將UTC-time轉成當地時間...自動加8小時
			dbUsr, dbPassword, dbHost, dbPort, dbName,
		),
	}
}

func RetriveMySQLDBAccessObj() MySQLDBAccessObject {
	once.Do(func() {
		mysqlDb = &mysqlDBObj{}
	})
	return mysqlDb
}

func StartMySQLDB() error {
	var err error
	mysqlDb = RetriveMySQLDBAccessObj()
	mysqlDb, err = initMySqlDB(mysqlDbConfig)

	return err
}

type mysqlDBObj struct {
	DB *gorm.DB
}

func (db *mysqlDBObj) Close() error {
	return db.DB.Close()
}

func (db *mysqlDBObj) ExecSql(sqlStr string) error {
	return nil
}

type MySQLDBAccessObject interface {
	ExecSql(string) error
	Close() error

	CreateNotes(string, string) error
	ReadAllNotes() (map[string]string, error)
	UpdateNotes(string, string) error
	DeleteNotes(string) error
}

func initMySqlDB(c *MySQLConfig) (MySQLDBAccessObject, error) {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open("mysql", c.DBUri); err != nil {
		return nil, fmt.Errorf("Connection to MySQL DB error : %v", err)
	}

	db.DB().SetMaxOpenConns(c.DBMaxConnection)
	db.DB().SetMaxIdleConns(c.DBIdleConnection)

	if err = db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("Ping MySQL db error : %v", err)
	}

	db.LogMode(c.DBLogEnable)
	db.SingularTable(true)
	db.AutoMigrate(&NoteTable{})

	return &mysqlDBObj{DB: db}, nil
}
