package upload

import (
	"context"
	"io"

	"github.com/blacksheepaul/qiniu-oss-go-sdk/config"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/region"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
)

type Uploader struct {
	credentials *auth.Credentials
	options     uploader.UploadManagerOptions
	Bucket      string
	AccessUrl   string
	Regions     string
}

func NewUploader(cfg config.UploadConfig) *Uploader {
	mac := credentials.NewCredentials(cfg.AccessKey, cfg.SecretKey)

	var cregion string
	if cfg.Region == "" {
		cregion = "z2" // 默认华南
	} else {
		cregion = cfg.Region
	}

	options := uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
			Regions:     region.GetRegionByID(cregion, true),
		},
	}

	return &Uploader{
		credentials: mac,
		options:     options,
		Bucket:      cfg.Bucket,
		AccessUrl:   cfg.AccessUrl,
		Regions:     cfg.Region,
	}
}

func (u *Uploader) FileUpload(filepath string) {
	// do something
}

func (u *Uploader) StreamUpload(key string, reader io.Reader) error {
	uploadManager := uploader.NewUploadManager(&u.options)
	err := uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: u.Bucket,
		ObjectName: &key,
	}, nil)
	return err
}

func (u *Uploader) DirectoryUpload() {
	// do something
}

func (u *Uploader) FileSplitUpload() {
	// do something
}

// Not in plan
// func ResumableUpload () {}
