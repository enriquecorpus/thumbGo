package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)
const (
	CMD = "ffmpeg"
)
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

func mimeFromIncipit(incipit []byte) string {
	incipitStr := []byte(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(string(incipitStr), magic) {
			fmt.Println(mime)
			//return mime
		}
	}
	return ""
}
func main() {
	ffmpegCmd := exec.Command(CMD, "-i", "/Users/enrique/Downloads/ffmpeg/2.mov", "-vf", "fps=1/1",
	"-s", "150x100", "-f", "image2pipe", "-")
	stdOut,_:= ffmpegCmd.Output()
	chunk:=split(stdOut,47993)
	for i := 0; i < len(chunk); i++ {
		t:=time.Now().String()
		err := ioutil.WriteFile( t + "file.jpg", chunk[i], 0644)
		if err != nil {
			// handle error
		}
	}
}
func  GenerateThumbnailFromVideo(fileBytes []byte) ([]byte, error) {
	cmd := exec.Command(CMD, "-ss", "0:00:05.000", "-i", "-", "-vframes", "1", "-q:v", "2", "-f", "singlejpeg", "-")
	cmd.Stdin = bytes.NewBuffer(fileBytes)
	return cmd.Output()
}

func GenerateThumbnailFromPhoto(fileBytes []byte) ([]byte, error) {
	cmd := exec.Command(CMD, "-i", "-", "-vf", "scale='150:100'", "-f", "singlejpeg", "-")
	cmd.Stdin = bytes.NewBuffer(fileBytes)
	return cmd.Output()
}