package Server

import (
	"log"
	"net/http"
	"time"
	"web-token-server/docs"
	"web-token-server/pkg/generator"
)

//newToken generate new token
func newToken(name, host string) string {
	token := Store{
		Host:      host,
		ExpiresAt: time.Now().Add(time.Duration(conf.ExpiredAt) * time.Second),
		Name:      name,
		Token:     generator.UUIDV4(),
	}
	storeTokens = append(storeTokens, token)
	return token.Token
}

//Run running server
func Run() {
	// Go-routine for finding expired token in tokens array and remove
	go tokenStore()

	// Mux router
	mux := http.NewServeMux()
	docs.Include(mux)
	mux.HandleFunc("/new", NewToken)
	mux.HandleFunc("/info", TokenInfo)
	mux.HandleFunc("/check", ChekToken)
	mux.HandleFunc("/version", Version)

	//mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./docs/assets/"))))

	//Configuration server
	server := &http.Server{
		Addr:           ":" + conf.Port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Check TLS on/off
	if conf.TLS {
		log.Println("TLS HTTP Server listen Port", conf.Port, "version", version())
		log.Fatal(server.ListenAndServeTLS(conf.ServerCrt, conf.ServerKey))
	} else {
		log.Println("HTTP Server listen Port", conf.Port, "version", version())
		log.Fatal(server.ListenAndServe())
	}

}
