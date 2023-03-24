package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
    db *sql.DB
}

func newAdapter(driverName, dataSourceName string) (*Adapter, error) {
    // connect
    db, err := sql.Open(driverName, dataSourceName)
    if err != nil {
        log.Fatalf("Db connection failur: %v", err)
    }

    // test db connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Db ping failur: %v", err)
    }

    return &Adapter{db: db}, nil
}

func (da Adapter) CloseConnection() {
    err := da.db.Close()
    if err != nil {
        log.Fatalf("Db close failur: %v", err)
    }
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
    currentTime := time.Now()
    formatedTime := currentTime.Format("2006-01-02 15:04:05")
    query := fmt.Sprintf(
        "INSERT INTO arith_history (date, answer, operation) VALUES (%s, %b, %s)",
        formatedTime, answer, operation,
    )

    _, err := da.db.Exec(query)
    if err != nil {
        return err
    }

    return nil
}
