package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	shv "github.com/uninus-opensource/uninus-go-grpc-boilerplate/utils/sharevar"
	_interface "github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/interface"
	_ "github.com/lib/pq"
	"github.com/go-kit/kit/log/level"
)

const (
	dbMsgDataNotFound = "Data Not Found"

)

const (
	updateXX = `UPDATE XX`
	readXX   = `SELECT XX `
)

// dbReadWriter is a struct having the sql db parameter
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(
	url string, port string,
	schema string, user string,
	password string, maxOpenConn int,
	maxConnLifeTime int,
) (_interface.ReadWriter, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", url, port, user, password, schema)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		level.Error(shv.Logger).Log(fmt.Sprintf("%s", err.Error()))
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxOpenConn)
	db.SetConnMaxLifetime(time.Duration(maxConnLifeTime) * time.Second)

	return &dbReadWriter{db: db}, nil
}

func closeRows(rs *sql.Rows) {
	if rs != nil {
		if err := rs.Close(); err != nil {
			_ = level.Error(shv.Logger).Log(fmt.Sprintf("error while closing result set %+v", err.Error()))
		}
	}
}

func rollbackTx(tx *sql.Tx) {
	if tx == nil {
		return
	}

	if err := tx.Rollback(); err != nil {
		// _ = level.Error(shv.Logger).Log(fmt.Sprintf("error while rolling back transaction %+v", err.Error()))
	}
}

// Close is used for closing the sql connection
func (rw *dbReadWriter) Close() error {
	if rw.db != nil {
		if err := rw.db.Close(); err != nil {
			return err
		}
		rw.db = nil
	}

	return nil
}

func (rw *dbReadWriter) Begin() (*sql.Tx, error) {
	return rw.db.Begin()
}

func (rw *dbReadWriter) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (rw *dbReadWriter) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

