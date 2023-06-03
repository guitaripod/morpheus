package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
)

func ProcessImage(imageFilePath string, outputPath string) (string, error) {
	file, err := os.Open(imageFilePath)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	grayImg := image.NewGray(img.Bounds())
	draw.Draw(grayImg, grayImg.Bounds(), img, image.ZP, draw.Src)

	outputFileName := strings.TrimSuffix(
		filepath.Base(imageFilePath),
		filepath.Ext(imageFilePath),
	) + ".png"

	outputFile, err := os.Create(filepath.Join(outputPath, outputFileName))
	if err != nil {
		return "", err
	}

	if strings.ToLower(filepath.Ext(outputFileName)) == ".png" {
		err = png.Encode(outputFile, grayImg)
	} else {
		err = jpeg.Encode(outputFile, grayImg, &jpeg.Options{Quality: 100})
	}

	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
