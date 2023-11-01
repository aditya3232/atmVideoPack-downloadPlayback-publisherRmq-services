package publisher_download_playback

type Service interface {
	CreateQueueDownloadPlayback(input RmqPublisherDownloadPlaybackInput) (RmqPublisherDownloadPlayback, error)
}

type service struct {
	downloadPlaybackRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) CreateQueueDownloadPlayback(input RmqPublisherDownloadPlaybackInput) (RmqPublisherDownloadPlayback, error) {

	newRmqPublisherDownloadPlayback := RmqPublisherDownloadPlayback(input)

	newDownloadPlayback, err := s.downloadPlaybackRepository.CreateQueueDownloadPlayback(newRmqPublisherDownloadPlayback)
	if err != nil {
		return newDownloadPlayback, err
	}

	return newDownloadPlayback, nil
}
