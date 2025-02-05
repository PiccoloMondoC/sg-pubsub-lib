// sg-pubsub/pkg/clientlib/pubsublib/subscriber.go
package pubsublib

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

// Subscribe starts listening for messages on a subscription
func (p *PubSubClient) Subscribe(ctx context.Context, subscriptionID string, handler func(context.Context, *pubsub.Message)) error {
	sub := p.client.Subscription(subscriptionID)
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handler(ctx, msg)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("failed to receive messages: %w", err)
	}
	return nil
}
