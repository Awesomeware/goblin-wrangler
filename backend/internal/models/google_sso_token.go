package models

type GoogleSSOTokenReq = GoogleSSOToken

type GoogleSSOToken struct {
	Credential   string `form:"credential"`
	G_CSRF_Token string `form:"g_csrf_token"`
}

type GoogleSignupReq struct {
	Token    GoogleSSOToken
	Username string
}

type GoogleSSOTokenClaims struct {
	Claims map[string]interface{} `json:"claims"`
}
