package facerecognition

import (
	"fmt"

	"gocv.io/x/gocv"
)

// FaceRecognizer struct manages face detection and recognition
type FaceRecognizer struct {
	faceCascade    gocv.CascadeClassifier
	faceRecognizer gocv.FaceRecognizer
}

// NewFaceRecognizer initializes a FaceRecognizer with Haar cascade and model
func NewFaceRecognizer(cascadePath, modelPath string) (*FaceRecognizer, error) {
	faceCascade := gocv.NewCascadeClassifier()
	if !faceCascade.Load(cascadePath) {
		return nil, fmt.Errorf("Error loading cascade file: %v", cascadePath)
	}

	faceRecognizer := gocv.NewFaceRecognizer()
	if !faceRecognizer.Load(modelPath) {
		fmt.Println("Warning: No trained model found, using an empty recognizer.")
	}

	return &FaceRecognizer{faceCascade: faceCascade, faceRecognizer: faceRecognizer}, nil
}

// RecognizeFace detects and identifies a face in an image
func (fr *FaceRecognizer) RecognizeFace(imgPath string) (int, float64, error) {
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	if img.Empty() {
		return -1, 0, fmt.Errorf("Error reading image: %s", imgPath)
	}
	defer img.Close()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	faces := fr.faceCascade.DetectMultiScale(gray)
	if len(faces) == 0 {
		return -1, 0, fmt.Errorf("No face detected")
	}

	faceROI := gray.Region(faces[0])
	defer faceROI.Close()

	label, confidence := fr.faceRecognizer.Predict(faceROI)
	return label, confidence, nil
}

// RegisterFace adds a new face to the model
func (fr *FaceRecognizer) RegisterFace(imgPath string, label int) error {
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	if img.Empty() {
		return fmt.Errorf("Error reading image: %s", imgPath)
	}
	defer img.Close()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	faces := fr.faceCascade.DetectMultiScale(gray)
	if len(faces) == 0 {
		return fmt.Errorf("No face detected")
	}

	faceROI := gray.Region(faces[0])
	defer faceROI.Close()

	fr.faceRecognizer.Update([]gocv.Mat{faceROI}, []int{label})
	return nil
}
