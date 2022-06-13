package Server

import (
	"time"
)

var storeTokens []Store

//tokenStore finding expired token in tokens array and remove
func tokenStore() {
	for {
		for index, token := range storeTokens {
			if token.ExpiresAt.UnixMilli()-time.Now().UnixMilli() < 0 {
				if len(storeTokens) > index+1 {
					storeTokens = append(storeTokens[:index], storeTokens[index+1:]...)
				} else if len(storeTokens) == 1 {
					storeTokens = storeTokens[:0]
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

//isExpired checking expired token in tokens array
func isExpired(token string) bool {
	var expired bool
	for _, t := range storeTokens {
		if t.Token == token {
			expired = true
		}
	}
	return expired
}
