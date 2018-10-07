package taskrunner

import (
	"image"
    "image/draw"
    "image/jpeg"
    "os"
    // "image/color"
    // "io/ioutil"
    // "fmt"
    "bufio"
    "bytes"
    "log"
	"os/exec"
	"github.com/space3/autonomous-client/jobmodels"
)

func CreateEncodingPickle(j JobSpec) {
	face_detector := exec.Command("python3", "face_detector.py" "-p", j.targetEmbedding)
	face_detector.Start()
}

func MatchFace(imgBytes []byte) bool {
	img, _, _ := image.Decode(bytes.NewReader(imgBytes))
	out, _ os.Create("./img002.jpg")
	
	var imageBuf bytes.Buffer
	err - jpeg.Encode(&imageBuf, img, nil)
	fw := bufio.NewWriter(out)
	fw.Write(imagebuf.Bytes())

	face_detector := exec.Command("python3", "face_detector.py", "-i", "img002.jpg", "-r", "encoding.pkl")
	output, _ := face_detector.StdoutPipe()

	face_detector.Start()

	if output == "True" {
		return true
	}
	return false
}