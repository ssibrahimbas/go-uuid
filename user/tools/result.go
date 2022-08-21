package tools

type Result struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResult struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Result) Error() string {
	return r.Message
}

func (r *DataResult) Error() string {
	return r.Message
}

func SuccessResult(msg string) *Result {
	return &Result{
		Success: true,
		Message: msg,
	}
}

func ErrorResult(msg string) *Result {
	return &Result{
		Success: false,
		Message: msg,
	}
}

func SuccessDataResult(msg string, data interface{}) *DataResult {
	return &DataResult{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func ErrorDataResult(msg string, data interface{}) *DataResult {
	return &DataResult{
		Success: false,
		Message: msg,
		Data:    data,
	}
}
