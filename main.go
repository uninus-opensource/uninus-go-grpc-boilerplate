package main

import (
	ssl "crypto/tls"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/sd"
	cfg "github.com/uninus-opensource/uninus-go-architect-common/config"
	run "github.com/uninus-opensource/uninus-go-architect-common/grcp"
	"github.com/uninus-opensource/uninus-go-architect-common/log"
	umvc "github.com/uninus-opensource/uninus-go-architect-common/microservice"
	ep "github.com/uninus-opensource/uninus-go-grpc-boilerplate/endpoint"
	pb "github.com/uninus-opensource/uninus-go-grpc-boilerplate/proto"
	"github.com/uninus-opensource/uninus-go-grpc-boilerplate/repository"
	svc "github.com/uninus-opensource/uninus-go-grpc-boilerplate/server"
	transport "github.com/uninus-opensource/uninus-go-grpc-boilerplate/transport/rpc"
	cs "github.com/uninus-opensource/uninus-go-grpc-boilerplate/utils/constant"
	shv "github.com/uninus-opensource/uninus-go-grpc-boilerplate/utils/sharevar"
	"go.elastic.co/apm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"os"
)

const (
	gCredential = "GOOGLE_APPLICATION_CREDENTIALS"
)

type conf struct {
	DiscHost          []string
	Tls               credentials.TransportCredentials
	TlsCert           ssl.Certificate
	IP                string
	Port              string
	Address           string
	DbHost            string
	DbPort            string
	DbName            string
	DbUser            string
	DbPwd             string
	DBMaxOpenConn     int
	DBMaxConnLifeTime int
	RedisHost         string
	RedisPort         string
	RedisPassword     string
	SecurityHost      string
	//GProjectId         string
	//GCredentialFile    string
	//SubsNum            int
	//TinodeHost         string
	//TinodePort         string
	//TopicCreateRoom    string
	//TopicCreateRoomNew string
	//TopicSavingRoom    string
	NofiticationHost string
	NotificationPort string
	Ca               string
	Cer              string
	Key              string
	//KeyCloakURL        string
	//KeyCloakRealm      string
}

func parseAndLoadConfigs() (*conf, error) {
	confFlag := flag.String("config", "", "config file")
	ipFlag := flag.String("host", "", "host ip")
	portFlag := flag.String("port", "", "host port")
	flag.Parse()

	var ok bool
	if len(*confFlag) == 0 {
		ok = cfg.AppConfig.LoadConfig()
	} else {
		ok = cfg.AppConfig.LoadConfigFile(*confFlag)
	}
	if !ok {
		return nil, errors.New("failed to load configuration")
	}

	discHost := cfg.GetA("discoveryhost", "")
	ip := cfg.Get("serviceip", "")
	port := cfg.Get("serviceport", "")

	if len(*ipFlag) > 0 {
		ip = *ipFlag
	}
	if len(*portFlag) > 0 {
		port = *portFlag
	}
	//ca := cfg.Get("capath", "")
	//cert := cfg.Get("certpath", "")
	//prikey := cfg.Get("keypath", "")
	//var tlsCert ssl.Certificate
	//var tls credentials.TransportCredentials
	//fmt.Println("capath", ca, "certpath", cert, "keypath", prikey)
	//if len(ca) > 0 && len(cert) > 0 && len(prikey) > 0 {
	//	var err error
	//	tls, _ = run.TLSCredentialFromFile(ca, cert, prikey, true)
	//	if err != nil {
	//		return nil, err
	//	}
	//	tlsCert, _ = ssl.LoadX509KeyPair(cert, prikey)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	gcredfile := cfg.Get("gcredfile", "")
	// fmt.Println(gcredfile)

	if len(gcredfile) > 0 {
		os.Setenv(gCredential, gcredfile)
	}

	config := &conf{
		DiscHost: discHost,
		//Tls:               tls,
		//TlsCert:           tlsCert,
		IP:                ip,
		Port:              port,
		Address:           fmt.Sprintf("%s:%s", ip, port),
		DbHost:            cfg.Get(cfg.DBhost, ""),
		DbPort:            cfg.Get(cfg.DBport, ""),
		DbName:            cfg.Get(cfg.DBname, ""),
		DbUser:            cfg.Get(cfg.DBuid, ""),
		DbPwd:             cfg.Get(cfg.DBpwd, ""),
		DBMaxOpenConn:     cfg.GetI("dbmaxopenconn", 50),
		DBMaxConnLifeTime: cfg.GetI("dbmaxlifetimeconn", 2),
		RedisHost:         cfg.Get(cfg.RDhost, ""),
		RedisPort:         cfg.Get(cfg.RDport, ""),
		RedisPassword:     cfg.Get(cfg.RDpassword, ""),
		//GProjectId:        cfg.Get(cfg.GProjectId, ""),
		//SubsNum:           cfg.GetI(cfg.SubsNum, 1),
		//TinodeHost:        cfg.Get("tinodeHost", ""),
		//TinodePort:        cfg.Get("tinodePort", ""),
		//TopicCreateRoom:   cfg.Get("topiccreateroom", ""),
		//TopicSavingRoom:   cfg.Get("topicsavingroom", "stage.order.saving.room"),
		//NofiticationHost:  cfg.Get("notificationservicehost", "http://stg-bbone-notification-token.gcp.bluebird.id"),
		//NotificationPort:  cfg.Get("notificationserviceport", "443"),
		//Key:               cfg.Get("keypath", ""),
		//Ca:                cfg.Get("capath", ""),
		//Cer:               cfg.Get("certpath", ""),
		//KeyCloakURL:       cfg.Get("keycloakUrl", ""),
		//KeyCloakRealm:     cfg.Get("keycloakRealm", ""),
	}
	//fmt.Println("config: ", config.NofiticationHost)
	//fmt.Println("post :", config.NotificationPort)
	//fmt.Println("tls :", config.Key)
	//fmt.Println("tls :", config.Ca)
	//fmt.Println("tls :", config.Cer)

	return config, nil

}

