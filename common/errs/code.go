package errs

// ReturnFormat 返回格式
type ReturnFormat struct {
	Code interface{} `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	// Success 成功;
	Success = iota

	// RequestParamIsNull 请求参数为空;
	RequestParamIsNull
)

const (
	// XXX example srv error code const;
	XXX = iota + 1000
)
