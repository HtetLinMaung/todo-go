package service

import (
	"database/sql"

	"github.com/HtetLinMaung/todo/internal/model"
	"github.com/HtetLinMaung/todo/internal/utils"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) AddUser(data *model.UserRequest) (int64, error) {
	var userId int64

	stmt, err := s.db.Prepare("insert into users (name, username, password, role, email, phone, profile_image, account_status) values ($1, $2, $3, $4, $5, $6, $7, $8) returning user_id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(data.Name, data.Username, hashedPassword, data.Role, data.Email, data.Phone, data.ProfileImage, data.AccountStatus).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserService) IsUserExists(username string) (bool, error) {
	var exists bool

	stmt, err := s.db.Prepare("select exists(select 1 from users where username = $1 and deleted_at is null)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	var user model.User

	stmt, err := s.db.Prepare("select user_id, name, username, password, role, email, phone, profile_image, account_status from users where username = $1 and deleted_at is null")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.UserID, &user.Name, &user.Username, &user.Password, &user.Role, &user.Email, &user.Phone, &user.ProfileImage, &user.AccountStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			// No user was found, return nil user and no error
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
