package response

type Register struct {
	Id int `json:"id"`
}

type Login struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
