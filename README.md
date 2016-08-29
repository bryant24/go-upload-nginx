**This package used for upload big file to nginx upload module extension.**

**You need install nginx upload module in sever env first.**

Usage:
```golang
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
```
	