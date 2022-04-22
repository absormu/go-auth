package dto

import "github.com/absormu/go-auth/app/entity"

func SetEmail(to string, subject string, message string) (req entity.Email) {

	req.To = to
	req.Subject = subject
	req.Message = message

	return
}
