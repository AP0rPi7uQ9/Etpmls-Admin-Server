package core

import (
	"Etpmls-Admin-Server/library"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

	LogError.Output(MessageWithLineNum("Token acquisition failed!"))
	return token, errors.New("Token acquisition failed！")
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


// Message(or Error) with line number
// 消息(或错误)带行号
func MessageWithLineNum(msg string) string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filepath.Dir(file))
	sourceDir := strings.ReplaceAll(dir, "\\", "/")

	var list []string
	for i := 1 ; i < 20; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && strings.HasPrefix(file, sourceDir) {
			list = append(list, file + ":" + strconv.Itoa(line))
		} else {
			break
		}
	}
	return strings.Join(list, " => ") + " => Message: " + msg
}