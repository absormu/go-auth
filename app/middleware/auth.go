package middleware

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	CompanyID int    `json:"company_id"`
	SellerID  int    `json:"seller_id"`
	RoleID    int    `json:"role_id"`
	jwt.StandardClaims
}
