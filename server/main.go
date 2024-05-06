package main

import (
	"makesweet/handlers"
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	imageFolderPath := os.Getenv("SAVE_IMAGE_FOLDER")
	if len(strings.TrimSpace(imageFolderPath)) == 0 {
		log.Fatal("SAVE_IMAGE_FOLDER environment variable invalid or not set")
	}
	_, err := os.Stat(imageFolderPath)
	if err != nil {
		log.Info("Creating folder to save input and output images")
		err = os.MkdirAll(imageFolderPath, os.ModeAppend)
		if err != nil {
			log.Fatal("Failed to create image directory")
		}
	}

	r := gin.Default()

	gifGroup := r.Group("/gif")
	gifGroup.POST("/billboard", handlers.CreateBillboardGif)
	gifGroup.POST("/flag", handlers.CreateFlagGif)
	gifGroup.POST("/heart-locket", handlers.CreateHeartLocketGif)

	r.Run(":8080")
}
