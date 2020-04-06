package model

import (
	"go-project/common"
	db "go-project/database"
)

// User model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin"`
}

// Auth : User authentication
func (u *User) Auth(password string) bool {
	return common.CheckPasswordHash(password, u.Password)
}

// Create : Create User
func (u *User) Create() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	pass, err := common.HashPassword(u.Password)
	if err != nil {
		return err
	}

	conn.QueryRow("INSERT INTO users (username, password, name) VALUES ($1, $2, $3)", u.Username, pass, u.Name)

	return nil
}

// Update : Update User
func (u *User) Update() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `UPDATE users SET 
		username = $1,
		name = $2,
		is_admin = $3
	WHERE id = $4`
	conn.QueryRow(sql, u.Username, u.Name, u.IsAdmin, u.ID)

	return nil
}

// Delete : Delete User
func (u *User) Delete() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `DELETE FROM users WHERE id = $1`
	conn.QueryRow(sql, u.ID)

	return nil
}

// GetUser : Get User using id
func (u *User) GetUser(id int) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `SELECT id,username,password,name,is_admin FROM users WHERE id = $1 LIMIT 1`
	row := conn.QueryRow(sql, id)

	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		return err
	}

	return nil
}

// GetByUsername : Get User using username
func (u *User) GetByUsername(username string) error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	sql := `SELECT id,username,password,name,is_admin FROM users WHERE username = $1 LIMIT 1`
	row := conn.QueryRow(sql, username)

	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.IsAdmin)
	if err != nil {
		return err
	}

	return nil
}

// Validate : User validation
func (u *User) Validate() (map[string]string, bool) {
	errMessages := map[string]string{}
	err := false

	switch {
	case u.Name == "":
		errMessages["name"] = "Nama can't be empty."
		err = true
		fallthrough
	case u.Username == "":
		errMessages["username"] = "Username can't be empty."
		err = true
		fallthrough
	case u.Password == "":
		errMessages["password"] = "Password can't be empty."
		err = true
	}

	return errMessages, err
}

// Get All Users
func GetAllUsers() ([]User, error) {
	var users []User
	var user User

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id,username,name,is_admin FROM users ORDER BY id DESC")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Name, &user.IsAdmin)
		users = append(users, user)
	}

	return users, nil
}
