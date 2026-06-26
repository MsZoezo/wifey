package backends

import (
	"context"
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/queues"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

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

func (c *CloudflareBackend) Send(message Message) error {
	_, err := c.client.Queues.Messages.Push(
		context.TODO(),
		c.queueId,
		queues.MessagePushParams{
			AccountID: cloudflare.F(c.accountId),
			Body: queues.MessagePushParamsBody{
				Body:        cloudflare.F[interface{}](message),
				ContentType: cloudflare.F(queues.MessagePushParamsBodyContentTypeJson),
			},
		},
	)

	return err
}

func (c *CloudflareBackend) Receive() (message *Message, err error) {
	response, err := c.client.Queues.Messages.Pull(
		context.TODO(),
		c.queueId,
		queues.MessagePullParams{
			AccountID: cloudflare.F(c.accountId),
		},
	)

	if err != nil {
		return nil, err
	}

	if len(response.Messages) == 0 {
		return nil, nil
	}

	queueMessage := response.Messages[0]

	m := new(Message)

	jsonErr := json.Unmarshal([]byte(queueMessage.Body), m)

	if jsonErr != nil {
		return nil, err
	}

	_, ackErr := c.client.Queues.Messages.Ack(
		context.TODO(),
		c.queueId,
		queues.MessageAckParams{
			AccountID: cloudflare.F(c.accountId),
			Acks: cloudflare.F([]queues.MessageAckParamsAck{
				{
					LeaseID: cloudflare.F(queueMessage.LeaseID),
				},
			}),
		},
	)

	if ackErr != nil {
		return nil, ackErr
	}

	return m, err
}
