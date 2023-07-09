package response

type Basic struct {
	Code  uint        `json:"code,omitempty"`
	Error error       `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func Success() Basic {
	return Basic{
		Code: 200,
	}
}
