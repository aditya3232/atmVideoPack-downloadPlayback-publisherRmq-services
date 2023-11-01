package publisher_status_mc_detection

import "github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/model/tb_tid"

type Service interface {
	CreateQueueStatusMcDetection(input RmqPublisherStatusMcDetectionInput) (RmqPublisherStatusMcDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
	tbTidRepository             tb_tid.Repository
}

func NewService(repository Repository, tbTidRepository tb_tid.Repository) *service {
	return &service{repository, tbTidRepository}
}

// public message to rmq
func (s *service) CreateQueueStatusMcDetection(input RmqPublisherStatusMcDetectionInput) (RmqPublisherStatusMcDetection, error) {
	var rmqPublisherStatusMcDetection RmqPublisherStatusMcDetection

	// get id from input tid
	tidID, err := s.tbTidRepository.GetOneByTid(input.Tid)
	if err != nil {
		return rmqPublisherStatusMcDetection, err
	}

	newRmqPublisherStatusMcDetection := RmqPublisherStatusMcDetection{
		TidID:         &tidID.ID,
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		StatusSignal:  input.StatusSignal,
		StatusStorage: input.StatusStorage,
		StatusRam:     input.StatusRam,
		StatusCpu:     input.StatusCpu,
	}

	newStatusMcDetection, err := s.statusMcDetectionRepository.CreateQueueStatusMcDetection(newRmqPublisherStatusMcDetection)
	if err != nil {
		return newStatusMcDetection, err
	}

	return newStatusMcDetection, nil
}
