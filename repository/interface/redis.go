package _interface

import (
	"io"
)

type RedisCache interface {
	io.Closer
}
