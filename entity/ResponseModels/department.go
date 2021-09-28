package ResponseModels

// LoginMessage 账号密码信息
type LoginMessage struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
