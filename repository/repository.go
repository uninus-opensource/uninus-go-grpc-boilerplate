package repository

import (
	"io"
	_interface "github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/interface"
)

type DBConf struct {
	URL, Port, Schema, User, Password string
	MaxOpenConn, MaxConnLifeTime      int
}

type RestConf struct {
	//Port string
	URL string
}

type MicroConf struct {
	DiscHost     []string
	SecurityHost string
	//Tls                     credentials.TransportCredentials
	NotificationServiceHost string
	NotificationSercicePort string
	Ca                      string
	Cer                     string
	Key                     string
}

type RedisConf struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

type Repo struct {
	DBReadWriter _interface.ReadWriter
	Redis        _interface.RedisCache
	//GPubsub      _interface.GPubsub
	//Tinode       _interface.Tinode
	//Microservice microservices.Microservices
	RestService _interface.Rest
	//KeyCloak    _interface.KeyCloak
	io.Closer
}
