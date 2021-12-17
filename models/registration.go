package models

var (
	UserInfo map[string]*UserData
)

func init() {
	UserInfo = make(map[string]*UserData)

}

type UserData struct {
	FirstName   string `valid:"Required" json:"firstname" `
	LastName    string `valid:"Required" json:"lastname" `
	Phone       string `valid:"Mobile" json:"phone" `
	Email       string `valid:"Email" json:"email" `
	Password    string `valid:"Required" json:"password"`
	DateOfBrith string `valid:"Required" json:"dob"`
}
