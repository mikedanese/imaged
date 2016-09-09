package client

import (
	"net"
	"time"

	"github.com/mikedanese/imaged/api"
	"google.golang.org/grpc"
)

func New() (*Client, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(30 * time.Second),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", "foo.sock", timeout)
		}),
	}
	conn, err := grpc.Dial("unix://run/imaged/foo.sock", dialOpts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		images:  api.NewImagesClient(conn),
		bundles: api.NewBundlesClient(conn),
		blobs:   api.NewBlobsClient(conn),
	}, nil
}

type Client struct {
	images  api.ImagesClient
	bundles api.BundlesClient
	blobs   api.BlobsClient
}

func (c *Client) Images() api.ImagesClient   { return c.images }
func (c *Client) Bundles() api.BundlesClient { return c.bundles }
func (c *Client) Blobs() api.BlobsClient     { return c.blobs }
