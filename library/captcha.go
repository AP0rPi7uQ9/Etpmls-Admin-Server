package library

import Package_Captcha "github.com/dchest/captcha"

type captcha struct {

}

func NewCaptcha() *captcha {
	return &captcha{}
}

// Verify whether the id and content of the verification code are associated
// 验证验证码的id及内容是否关联
func (this *captcha) Verify(id string, captcha string) bool {
	return Package_Captcha.VerifyString(id, captcha)
}

// Generate a new verification code ID
// 生成新的验证码ID
func (this *captcha) GenerateId() string {
	d := struct {
		CaptchaId string
	}{
		Package_Captcha.New(),
	}
	return d.CaptchaId
}
