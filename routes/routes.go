package routes

import (
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/connection"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/handler"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/middleware"
	"github.com/aditya3232/atmVideoPack-downloadPlayback-publisherRmq-services.git/model/publisher_download_playback"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherDownloadPlaybackRepository := publisher_download_playback.NewRepository(connection.RabbitMQ())

	// Initialize services
	publisherDownloadPlaybackService := publisher_download_playback.NewService(publisherDownloadPlaybackRepository)

	// Initialize handlers
	publisherDownloadPlaybackHandler := handler.NewPublisherDownloadPlaybackHandler(publisherDownloadPlaybackService)

	// Configure routes
	api := router.Group("/publisher/atmvideopack/v1")

	DownloadPlaybackRoutes := api.Group("/downloadplayback", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureDownloadPlaybackRoutes(DownloadPlaybackRoutes, publisherDownloadPlaybackHandler)

}

func configureDownloadPlaybackRoutes(api *gin.RouterGroup, handler *handler.PublisherDownloadPlaybackHandler) {
	api.POST("/create", handler.CreateQueueDownloadPlayback)
}
