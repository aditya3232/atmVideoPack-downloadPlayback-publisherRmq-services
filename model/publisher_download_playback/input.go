package publisher_download_playback

type RmqPublisherDownloadPlaybackInput struct {
	Tid             string `form:"tid" binding:"required"`
	DateModified    string `form:"date_modified" binding:"required"`
	DurationMinutes string `form:"duration_minutes" binding:"required"`
	FileSizeBytes   string `form:"file_size_bytes" binding:"required"`
	Filename        string `form:"filename" binding:"required"`
	Url             string `form:"url" binding:"required"`
}
