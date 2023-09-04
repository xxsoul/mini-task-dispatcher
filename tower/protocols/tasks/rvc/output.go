package rvc

// TrainOutputConfig 模型训练输出相关的参数
type TrainOutputConfig struct {
	modelFileOssPath string // 模型文件在oss的上传路径
	indexFileOssPath string // 特征索引文件在oss的上传路径
}
