package user

type LoginParams struct {
	Username     string `json:"username" form:"username" gorm:"commit:用户名"  binding:"required"`
	Password     string `json:"password" form:"password" gorm:"commit:密码"  binding:"required"`
	Redirect     string `json:"redirect_url" form:"redirect_url" gorm:"commit:重定向url"  binding:"required"`
	CompanyId    string `json:"companyId" form:"companyId" gorm:"commit:公司唯一标识，通过管理系统申请"  binding:"required"`
	ResponseType string `json:"response_type" form:"response_type" gorm:"commit:请求类型，比如登录"  binding:"required"`
}
