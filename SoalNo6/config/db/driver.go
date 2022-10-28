package db

import (
	"fmt"
	"time"

	l "Github.com/Dert12318/TNT-Telkom.git/SoalNo6/config/log"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models"
	"Github.com/Dert12318/TNT-Telkom.git/SoalNo6/models/contract"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
	l  *l.LogCustom
}

func NewDB(conf models.ServerConfig, isDbLog bool) *Database {
	var DB *gorm.DB
	var err error

	var host, user, password, name, port string

	log := l.NewLogCustom(conf)

	defer func() {
		if r := recover(); r != nil {
			log.Error(errors.New("recover"), "config/db: recover from error db init", "",
				nil, nil, nil)
		}
	}()

	// check DB version
	if isDbLog {
		host = conf.DBConfig.HostLogDb
		port = conf.DBConfig.PortLogDb
		user = conf.DBConfig.UserLogDb
		password = conf.DBConfig.PasswordLogDb
		name = conf.DBConfig.NameLogDb
	} else {
		host = conf.DBConfig.Host
		port = conf.DBConfig.Port
		user = conf.DBConfig.User
		password = conf.DBConfig.Password
		name = conf.DBConfig.Name
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, name, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "config/db: gorm open connect", nil)
	}

	dbSQL, err := DB.DB()
	if err != nil {
		log.Fatal(err, "config/DB: gorm open connect", nil)
	}

	//Database Connection Pool
	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	err = dbSQL.Ping()
	if err != nil {
		log.Fatal(err, "config/DB: can't ping the DB, WTF", nil)
	} else {
		go doEvery(10*time.Minute, pingDb, DB, log)
		return &Database{
			DB: DB,
			l:  log,
		}
	}

	return &Database{
		DB: DB,
		l:  log,
	}
}

func doEvery(d time.Duration, f func(*gorm.DB, *l.LogCustom), x *gorm.DB, y *l.LogCustom) {
	for _ = range time.Tick(d) {
		f(x, y)
	}
}

func pingDb(db *gorm.DB, log *l.LogCustom) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Error(err, "config/db: can't ping the db, WTF", "", nil, nil, nil)
	}

	err = dbSQL.Ping()
	if err != nil {
		log.Error(err, "config/db: can't ping the db, WTF", "", nil, nil, nil)
	}
}

func (d *Database) AutoMigrate(schemas ...interface{}) {

	for _, schema := range schemas {

		if err := d.DB.AutoMigrate(schema); err != nil {
			d.l.Error(errors.New(contract.ErrGeneralError), "", "", nil, nil, nil)
		}
	}
}

func (db *Database) DropTable(schemas ...interface{}) error {
	for _, schema := range schemas {

		if err := db.DB.Migrator().DropTable(schema); err != nil {
			db.l.Error(errors.New(contract.ErrGeneralError), "", "", nil, nil, nil)
			return err
		}
	}
	return nil
}
