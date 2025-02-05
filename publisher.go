// sg-pubsub/pkg/clientlib/pubsublib/publisher.go
package pubsublib

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)


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
