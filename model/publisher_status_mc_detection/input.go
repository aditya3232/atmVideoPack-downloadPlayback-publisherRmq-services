package publisher_status_mc_detection

type RmqPublisherStatusMcDetectionInput struct {
	Tid           string `form:"tid" binding:"required"`
	DateTime      string `form:"date_time" binding:"required"`
	StatusSignal  string `form:"status_signal" binding:"required"`
	StatusStorage string `form:"status_storage" binding:"required"`
	StatusRam     string `form:"status_ram" binding:"required"`
	StatusCpu     string `form:"status_cpu" binding:"required"`
	TidID         *int   `form:"tid_id"` // ini untuk menampung hasil pencarian id berdasarkan tid yang diinput di api
}
