package library

var (
	Jwt_Token = Interface_Jwt(&JwtGo{MySigningKey: []byte(Config.App.Key)})
	// Cache = Interface_Cache(&Redis)
)


// Jwt token interface
// JWT 令牌接口
type Interface_Jwt interface {
	CreateToken(interface{}) (string,  error)
	ParseToken(string) (interface{}, error)
	GetIdByToken(string) (uint, error)						// Get user ID
	GetIssuerByToken(string) (issuer string, err error)		// Get Username
}


// Cache interface
// 缓存接口
type Interface_Cache interface {

}

