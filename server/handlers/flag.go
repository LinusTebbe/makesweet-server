package handlers

import (
	"fmt"
	"makesweet/utils"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateFlagGif(ctx *gin.Context) {
	imageFilePath, err := utils.SaveImageFromContext(ctx, "image")
	if err != nil {
		if err.Error() == "Fail to save 'image' in the server" {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer os.Remove(imageFilePath)

	destFolderPath := os.Getenv("SAVE_IMAGE_FOLDER")
	outputID := uuid.New()
	outputFileName := fmt.Sprintf("%s.gif", outputID.String())
	outputPath := fmt.Sprintf("%s/%s", destFolderPath, outputFileName)

	flagCreateCommand := utils.NewCommandBuilder().Flag(imageFilePath, outputPath)
	err = flagCreateCommand.Run()
	if err != nil {
		log.Error("Flag gif make fail.", "err", err)
		ctx.JSON(http.StatusInternalServerError, "Fail to create gif")
		return
	}
	defer os.Remove(outputPath)

	_, err = os.Stat(outputPath)
	if err != nil {
		log.Error("Flag output check fail.", "err", err)
		ctx.JSON(http.StatusInternalServerError, "Fail to create gif")
		return
	}
	ctx.File(outputPath)
}
