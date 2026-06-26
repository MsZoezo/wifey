package backends

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/queues"
)

type CloudflareBackend struct {
	client cloudflare.Client

	queueId   string
	accountId string
}

func New(token string, queueId string, accountId string) *CloudflareBackend {
	c := new(CloudflareBackend)

	c.client = *cloudflare.NewClient(
		option.WithAPIToken(token),
	)

	c.queueId = queueId
	c.accountId = accountId

	return c
}

func (c *CloudflareBackend) Send(message string) error {
	_, err := c.client.Queues.Messages.Push(
		context.TODO(),
		c.queueId,
		queues.MessagePushParams{
			AccountID: cloudflare.F(c.accountId),
			Body: queues.MessagePushParamsBody{
				Body:        cloudflare.F[interface{}](message),
				ContentType: cloudflare.F(queues.MessagePushParamsBodyContentTypeText),
			},
		},
	)

	return err
}
