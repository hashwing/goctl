pkg/server/apis/v1/type/login.go
package typev1

type LogInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}
