package server

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
)

func WriteProto(tx *bolt.Tx, bucket, key string, m proto.Message) error {
	b := tx.Bucket([]byte(bucket))
	r, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	return b.Put([]byte(key), r)
}

func CreateBuckets(db *bolt.DB, names []string) error {
	return db.Update(func(tx *bolt.Tx) error {
		for _, name := range names {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})
}
