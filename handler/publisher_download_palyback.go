package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/constant"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/model/publisher_download_playback"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PublisherDownloadPlaybackHandler struct {
	publisherDownloadPlaybackService publisher_download_playback.Service
}

func NewPublisherDownloadPlaybackHandler(publisherDownloadPlaybackService publisher_download_playback.Service) *PublisherDownloadPlaybackHandler {
	return &PublisherDownloadPlaybackHandler{publisherDownloadPlaybackService}
}

// public message to rmqg
func (h *PublisherDownloadPlaybackHandler) CreateQueueDownloadPlayback(c *gin.Context) {
	var input publisher_download_playback.RmqPublisherDownloadPlaybackInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	downloadPlayback, err := h.publisherDownloadPlaybackService.CreateQueueDownloadPlayback(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, publisher_download_playback.PublisherDownloadPlaybackFormat(downloadPlayback))
	log_function.Info("Queue download playback berhasil dibuat")
	c.JSON(http.StatusOK, response)
}
