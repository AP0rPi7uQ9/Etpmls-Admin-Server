package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CaptchaGetPicture(c *gin.Context)  {
	captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeHTTP(c.Writer, c.Request)
	return
}

func CaptchaGetOne(c *gin.Context)  {
	captchaId := library.CaptchaGenerateId()
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Get, captchaId)
	return
}