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

func CreateHeartLocketGif(ctx *gin.Context) {
	leftImageFilePath, err := utils.SaveImageFromContext(ctx, "image-left")
	if err != nil {
		if err.Error() == "Fail to save 'image-left' in the server" {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer os.Remove(leftImageFilePath)

	rightImageFilePath, err := utils.SaveImageFromContext(ctx, "image-right")
	if err != nil {
		if err.Error() == "Fail to save 'image-right' in the server" {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer os.Remove(rightImageFilePath)

	destFolderPath := os.Getenv("SAVE_IMAGE_FOLDER")
	outputID := uuid.New()
	outputFileName := fmt.Sprintf("%s.gif", outputID.String())
	outputPath := fmt.Sprintf("%s/%s", destFolderPath, outputFileName)

	heartLocketCreateCommand := utils.NewCommandBuilder().HeartLocket(leftImageFilePath, rightImageFilePath, outputPath)
	err = heartLocketCreateCommand.Run()
	if err != nil {
		log.Error("Heart locket gif make fail.", "err", err)
		ctx.JSON(http.StatusInternalServerError, "Fail to create gif")
		return
	}
	defer os.Remove(outputPath)

	_, err = os.Stat(outputPath)
	if err != nil {
		log.Error("Heart locket output check fail.", "err", err)
		ctx.JSON(http.StatusInternalServerError, "Fail to create gif")
		return
	}
	ctx.File(outputPath)
}
