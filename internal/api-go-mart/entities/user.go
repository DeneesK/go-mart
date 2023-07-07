package entities

type User struct {
	ID            int64  `json:"-"`
	Login         string `json:"json"`
	Password      string `json:"password"`
	CryptPassword []byte `json:"-"`
}
