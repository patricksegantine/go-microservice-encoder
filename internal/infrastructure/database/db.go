package database

import (
	"encoder/internal/domain/entities"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DbType        string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	d.Db, err = gorm.Open(d.DbType, d.Dsn)

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&entities.Video{}, &entities.Job{})

		d.Db.Model(entities.Job{}).AddForeignKey(
			"video_id",
			"videos (id)",
			"CASCADE",
			"CASCADE",
		)
	}

	return d.Db, nil
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbType = "sqlite3"
	dbInstance.Dsn = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %v", err)
	}

	return connection
}
