package request

type User struct {
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mail      string `json:"mail"`
	Sex       string `json:"sex"`
	Birthday  string `json:"birthday"`
	Weight    string `json:"weight"`
	Height    string `json:"height"`
}
