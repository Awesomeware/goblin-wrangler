package models

type GoogleSSOTokenReq struct {
	Credential   string `form:"credential"`
	G_CSRF_Token string `form:"g_csrf_token"`
}

type GoogleSSOTokenClaims struct {
	Claims map[string]interface{} `json:"claims"`
}
