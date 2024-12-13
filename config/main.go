package config

type UploadConfig struct {
	AccessKey string
	SecretKey string
	Bucket    string // 对应七牛云的"空间"
	AccessUrl string // 对应七牛云空间设置的外链域名
	Region    string
}

const (
	RegionsHuaDong          = "z0"
	RegionsHuaDongZheJiang2 = "cn-east-2"
	RegionsHuaBei           = "z1"
	RegionsHuaNan           = "z2"
	RegionsBeiMei           = "na0"
	RegionsXinJiaPo         = "as0"
	RegionsYaTaiHeNei       = "ap-southeast-2"
	RegionsYaTaiHuZhiMing   = "ap-southeast-3"
)

type DownloadConfig struct {
	AccessKey string
	SecretKey string
	Domain    string // 和上传的 AccessUrl 对应
	Deadline  int64
}
