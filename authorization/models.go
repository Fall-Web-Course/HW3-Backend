package authorization

type JWTToken struct {
	AccessToken   string
	RefreshToken  string
	AccessExpire  int64
	RefreshExpire int64
}
