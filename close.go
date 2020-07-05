package gomongo

import (
	"context"
)

func (c *MongoDB) Close() error {
	return c.Client.Disconnect(context.TODO())
}
