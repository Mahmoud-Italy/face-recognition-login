package facerecognition

import (
	"testing"
)

func TestNewFaceRecognizer(t *testing.T) {
	_, err := NewFaceRecognizer("./assets/haarcascade_frontalface_default.xml", "./models/face_model.yml")
	if err != nil {
		t.Errorf("Failed to initialize FaceRecognizer: %v", err)
	}
}

func TestRecognizeFace_NoFace(t *testing.T) {
	faceRec, _ := NewFaceRecognizer("./assets/haarcascade_frontalface_default.xml", "./models/face_model.yml")
	_, _, err := faceRec.RecognizeFace("./data/test_faces/no_face.jpg")
	if err == nil {
		t.Errorf("Expected an error for no face detected, but got none")
	}
}

func TestRegisterFace(t *testing.T) {
	faceRec, _ := NewFaceRecognizer(".assets/haarcascade_frontalface_default.xml", "./models/face_model.yml")
	err := faceRec.RegisterFace("./data/test_faces/face1.jpg", 1)
	if err != nil {
		t.Errorf("Failed to register face: %v", err)
	}
}
