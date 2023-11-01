package publisher_download_playback

import (
	"context"
	"encoding/json"

	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	CreateQueueDownloadPlayback(rmqPublisherDownloadPlayback RmqPublisherDownloadPlayback) (RmqPublisherDownloadPlayback, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueDownloadPlayback(rmqPublisherDownloadPlayback RmqPublisherDownloadPlayback) (RmqPublisherDownloadPlayback, error) {

	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		return rmqPublisherDownloadPlayback, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"DownloadPlaybackQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return rmqPublisherDownloadPlayback, err
	}

	// yang akan dimarshal
	var inputReadytoMarshal = RmqPublisherDownloadPlayback{
		Tid:             rmqPublisherDownloadPlayback.Tid,
		DateModified:    rmqPublisherDownloadPlayback.DateModified,
		DurationMinutes: rmqPublisherDownloadPlayback.DurationMinutes,
		FileSizeBytes:   rmqPublisherDownloadPlayback.FileSizeBytes,
		Filename:        rmqPublisherDownloadPlayback.Filename,
		Url:             rmqPublisherDownloadPlayback.Url,
	}

	// Convert the StatusMc struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return rmqPublisherDownloadPlayback, err
	}

	ctx := context.Background() // Create a context
	err = ch.PublishWithContext(ctx,
		"",
		"DownloadPlaybackQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return rmqPublisherDownloadPlayback, err
	}

	// Return the sent DownloadPlayback struct
	return rmqPublisherDownloadPlayback, err
}
