package publisher_status_mc_detection

type RmqPublisherStatusMcDetectionFormatter struct {
	Tid           string `json:"tid"`
	TidID         *int   `json:"tid_id"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}

// data ditampilkan dari input, bukan dari entity yg sudah tercreate
func PublisherStatusMcDetectionFormat(rmqPublisherStatusMcDetection RmqPublisherStatusMcDetection) RmqPublisherStatusMcDetectionFormatter {
	var formatter RmqPublisherStatusMcDetectionFormatter

	formatter.Tid = rmqPublisherStatusMcDetection.Tid
	formatter.TidID = rmqPublisherStatusMcDetection.TidID
	formatter.DateTime = rmqPublisherStatusMcDetection.DateTime
	formatter.StatusSignal = rmqPublisherStatusMcDetection.StatusSignal
	formatter.StatusStorage = rmqPublisherStatusMcDetection.StatusStorage
	formatter.StatusRam = rmqPublisherStatusMcDetection.StatusRam
	formatter.StatusCpu = rmqPublisherStatusMcDetection.StatusCpu

	return formatter
}
