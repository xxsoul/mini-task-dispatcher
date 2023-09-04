package rvc

// TrainInputConfig 应用于rvc模型训练的输入参数
type TrainInputConfig struct {
	// 模型所需的参数
	TaskName       string  `json:"taskName"`       // 任务名 // 任务名
	SamplingRate   *int    `json:"samplingRate"`   // 采样率 // 采样率
	F0Guide        *bool   `json:"f0Guide"`        // 声高指导 // 声高指导
	F0Method       *string `json:"f0Method"`       // 声高模型 // 声高模型
	TotalEpoch     *int    `json:"totalEpoch"`     // 总迭代次数 // 总迭代次数
	OnlyLatestCkpt *bool   `json:"onlyLatestCkpt"` // 是否仅保存最后一 // 是否仅保存最后一个模型

	// pilot所需的参数
	OssTrainSetUrl string `json:"ossTrainSetUrl"` // 训练集在oss上的路径
}
