package rtdb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DSNMySQL struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type DSNTiDB struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

// Returns a new MySQL DB connection.
//
// Pass the DSN as an instance of rtdb.DSNMySQL struct.
func ConnectMySQL(DSNMySQL *DSNMySQL) (*gorm.DB, error) {
	//dsn: "username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		DSNMySQL.Username,
		DSNMySQL.Password,
		DSNMySQL.Host,
		DSNMySQL.Port,
		DSNMySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	return db, nil
}

func ConnectTiDB(DSNTiDB *DSNTiDB) *gorm.DB {
	//dsn: "username:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		DSNTiDB.Username,
		DSNTiDB.Password,
		DSNTiDB.Host,
		DSNTiDB.Port,
		DSNTiDB.Database,
	)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		db.Error = err
		return db
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(5)
	return db
}
