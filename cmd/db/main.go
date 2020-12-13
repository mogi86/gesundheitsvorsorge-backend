package main

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (
	user     string
	password string
	host     string
	port     string
	DBName   string
)

// main run db migration
// see: https://github.com/golang-migrate/migrate/blob/master/database/mysql/README.md
func main() {
	intPort, err := strconv.Atoi(port)
	if err != nil {
		logrus.Errorf("port(%s) convert failed: %+v", port, err)
		return
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true", user, password, host, intPort, DBName))
	if err != nil {
		logrus.Errorf("DB open failed: %+v", err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logrus.Errorf("get driver failed: %+v", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		logrus.Errorf("get new migration instance failed: %+v", err)
		return
	}

	// Step はversionではない。例えばstep=1を指定した場合マイグレーションバージョンが1つ上がる。downはその逆です。
	// MEMO: Step isn't version.
	//       For example, if you specify step=1, the migration version goes up by 1, and vice versa for down.
	//err = m.Steps(intVersion)

	err = m.Up()
	if err != nil {
		logrus.Errorf("migration failed: %+v", err)
		return
	}
}