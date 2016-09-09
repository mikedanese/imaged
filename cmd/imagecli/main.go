package main

import (
	"log"
	"os"

	"google.golang.org/grpc/grpclog"

	"github.com/golang/glog"
	"github.com/mikedanese/imaged/client"
	"github.com/spf13/cobra"
)

func main() {
	grpclog.SetLogger(log.New(os.Stderr, "", log.LstdFlags))

	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(putCmd)

	rootCmd.Execute()
}

var getCmd = &cobra.Command{
	Use: "get [type]",
	Run: doCmd(func(c *client.Client, cmd *cobra.Command, args []string) {
		crudFor(os.Args[2]).List(c)
	}),
}
var putCmd = &cobra.Command{
	Use: "put [type] [value]",
	Run: doCmd(func(c *client.Client, cmd *cobra.Command, args []string) {
		crudFor(os.Args[2]).Put(c, os.Args[2], os.Args[3])
	}),
}

func doCmd(f func(c *client.Client, cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		c, err := client.New()
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		f(c, cmd, args)
	}
}

type CRUD interface {
	List(c *client.Client)
	Put(*client.Client, string, string)
	Delete(*client.Client, string)
}

func crudFor(t string) CRUD {
	c, ok := crudHelpers[t]
	if !ok {
		glog.Fatalf("no crud for %q", t)
	}
	return c
}

//go:generate go-template -in crud_gen.go.template -out crud_gen.go -cfg {"Types":["Image","Blob","Bundle"]}
