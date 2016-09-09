// auto-generated
// do not modify this file by hand

package server

import (
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"github.com/mikedanese/imaged/api"
	"golang.org/x/net/context"
)

var (
	imagesBucket = "images"

	bundlesBucket = "bundles"

	blobsBucket = "blobs"

	resourceBuckets = []string{

		imagesBucket,

		bundlesBucket,

		blobsBucket,
	}
)

func (s *imagesServerImpl) List(context.Context, *api.Void) (*api.ImageList, error) {
	var es api.ImageList
	return &es, s.s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(imagesBucket)).ForEach(func(k, v []byte) error {
			var e api.Image
			if err := proto.Unmarshal(v, &e); err != nil {
				return err
			}
			es.Images = append(es.Images, &e)
			return nil
		})
	})
}

func (s *imagesServerImpl) Put(c context.Context, req *api.Image) (*api.Image, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return WriteProto(tx, imagesBucket, req.Meta.Name, req)
	})
}

func (s *imagesServerImpl) Delete(c context.Context, req *api.Image) (*api.Image, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(imagesBucket)).Delete([]byte(req.Meta.Name))
	})
}

func (s *bundlesServerImpl) List(context.Context, *api.Void) (*api.BundleList, error) {
	var es api.BundleList
	return &es, s.s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bundlesBucket)).ForEach(func(k, v []byte) error {
			var e api.Bundle
			if err := proto.Unmarshal(v, &e); err != nil {
				return err
			}
			es.Bundles = append(es.Bundles, &e)
			return nil
		})
	})
}

func (s *bundlesServerImpl) Put(c context.Context, req *api.Bundle) (*api.Bundle, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return WriteProto(tx, bundlesBucket, req.Meta.Name, req)
	})
}

func (s *bundlesServerImpl) Delete(c context.Context, req *api.Bundle) (*api.Bundle, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bundlesBucket)).Delete([]byte(req.Meta.Name))
	})
}

func (s *blobsServerImpl) List(context.Context, *api.Void) (*api.BlobList, error) {
	var es api.BlobList
	return &es, s.s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(blobsBucket)).ForEach(func(k, v []byte) error {
			var e api.Blob
			if err := proto.Unmarshal(v, &e); err != nil {
				return err
			}
			es.Blobs = append(es.Blobs, &e)
			return nil
		})
	})
}

func (s *blobsServerImpl) Put(c context.Context, req *api.Blob) (*api.Blob, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return WriteProto(tx, blobsBucket, req.Meta.Name, req)
	})
}

func (s *blobsServerImpl) Delete(c context.Context, req *api.Blob) (*api.Blob, error) {
	return req, s.s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(blobsBucket)).Delete([]byte(req.Meta.Name))
	})
}
