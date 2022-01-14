package user

type User struct {
	// gorm.Model
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	IsAdmin		bool	`json:"id_admin"`
}