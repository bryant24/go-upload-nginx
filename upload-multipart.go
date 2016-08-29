package go_multiple_upload

import (
	"os"
	"fmt"
	"net/http"
	"bytes"
	"strconv"
	"math"
	"io/ioutil"
)

type Uploader struct {
	UploadApi  string
	SessionId  string
	TargetFile string
}

const ChunkSize = 2//CHUNK SIZE

func (this *Uploader)SplitAndUpload() {
	file, err := os.Open(this.TargetFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var fileSize int64 = fileInfo.Size()
	const fileChunk = ChunkSize * (1 << 20)
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Split to %d pieces.\n", totalPartsNum)
	for i := uint64(0); i < totalPartsNum; i++ {
		partSize := int(math.Min(fileChunk, float64(fileSize - int64(i * fileChunk))))
		partBuffer := make([]byte, partSize)
		startPos := int64(i * fileChunk)
		stopPos := startPos + int64(partSize) - 1

		file.Read(partBuffer)
		//here you can use more go routine to speed up,like multiple thread upload
		this.HttpDo(partBuffer, startPos, stopPos, int64(partSize), fileSize, fileInfo.Name())
		fmt.Printf("Start:%d,Stop:%d,ChunkSize:%d\n", startPos, stopPos, partSize)
	}

}
func (this *Uploader)HttpDo(buf []byte, start, stop, chunkSize, total int64, filename string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", this.UploadApi, bytes.NewReader(buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Session-ID", this.SessionId)
	req.Header.Set("X-Content-Range", "bytes " + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(stop, 10) + "/" + strconv.FormatInt(total, 10))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Content-Disposition", "attachment;filename=\"" + filename + "\"")
	req.Header.Set("Content-Length", strconv.FormatInt(chunkSize, 10))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body:", string(b))
}