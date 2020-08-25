package token

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Tokens map[string]Token

type Token struct {
	Token      string
	ExpireTime int64
}

func CreateToken(tokens Tokens) (tokenStr string) {
	tokenStr = uuid.NewV4().String()
	now := time.Now().Unix()
	expire := now + 300 // 5 min after
	tokens[tokenStr] = Token{tokenStr, expire}
	return tokenStr
}

func RemoveExpiredToken(tokens Tokens) {
	for tokenStr, t := range tokens{
		if t.ExpireTime < time.Now().Unix() {
			delete(tokens, tokenStr)
		}
	}
}
