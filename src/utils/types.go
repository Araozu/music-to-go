package utils

// The result of a Login
type AuthSuccess struct {
	Id            string `json:"id"`
	IsAdmin       bool   `json:"isAdmin"`
	Name          string `json:"name"`
	SubsonicSalt  string `json:"subsonicSalt"`
	SubsonicToken string `json:"subsonicToken"`
	Token         string `json:"token"`
	Username      string `json:"username"`
}
