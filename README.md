# qiniu-oss-go-sdk

七牛云对象存储 Go 快速接入

### 传输过程

[参考](https://developer.qiniu.com/kodo/1205/programming-model)

### 文件流上传

1. 构造 `config.UploadConfig`

2. 调用 `NewUploader`，实例化一个 uploader

3. 调用 `StreamUpload` 进行上传

### 下载

公开空间

可直接下载，拼接 path+key 即为下载 URL

私有空间

1. 构造 `config.DownloadConfig`

2. 调用 `GetPrivateAccessURL` 获取下载 URL (注意有效期)
