package bluemovement

import (
	"encoding/gob"
	"log/slog"
	"os"

	"github.com/taylormonacelli/goldbug"
)

type User struct {
	Username string
	Password string
}

func Main() int {
	goldbug.SetDefaultLoggerText(slog.LevelDebug)
	var err error

	user := User{
		"zola",
		"supersecretpassword",
	}

	err = encodeUser(user)
	if err != nil {
		return 1
	}

	err = decodeUser()
	if err != nil {
		return 1
	}

	return 0
}

func decodeUser() error {
	file, err := os.Open("user.gob")
	if err != nil {
		slog.Debug("file access", "error", err.Error())
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var user User

	if err := decoder.Decode(&user); err != nil {
		slog.Debug("decode", err.Error())
		return err
	}

	slog.Debug("user", "username", user.Username, "password", user.Password)

	return nil
}

func encodeUser(user User) error {
	file, _ := os.Create("user.gob")
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(user)

	return nil
}
