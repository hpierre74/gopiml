package main

import (
	"image/color"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/machinebox/sdk-go/facebox"
	"gocv.io/x/gocv"
)

var (
	blue          = color.RGBA{0, 0, 255, 0}
	faceAlgorithm = "haarcascade_frontalface_default.xml"
	fbox          = facebox.New("http://localhost:8080")
	data          = "face.jpg"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r)
	w.Write([]byte("Gorilla!\n"))
}

func mainn() {
	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(faceAlgorithm)
	defer classifier.Close()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

	// detect faces
	// rects := classifier.DetectMultiScale(img)
	// for _, r := range rects {
	// 	// Save each found face into the file
	// 	imgFace := img.Region(r)
	// 	imgName := fmt.Sprintf("%d.jpg", time.Now().UnixNano())
	// 	gocv.IMWrite(imgName, imgFace)
	// 	buf, err := gocv.IMEncode(".jpg", imgFace)
	// 	imgFace.Close()
	// 	if err != nil {
	// 		log.Printf("unable to encode matrix: %v", err)
	// 		continue
	// 	}

	// 	faces, err := fbox.Check(bytes.NewReader(buf))
	// 	if err != nil {
	// 		log.Printf("unable to recognize face: %v", err)
	// 	}

	// 	var caption = "I don't know you"
	// 	if len(faces) > 0 {
	// 		caption = fmt.Sprintf("I know you %s", faces[0].Name)
	// 	}

	// 	// draw rectangle for the face
	// 	size := gocv.GetTextSize(caption, gocv.FontHersheyPlain, 3, 2)
	// 	pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
	// 	gocv.PutText(&img, caption, pt, gocv.FontHersheyPlain, 3, blue, 2)
	// 	gocv.Rectangle(&img, r, blue, 3)

	// }
}
