package main

import (
	"fmt"
	"log"
	"net"

	"github.com/boltdb/bolt"
	"github.com/mikedanese/imaged/api"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for _, name := range []string{"images", "bundles", "operations"} {
		if err := CreateBucket(db, name); err != nil {
			log.Fatal(err)
		}
	}

	s := grpc.NewServer()
	ds, err := net.Listen("unix", "/tmp/imaged.sock")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	api.RegisterImagesServer(s, &ImagesServerImpl{})
	s.Serve(ds)
}

func CreateBucket(db *bolt.DB, name string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

type ImagesServerImpl struct {
	db *bolt.DB
}

func (i *ImagesServerImpl) List(context.Context, *api.Void) (*api.ImageList, error) {
	return nil, nil
}

func (i *ImagesServerImpl) Create(c context.Context, ir *api.Image) (*api.Image, error) {
	return nil, i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("images"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
}

func (i *ImagesServerImpl) Delete(c context.Context, ir *api.Image) (*api.Image, error) {
	return nil, i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("images"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})
}