func (cnf *conf) createRegistrar() (sd.Registrar, error) {
	if len(cnf.DiscHost) > 0 {
		registrar, err := umvc.ServiceRegistry(cnf.DiscHost, cs.ServiceID, cnf.Address, shv.Logger)
		if err != nil {
			return nil, err
		}
		registrar.Register()
		return registrar, nil
	}
	return nil, nil
}

func (cnf *conf) UserServiceServer() (*pb.UserServiceServer, *repository.Repo, error) {
	repoConf := repository.RepoConfigs{
		//MicroConf: repository.MicroConf{
		//	DiscHost:     cnf.DiscHost,
		//	SecurityHost: cnf.SecurityHost,
		//	//NotificationServiceHost: "",
		//	//NotificationSercicePort: "",
		//	Ca:  cnf.Ca,
		//	Cer: cnf.Cer,
		//	Key: cnf.Key,
		//},
		DBConf: repository.DBConf{
			URL:             cnf.DbHost,
			Port:            cnf.DbPort,
			Schema:          cnf.DbName,
			User:            cnf.DbUser,
			Password:        cnf.DbPwd,
			MaxOpenConn:     cnf.DBMaxOpenConn,
			MaxConnLifeTime: cnf.DBMaxConnLifeTime,
		},
		RedisConf: repository.RedisConf{
			RedisHost:     cnf.RedisHost,
			RedisPort:     cnf.RedisPort,
			RedisPassword: cnf.RedisPassword,
		},
		RestConf: repository.RestConf{
			URL: cnf.NofiticationHost,
		},
	}

	newRepo, err := repository.NewUserServiceRepo(repoConf)
	if err != nil {
		level.Error(shv.Logger).Log(log.LogError, err.Error())
		return nil, nil, err
	}

	userService := svc.NewUserServer(*newRepo)

	endpoint, err := ep.NewUserEndpoint(userService)

	server := transport.NewGRPCUserServer(endpoint, shv.Logger)

	return &server, newRepo, nil
}

func (cnf *conf) initGRPCServer(service *pb.UserServiceServer) *grpc.Server {
	grpcServer := grpc.NewServer(
		run.DefaultServerOptions(shv.Logger)...,
	)
	pb.RegisterUserServiceServer(grpcServer, *service)
	reflection.Register(grpcServer)
	return grpcServer

}
func main() {
	shv.Logger = log.StackDriverLogger()
	shv.TracerAPM = apm.DefaultTracer

	config, err := parseAndLoadConfigs()
	if err != nil {
		level.Error(shv.Logger).Log(log.LogError, err.Error())
		return
	}

	// create service registry
	registrar, err := config.createRegistrar()
	if err != nil {
		level.Error(shv.Logger).Log(log.LogError, err.Error())
		return
	}

	if registrar != nil {
		defer registrar.Deregister()
	}

	serviceServer, _, err := config.UserServiceServer()
	if err != nil {
		level.Error(shv.Logger).Log(log.LogError, err.Error())
		return
	}
	//defer repo.Close()

	grpcServer := config.initGRPCServer(serviceServer)

	go run.ServeGRPCAndHTTP(
		config.Address, config.Port, grpcServer,
		pb.RegisterUserServiceHandlerFromEndpoint,
		nil, ssl.Certificate{}, shv.Logger,
		run.DefaultHTTPOption(), run.DefaultHTTPHandler,
	)

	level.Info(shv.Logger).Log("[RUNNING]", "Server started")

	umvc.OnShutdown(func() {
		grpcServer.GracefulStop()
	})

}
