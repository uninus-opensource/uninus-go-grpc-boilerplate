package repository

import (
	_interface "github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/interface"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/postgresql"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository/redis"
	"io"
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

type RepoConfigs struct {
	MicroConf
	DBConf        DBConf
	DBReplicaConf DBConf
	RedisConf
	RestConf RestConf
}

func NewUserServiceRepo(rc RepoConfigs) (*Repo, error) {
	readWriter, err := postgresql.NewDBReadWriter(rc.DBConf.URL, rc.DBConf.Port, rc.DBConf.Schema, rc.DBConf.User,
		rc.DBConf.Password, rc.DBConf.MaxOpenConn,
		rc.DBConf.MaxConnLifeTime,
	)

	if err != nil {
		return nil, err
	}

	redisC, err := redis.NewRedis(rc.RedisHost, rc.RedisPort, rc.RedisPassword)
	if err != nil {
		return nil, err
	}

	// restService := rest.NewRestClient(rc.RestConf.URL)

	return &Repo{
		DBReadWriter: readWriter,
		Redis:        redisC,
		RestService:  nil,
		Closer:       nil,
	}, nil

}
