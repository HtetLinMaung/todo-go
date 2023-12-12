package model

type User struct {
	UserID        int64  `json:"user_id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ProfileImage  string `json:"profile_image"`
	AccountStatus string `json:"account_status"`
}

type UserRequest struct {
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ProfileImage  string `json:"profile_image"`
	AccountStatus string `json:"account_status"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
