package go_multiple_upload

import (
	"testing"
	"github.com/satori/go.uuid"
)

func TestUploader_SplitAndUpload(t *testing.T) {
	var uploader Uploader
	uploader = Uploader{
		//upload api
		UploadApi:"http://localhost:8000/v1.0/upload",
		//random uuid
		SessionId:uuid.NewV4().String(),
		//target file name
		TargetFile:"c:\\test.zip",
	}
	uploader.SplitAndUpload()
}
