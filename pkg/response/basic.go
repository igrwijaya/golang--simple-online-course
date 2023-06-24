package response

type Basic struct {
	Code  uint
	Error error
	Data  interface{}
}

func Success() Basic {
	return Basic{
		Code: 200,
	}
}
