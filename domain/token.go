package domain

type Token struct{
	Username string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
}

