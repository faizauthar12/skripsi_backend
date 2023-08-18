package email

import (
	"errors"
	"net/smtp"
	"os"
)

type DotEnv struct {
	EmailSender    string
	PasswordSender string
	EmailReceier1  string
	EmailReceier2  string
}

type loginAuth struct {
	username, password string
}

func ConfigDotEnvv() *DotEnv {

	dotEnv := &DotEnv{
		EmailSender:    os.Getenv("EMAIL_SENDER"),
		PasswordSender: os.Getenv("PASSWORD_SENDER"),
		EmailReceier1:  os.Getenv("EMAIL_RECEIVER_1"),
		EmailReceier2:  os.Getenv("EMAIL_RECEIVER_2"),
	}

	return dotEnv
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
