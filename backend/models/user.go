package models

type User struct {
	UserName string   `bson:"username" json:"username"`
	Password string   `bson:"password" json:"password"`
	Roles    []string `bson:"roles" json:"roles"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Roles []string `json:"roles"`
	Name  string   `json:"name"`
}
