package Server

//tokenStore finding expired token in tokens array and remove
func tokenStore() {

	go TimeOutLiveToken()
	go TimeOutLiveRefreshToken()

}

//isExpired checking expired token in tokens array
func isExpired(token string) bool {
	var expired bool
	for _, t := range SessionTokenStorage {
		if t.Token == token {
			expired = true
		}
	}
	return expired
}
