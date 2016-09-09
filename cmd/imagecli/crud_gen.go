// auto-generated
// do not modify this file by hand

package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/mikedanese/imaged/api"
	"github.com/mikedanese/imaged/client"
	"golang.org/x/net/context"
)

var (
	crudHelpers = map[string]CRUD{

		"image": &imagesCRUD{},

		"blob": &blobsCRUD{},

		"bundle": &bundlesCRUD{},
	}
)

type imagesCRUD struct{}

func (h *imagesCRUD) List(c *client.Client) {
	es, err := c.Images().List(context.Background(), &api.Void{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, es); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *imagesCRUD) Put(c *client.Client, k, v string) {
	var e api.Image
	err := json.Unmarshal([]byte(k), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	_, err = c.Images().Put(context.Background(), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, &e); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *imagesCRUD) Delete(c *client.Client, k string) {
}

type blobsCRUD struct{}

func (h *blobsCRUD) List(c *client.Client) {
	es, err := c.Blobs().List(context.Background(), &api.Void{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, es); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *blobsCRUD) Put(c *client.Client, k, v string) {
	var e api.Blob
	err := json.Unmarshal([]byte(k), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	_, err = c.Blobs().Put(context.Background(), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, &e); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *blobsCRUD) Delete(c *client.Client, k string) {
}

type bundlesCRUD struct{}

func (h *bundlesCRUD) List(c *client.Client) {
	es, err := c.Bundles().List(context.Background(), &api.Void{})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, es); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *bundlesCRUD) Put(c *client.Client, k, v string) {
	var e api.Bundle
	err := json.Unmarshal([]byte(k), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	_, err = c.Bundles().Put(context.Background(), &e)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err := proto.MarshalText(os.Stderr, &e); err != nil {
		log.Fatalf("err: %v", err)
	}
}

func (h *bundlesCRUD) Delete(c *client.Client, k string) {
}
