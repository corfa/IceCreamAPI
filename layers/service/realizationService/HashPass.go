package realizationService

import "golang.org/x/crypto/bcrypt"

func (s Service) HashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 7)
	return string(bytes)
}

func (s *Service) CheckHashPassword(pass string, hash string) bool {
	flag := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if flag == nil {
		return true
	}
	return false
}
