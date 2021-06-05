package mysqldb

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	DBLogMode        int
	DBMaxConnection  int
	DBIdleConnection int
	DBUri            string
}

func LoadMySQLDBConfig(dbName, dbHost, dbPort, dbUsr, dbPassword string, dbLogMode int, dbMaxConnection, dbIdleConnection int) {
	mysqlDbConfig = &MySQLConfig{
		DBName:           dbName,
		DBHost:           dbHost,
		DBPort:           dbPort,
		DBUsr:            dbUsr,
		DBPassword:       dbPassword,
		DBLogMode:        dbLogMode,
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
	d, err := db.DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

type MySQLDBAccessObject interface {
	Close() error

	NoteImp
}

func initMySqlDB(c *MySQLConfig) (MySQLDBAccessObject, error) {
	var db *gorm.DB
	var err error

	if db, err = gorm.Open(mysql.Open(c.DBUri), &gorm.Config{}); err != nil {
		return nil, fmt.Errorf("Connection to MySQL DB error : %v", err)
	}

	// config db open & idle connection nums
	d, err := db.DB()
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(c.DBMaxConnection)
	d.SetMaxIdleConns(c.DBIdleConnection)

	if err = d.Ping(); err != nil {
		return nil, fmt.Errorf("Ping MySQL db error : %v", err)
	}

	// 在 gorm.io/gorm 後 db.LogMode(c.DBLogEnable) 改為按照 log 等級
	db.Logger.LogMode(logger.LogLevel(c.DBLogMode)) // Silent=1, Error=2, Warn=3, Info=4
	// db.SingularTable(true) // 改用 db open mysql.Config{} 取代 .. singular 用意，對於table 命名做格式化，tables -> table

	if err = db.AutoMigrate(&NoteTable{}); err != nil {
		return nil, err
	}

	return &mysqlDBObj{DB: db}, nil
}
