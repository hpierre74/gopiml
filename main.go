package main

import (
	"fmt"
	"log"
	"time"

	"gocv.io/x/gocv"
)

var (
	faceAlgorithm = "haarcascade_frontalface_default.xml"
)

// renamed cause taken from camera api to rework
func detectGocvFaces() {
	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(faceAlgorithm)
	defer classifier.Close()

	// detect faces
	rects := classifier.DetectMultiScale(img)
	for _, r := range rects {
		// Save each found face into the file
		imgFace := img.Region(r)
		imgName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
		gocv.IMWrite(imgName, imgFace)
		buf, err := gocv.IMEncode(".jpg", imgFace)
		imgFace.Close()
		if err != nil {
			log.Printf("unable to encode matrix: %v", err)
			continue
		}

		log.Print(buf)
	}
}
