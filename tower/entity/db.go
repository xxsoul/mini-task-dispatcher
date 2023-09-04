package entity

import "time"

// Tasks 任务列表
type Tasks struct {
	TaskId       int        // 任务自增id
	SeqNo        string     // 流水号
	CallbackUrl  string     // 回调地址
	Purpose      string     // 任务目的 train_model-训练模型 audio_inference-音频推理
	InputConfig  string     // 输入参数，json字符串
	OutputConfig string     // 输出参数，json字符串
	StepDetail   string     // 任务步进参数
	Status       string     // 任务状态
	CreateTime   *time.Time // 创建时间
	UpdateTime   *time.Time // 更新时间
}
