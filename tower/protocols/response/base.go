package response

import (
	"encoding/json"
	"io"
)

type BaseResponse[T any] struct {
	ErrCode string `json:"errCode"` // 错误码
	ErrMsg  string `json:"errMsg"`  // 错误描述
	Data    T      `json:"data"`    // 返回数据
}

// NewResponse 创建泛型结果
func NewResponse[T any](errCode, errMsg string, data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}
}

// ResolveJsonResponse 从json解析结果
func ResolveJsonResponse[T any](bodyR io.Reader) (*BaseResponse[T], error) {
	resIns := new(BaseResponse[T])
	if bodyB, err := io.ReadAll(bodyR); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bodyB, resIns); err != nil {
		return nil, err
	} else {
		return resIns, nil
	}
}
