package library

import "time"

var (
	Jwt_Token = Interface_Jwt(&JwtGo{MySigningKey: []byte(Config.App.Key)})
	I18n = Interface_I18n(&Go_i18n{})
	Cache = Interface_Cache(&Redis{})
)


// Jwt token interface
// JWT 令牌接口
type Interface_Jwt interface {
	CreateToken(interface{}) (string,  error)
	ParseToken(string) (interface{}, error)
	GetIdByToken(string) (uint, error)						// Get user ID
	GetIssuerByToken(string) (issuer string, err error)		// Get Username
}


// i18n interface
// i18n 接口
type Interface_I18n interface {
	Translate (ctx string, language string) string
}


// Cache interface
// 缓存接口
type Interface_Cache interface {
	GetString (key string) (string, error)
	SetString (key string, value string, time time.Duration)
	DeleteString (list ...string)
	GetHash (key string, field string) (string, error)
	SetHash (key string, value interface{})
	DeleteHash (key string, list ...string)
	ClearAllCache()
}
