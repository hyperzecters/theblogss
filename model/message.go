package model

import (
	db "go-project/database"
)

// Message model
type Message struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Create : Create Message
func (m *Message) Create() error {
	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.QueryRow("INSERT INTO messages (name, email, message) VALUES ($1, $2, $3)", m.Name, m.Email, m.Message)

	return nil
}

// Validate : Message validation
func (m *Message) Validate() (map[string]string, bool) {
	errMessages := map[string]string{}
	err := false

	switch {
	case m.Name == "":
		errMessages["name"] = "Name can't be empty."
		err = true
		fallthrough
	case m.Email == "":
		errMessages["email"] = "Email can't be empty."
		err = true
		fallthrough
	case m.Message == "":
		errMessages["message"] = "Message can't be empty."
		err = true
	}

	return errMessages, err
}
