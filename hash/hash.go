package hash

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		log.Panic().Err(err).Msg("can`t hash user password")
	}
	return string(bytes)
}

//func Verify(hashedPassword, enteredPassword string) bool {
//
//	errPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
//	if errPassword == nil {
//		return true
//	}
//	return false
//}
