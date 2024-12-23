package download

import (
	"time"

	"github.com/blacksheepaul/qiniu-oss-go-sdk/config"

	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
)

type Downloader struct {
	credentials *credentials.Credentials
	Domain      string
	Deadline    int64
}

func NewDownloader(cfg config.DownloadConfig) *Downloader {
	mac := credentials.NewCredentials(cfg.AccessKey, cfg.SecretKey)

	return &Downloader{
		credentials: mac,
		Domain:      cfg.Domain,
		Deadline:    cfg.Deadline,
	}
}

func (d *Downloader) GetPrivateAccessURL(key string) (url string) {
	ddl := d.Deadline + time.Now().Unix()
	return storage.MakePrivateURLv2(d.credentials, d.Domain, key, ddl)
}
