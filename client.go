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

// Close closes the Pub/Sub client
func (p *PubSubClient) Close() error {
	return p.client.Close()
}
