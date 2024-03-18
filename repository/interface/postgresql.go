package _interface

import (
	"database/sql"
	"io"
)

type ReadWriter interface {
	io.Closer
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error

	//GetListUser(ctx context.Context, params model.GetListUserRequest) (model.GetListUserResponse, error)
}
