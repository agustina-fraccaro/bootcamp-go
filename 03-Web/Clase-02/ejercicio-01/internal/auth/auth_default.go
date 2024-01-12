package auth

func NewAuthTokenBasic(token string) *AuthBasic {
	return &AuthBasic{
		Token: token,
	}
}

type AuthBasic struct {
	Token string
}

func (a *AuthBasic) Auth(token string) (err error) {
	if a.Token != token {
		return ErrAuthTokenInvalid
	}
	return nil
}
