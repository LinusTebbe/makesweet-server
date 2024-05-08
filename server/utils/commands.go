package utils

import (
	"os/exec"
)

type commandBuilder struct{}
type CommandBuilder interface {
	Billboard(imagePath string, outputPath string) *exec.Cmd
	Flag(imagePath string, outputPath string) *exec.Cmd
	HeartLocket(leftImagePath string, rightImagePath string, outputPath string) *exec.Cmd
	Circuit(imagePath string, outputPath string) *exec.Cmd
	Bear(imagePath string, outputPath string) *exec.Cmd
	Doll(imageLeftPath string, imageMidPath string, imageRightPath string, outputPath string) *exec.Cmd
}

func NewCommandBuilder() CommandBuilder {
	return &commandBuilder{}
}
func (*commandBuilder) Billboard(imagePath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/billboard-cityscape.zip",
		"--in",
		imagePath,
		"--gif",
		outputPath,
	)

	return cmd
}

func (*commandBuilder) Flag(imagePath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/flag.zip",
		"--in",
		imagePath,
		"--gif",
		outputPath,
	)

	return cmd
}

func (*commandBuilder) HeartLocket(leftImagePath string, rightImagePath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/heart-locket.zip",
		"--in",
		leftImagePath,
		rightImagePath,
		"--gif",
		outputPath,
	)

	return cmd
}

func (*commandBuilder) Circuit(imagePath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/circuit-board.zip",
		"--in",
		imagePath,
		"--gif",
		outputPath,
	)

	return cmd
}

func (*commandBuilder) Bear(imagePath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/flying-bear.zip",
		"--in",
		imagePath,
		"--gif",
		outputPath,
	)

	return cmd
}

func (*commandBuilder) Doll(imageLeftPath string, imageMidPath string, imageRightPath string, outputPath string) *exec.Cmd {
	cmd := exec.Command(
		"/reanimator",
		"--zip",
		"/makesweet/templates/nesting-doll.zip",
		"--in",
		imageLeftPath,
		imageMidPath,
		imageRightPath,
		"--gif",
		outputPath,
	)

	return cmd
}
