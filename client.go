// sg-pubsub/pkg/clientlib/pubsublib/client.go
package pubsublib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type PubSubClient struct {
	client *pubsub.Client
}

// NewPubSubClient initializes a new Pub/Sub client
func NewPubSubClient(ctx context.Context, projectID string) (*PubSubClient, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}
	return &PubSubClient{client: client}, nil
}

// PublishMessage publishes a message to a given topic
func (p *PubSubClient) PublishMessage(ctx context.Context, topicID string, message interface{}) error {
	topic := p.client.Topic(topicID)
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	result := topic.Publish(ctx, &pubsub.Message{Data: data})
	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Printf("Message published to topic %s", topicID)
	return nil
}

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

// Close closes the Pub/Sub client
func (p *PubSubClient) Close() error {
	return p.client.Close()
}
