package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/mikedanese/imaged/server"
)

func main() {
	flag.Set("alsologtostderr", "true")

	glog.Infof("starting server")
	server.NewServer(server.Config{}).Serve()
}
