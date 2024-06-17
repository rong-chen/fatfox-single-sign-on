package email

import (
	"sync"
	"time"
)

type Params struct {
	Email string `json:"email" gorm:"not null" form:"email"`
}

type Mailer struct {
	Body
}
type Body struct {
	Account  string
	Password string
	F, T, C  string
	HtmlBody string
}

type MailLimiter struct {
	lock     sync.Mutex
	lastSent map[string]time.Time
}

var HtmlBody = `
	 <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f0f0f0;
                }
                .container {
                    max-width: 600px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #ffffff;
                    border-radius: 10px;
                    box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.1);
                }
                h1 {
                    color: #333333;
                }
                p {
                    color: #666666;
                }
				.button {
					display: inline-block;
					padding: 5px 10px;
					background-color: #007bff;
					color: #ffffff;
					text-decoration: none;
					border-radius: 5px;
				}
            </style>
        </head>
        <body>
            <div class="container">
                <h1>用户注册提醒</h1>
                <p>您的邮箱有新用户注册了，请及时处理。</p>
               	<h1>验证码：%s</h1>
            </div>
        </body>
        </html>
`
