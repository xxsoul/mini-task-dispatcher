package tasks

import "encoding/json"

// Step 任务执行步骤
type Step struct {
	Id   int               `json:"id"`   // 步骤id，排序依据，升序排列
	Type string            `json:"type"` // 类型os_cmd-执行系统命令 method-执行内置指令
	Env  map[string]string `json:"env"`  // 环境变量，如果为空则使用上一步的环境变量，如果是第一步则使用任务的环境变量
	Cmd  string            `json:"cmd"`  // 准备执行的命令，cmd中需包含完整的参数列表
	Pwd  *string           `json:"pwd"`  // 执行命令的工作目录
}

// Task 任务信息
type Task struct {
	SeqNo        string            `json:"seqNo"`        // 任务流水号
	Env          map[string]string `json:"env"`          // 执行任务所需的环境变量
	InputConfig  json.RawMessage   `json:"inputConfig"`  // 输入参数
	OutputConfig json.RawMessage   `json:"outputConfig"` // 暑促参数
	Steps        []Step            `json:"steps"`        // 任务执行步骤
}
