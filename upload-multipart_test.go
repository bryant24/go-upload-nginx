package go_multiple_upload

import (
	"testing"
)

func TestUploader_SplitAndUpload(t *testing.T) {
	var uploader Uploader
	uploader = Uploader{
		//upload api
		UploadApi:"http://localhost:8000/v1.0/upload",
		//target file name
		TargetFile:"c:\\test.zip",
		//chunksize:2MB
		ChunkSize:2,
	}
	uploader.SplitAndUpload()
}
