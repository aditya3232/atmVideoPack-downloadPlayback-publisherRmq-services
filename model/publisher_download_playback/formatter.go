package publisher_download_playback

type RmqPublisherDownloadPlaybackFormatter struct {
	Tid             string `json:"tid"`
	DateModified    string `json:"date_modified"`
	DurationMinutes string `json:"duration_minutes"`
	FileSizeBytes   string `json:"file_size_bytes"`
	Filename        string `json:"filename"`
	Url             string `json:"url"`
}

// data ditampilkan dari input, bukan dari entity yg sudah tercreate
func PublisherDownloadPlaybackFormat(rmqPublisherDownloadPlayback RmqPublisherDownloadPlayback) RmqPublisherDownloadPlaybackFormatter {
	var formatter RmqPublisherDownloadPlaybackFormatter

	formatter.Tid = rmqPublisherDownloadPlayback.Tid
	formatter.DateModified = rmqPublisherDownloadPlayback.DateModified
	formatter.DurationMinutes = rmqPublisherDownloadPlayback.DurationMinutes
	formatter.FileSizeBytes = rmqPublisherDownloadPlayback.FileSizeBytes
	formatter.Filename = rmqPublisherDownloadPlayback.Filename
	formatter.Url = rmqPublisherDownloadPlayback.Url

	return formatter
}
