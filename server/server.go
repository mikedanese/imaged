package server

import (
	"net"

	"github.com/boltdb/bolt"
	"github.com/golang/glog"
	"github.com/mikedanese/imaged/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewServer(c Config) *Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			glog.Infof("ctx: %+v", ctx)
			glog.Infof("info: %+v", info)
			return handler(ctx, req)
		}),
	)
	serv := Server{
		s: s,
		c: c,
	}

	api.RegisterImagesServer(s, &imagesServerImpl{s: &serv})
	api.RegisterBlobsServer(s, &blobsServerImpl{s: &serv})
	api.RegisterBundlesServer(s, &bundlesServerImpl{s: &serv})

	return &serv
}

type Config struct{}

type Server struct {
	db *bolt.DB
	s  *grpc.Server
	c  Config
}

func (s *Server) Serve() error {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return err
	}
	if err := CreateBuckets(db, resourceBuckets); err != nil {
		return err
	}
	s.db = db

	ds, err := net.Listen("unix", "foo.sock")
	if err != nil {
		return err
	}
	return s.s.Serve(ds)
}

func (s *Server) Stop() {
	s.db.Close()
	s.s.Stop()
}

type imagesServerImpl struct {
	s *Server
}

type bundlesServerImpl struct {
	s *Server
}

type blobsServerImpl struct {
	s *Server
}

//go:generate go-template -in crud_gen.go.template -out crud_gen.go -cfg {"Types":["Image","Bundle","Blob"]}
