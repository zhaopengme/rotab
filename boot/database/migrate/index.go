package migrate

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/os/glog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
)

func AutoMigrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/sqlite",
	}
	db, err := sql.Open("sqlite3", "./rotab.db")
	if err != nil {
		glog.Fatal(err)
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
