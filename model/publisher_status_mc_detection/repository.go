package publisher_status_mc_detection

import (
	"context"
	"encoding/json"

	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	CreateQueueStatusMcDetection(rmqPublisherStatusMcDetection RmqPublisherStatusMcDetection) (RmqPublisherStatusMcDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueStatusMcDetection(rmqPublisherStatusMcDetection RmqPublisherStatusMcDetection) (RmqPublisherStatusMcDetection, error) {

	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		return rmqPublisherStatusMcDetection, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"StatusMcDetectionQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return rmqPublisherStatusMcDetection, err
	}

	// yang akan dimarshal
	var inputReadytoMarshal = RmqPublisherStatusMcDetection{
		TidID:         rmqPublisherStatusMcDetection.TidID,
		Tid:           rmqPublisherStatusMcDetection.Tid,
		DateTime:      rmqPublisherStatusMcDetection.DateTime,
		StatusSignal:  rmqPublisherStatusMcDetection.StatusSignal,
		StatusStorage: rmqPublisherStatusMcDetection.StatusStorage,
		StatusRam:     rmqPublisherStatusMcDetection.StatusRam,
		StatusCpu:     rmqPublisherStatusMcDetection.StatusCpu,
	}

	// Convert the StatusMc struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return rmqPublisherStatusMcDetection, err
	}

	ctx := context.Background() // Create a context
	err = ch.PublishWithContext(ctx,
		"",
		"StatusMcDetectionQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return rmqPublisherStatusMcDetection, err
	}

	// Return the sent StatusMcDetection struct
	return rmqPublisherStatusMcDetection, err
}
