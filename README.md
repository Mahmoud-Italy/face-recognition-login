# Face Recognition Login in Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mahmoud-italy/face-recognition-login.svg)](https://pkg.go.dev/github.com/mahmoud-italy/face-recognition-login)
![License](https://img.shields.io/github/license/mahmoud-italy/face-recognition-login)
![Go Report Card](https://goreportcard.com/badge/github.com/mahmoud-italy/face-recognition-login)
![Version](https://img.shields.io/github/tag/mahmoud-italy/face-recognition-login)
![Go Version](https://img.shields.io/badge/go-1.23.6-blue)
  
This is a face recognition login package using Go and OpenCV (`gocv`).
  
## 🚀 Features
- Face detection using Haar cascades
- Face recognition using OpenCV's FaceRecognizer
- Model training and updating
- Easy API for face registration and recognition

## 📦 Installation

```bash
go get -u github.com/mahmoud-italy/face-recognition-login
```

## 🏗️ Setup
Install pkg-config and OpenCV <br />
<i>Depending on your OS using the appropriate command:</i>

Ubuntu/Debian
```bash
sudo apt update
sudo apt install -y pkg-config
sudo apt install -y libopencv-dev
```

Mac (Homebrew)
```bash
brew install pkg-config
brew install opencv
```

Windows (Using MSYS2) If you're on Windows, install MSYS2 and then run
```bash
pacman -S mingw-w64-x86_64-pkg-config
```
Download and install OpenCV from the official website: <a href="https://opencv.org/releases/" target="_blank">https://opencv.org/releases</a>
Set the OpenCV_DIR environment variable to the installation path.

Verify Installation
Run the following command to check if pkg-config is properly installed:
```bash
pkg-config --version
```

If it prints a version number, pkg-config is working.
Also, check if OpenCV is found by running:
```bash
pkg-config --modversion opencv4
```
If it returns an OpenCV version number (e.g., 4.5.2), everything is good. <br />

## 🔧 Example Usage
```bash
import (
    "fmt"
	facerecognition "github.com/mahmoud-italy/face-recognition-login"
)

func main() {
    faceRec, _ := facerecognition.NewFaceRecognizer("assets/haarcascade_frontalface_default.xml", "models/face_model.yml")
    faceRec.RegisterFace("data/test_faces/face1.jpg", 1)
    label, confidence, _ := faceRec.RecognizeFace("data/test_faces/face1.jpg")
    fmt.Printf("Recognized label: %d, Confidence: %f\n", label, confidence)
}
```

## 🛠️ Running Tests
```bash
go test
```

## 🔖 License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

## 🤝 Contributing
Feel free to contribute by submitting issues or pull requests.
