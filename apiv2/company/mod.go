package company

type RegisterAuthorityParams struct {
	SecretKey     string `json:"secret_key" binding:"required"`
	SecretCode    string `json:"secret_code" binding:"required"`
	TemporaryCode string `json:"temporaryCode" binding:"required"`
}
