package routes

import (
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/connection"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/handler"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/middleware"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/model/publisher_status_mc_detection"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-publisherRmq-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherStatusMcDetectionRepository := publisher_status_mc_detection.NewRepository(connection.RabbitMQ())
	tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())

	// Initialize services
	publisherStatusMcDetectionService := publisher_status_mc_detection.NewService(publisherStatusMcDetectionRepository, tbTidRepository)

	// Initialize handlers
	publisherStatusMcDetectionHandler := handler.NewPublisherStatusMcDetectionHandler(publisherStatusMcDetectionService)

	// Configure routes
	api := router.Group("/publisher/atmvideopack/v1")

	statusMcDetectionRoutes := api.Group("/statusmcdetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureStatusMcDetectionRoutes(statusMcDetectionRoutes, publisherStatusMcDetectionHandler)

}

func configureStatusMcDetectionRoutes(api *gin.RouterGroup, handler *handler.PublisherStatusMcDetectionHandler) {
	api.POST("/create", handler.CreateQueueStatusMcDetection)
}
