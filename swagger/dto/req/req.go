package req

type UserListReq struct {
	Size int `json:"size"`
	Page int `json:"page"`
}
