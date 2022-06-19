package auth

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type Authenticator struct {
}

func (a *Authenticator) Check(userID int64) bool {
	return userID > 0
}
