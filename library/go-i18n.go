package library

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
)

var (
	bundle *i18n.Bundle
)

// Initialization
// 初始化
// https://github.com/nicksnyder/go-i18n/tree/master/v2/example
func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	list, err := filepath.Glob("./storage/language/*.toml")
	if err != nil || len(list) < 1 {
		Log.Error("Failed to load language pack!")
		return
	}

	for _, v:=range list {
		bundle.MustLoadMessageFile(v)
	}
	return
}


type Go_i18n struct {}


// Translate
// 翻译
func (this *Go_i18n) Translate (s string, lang string) string {
	localizer := i18n.NewLocalizer(bundle, lang)
	ctx := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      s,
	})
	return ctx
}