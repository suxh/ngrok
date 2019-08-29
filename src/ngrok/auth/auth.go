package auth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Auth struct {
	Tokens map[string]string
}

var tokenFile = "/etc/ngrok/htpasswd"

func SetTokenFile(file string) {
	tokenFile = file
}

func NewAuth() (*Auth, error) {
	auth := &Auth{
		Tokens: make(map[string]string),
	}

	if err := auth.load(tokenFile); err != nil {
		return nil, err
	}

	return auth, nil
}

func (this *Auth) load(tokenFile string) error {
	file, err := os.Open(tokenFile)
	if err != nil {
		err = fmt.Errorf("Failed to read auth token file %s: %v", tokenFile, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		toParse := strings.Split(text, "#")[0]
		fields := strings.Fields(toParse)
		if len(fields) != 2 {
			continue
		}
		this.Tokens[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (this *Auth) Auth(token string) error {
	fields := strings.Split(token, ":")
	if len(fields) != 2 {
		return errors.New("wrong format of token")
	}
	username := fields[0]
	password := fields[1]

	if val, ok := this.Tokens[username]; ok {
		if val == password {
			return nil
		}
	}
	return errors.New("user or password error")
}
