package Server

import (
	"log"
	"time"
	"web-token-server/pkg/generator"
)

var SessionTokenStorage []SessionTokenStorageStruct
var RefreshTokenStorage []RefreshTokenStorageStruct

//NewSessionToken generate new token
func NewSessionToken(name, host string) string {
	token := SessionTokenStorageStruct{
		Host:      host,
		ExpiresAt: time.Now().Add(time.Duration(conf.ExpiredAt) * time.Second),
		Name:      name,
		Token:     generator.UUIDV4() + "." + name,
	}
	SessionTokenStorage = append(SessionTokenStorage, token)
	return token.Token
}

func NewRefreshToken(nameSessionToken string) string {
	refreshToken := RefreshTokenStorageStruct{
		NameSessionToken: nameSessionToken,
		RefreshToken:     generator.UUIDV4(),
		ExpiresAt:        time.Now().Add(time.Duration(600) * time.Second),
	}
	RefreshTokenStorage = append(RefreshTokenStorage, refreshToken)
	return refreshToken.RefreshToken
}

//TimeOutLiveToken finding expired token in tokens array and remove
func TimeOutLiveToken() {
	for {
		for index, token := range SessionTokenStorage {

			if token.ExpiresAt.UnixMilli()-time.Now().UnixMilli() < 0 {
				if len(SessionTokenStorage) > index+1 {

					SessionTokenStorage = append(SessionTokenStorage[:index], SessionTokenStorage[index+1:]...)

				} else if len(SessionTokenStorage) == 1 {

					SessionTokenStorage = SessionTokenStorage[:0]
				}
			}
		}
		if len(SessionTokenStorage) == 0 {
			SessionTokenStorage = make([]SessionTokenStorageStruct, 0, 0)
		}
		time.Sleep(time.Second)
		log.Println(cap(SessionTokenStorage), len(SessionTokenStorage), SessionTokenStorage)
	}

}

//TimeOutLiveRefreshToken finding expired token in tokens array and remove
func TimeOutLiveRefreshToken() {
	for {
		for index, token := range RefreshTokenStorage {

			if token.ExpiresAt.UnixMilli()-time.Now().UnixMilli() < 0 {
				if len(RefreshTokenStorage) > index+1 {

					RefreshTokenStorage = append(RefreshTokenStorage[:index], RefreshTokenStorage[index+1:]...)

				} else if len(RefreshTokenStorage) == 1 {

					RefreshTokenStorage = RefreshTokenStorage[:0]
				}
			}
		}
		if len(RefreshTokenStorage) == 0 {
			RefreshTokenStorage = make([]RefreshTokenStorageStruct, 0, 0)
		}
		time.Sleep(time.Second)
		log.Println(cap(RefreshTokenStorage), len(RefreshTokenStorage), RefreshTokenStorage)
	}

}
