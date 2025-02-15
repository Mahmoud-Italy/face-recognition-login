package main

import (
	"fmt"
	"log"

	facerecognition "github.com/mahmoud-italy/face-recognition-login"
)

func main() {
	cascadePath := "assets/haarcascade_frontalface_default.xml"
	modelPath := "models/face_model.yml"

	faceRec, err := facerecognition.NewFaceRecognizer(cascadePath, modelPath)
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// Register a new face
	err = faceRec.RegisterFace("data/test_faces/face1.jpg", 1)
	if err != nil {
		log.Fatalf("Error registering face: %v", err)
	}
	fmt.Println("Face registered successfully.")

	// Recognize face
	label, confidence, err := faceRec.RecognizeFace("data/test_faces/face1.jpg")
	if err != nil {
		log.Fatalf("Error recognizing face: %v", err)
	}
	fmt.Printf("Recognized label: %d, Confidence: %f\n", label, confidence)
}
