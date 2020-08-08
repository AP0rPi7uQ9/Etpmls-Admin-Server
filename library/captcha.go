package library

import (
	"github.com/dchest/captcha"
)

func CaptchaGenerateId() string {
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	return d.CaptchaId
}