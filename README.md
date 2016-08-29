#go-upload-nginx

**This package used for upload big file to nginx upload module extension.**

**You need install nginx upload module in sever env first.**


##Install
Install the package with:
```bash
go get https://github.com/bryant24/go-upload-nginx
```

Import it with:
```go
import "github.com/bryant24/go-upload-nginx"
```

##Usage:
```go
var uploader Uploader
uploader = Uploader{
	//upload api
	UploadApi:"http://localhost:8000/v1.0/upload",
	//target file name
	TargetFile:"c:\\test.zip",
	//chunk size:2MB
	ChunkSize:2,
}
uploader.SplitAndUpload()
```
	