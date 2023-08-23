package resp

type UserResp struct {
	Name string `json:"name" example:"ray"`
	Age  int    `json:"age" example:"10"`
}

type CommandResp struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"error msg"`
	Data any    `json:"data"`
}
