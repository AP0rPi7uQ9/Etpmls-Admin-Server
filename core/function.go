package core

import (
	"Etpmls-Admin-Server/library"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
)

func IsDebug() bool {
	if strings.ToUpper(os.Getenv("DEBUG")) == "TRUE" {
		return true
	}
	return false
}


// Get token by header Or query
// 从header或query中获取token
func GetToken(c *gin.Context) (token string, err error) {
	// Get Query Token
	token, b := c.GetQuery("token")
	if b {
		return token, err
	}

	// Get Header Token
	token = c.GetHeader("X-Token")
	if len(token) != 0 {
		return token, err
	}

	token = c.GetHeader("Token")
	if len(token) != 0 {
		return token, err
	}

	return token, errors.New("Token获取失败！")
}


// Translate
// 翻译
func Translate(c *gin.Context, ctx string) string {
	lang := c.GetHeader("language")
	if lang == "" {
		lang = "en"
	}
	return library.I18n.Translate(ctx, lang)
}