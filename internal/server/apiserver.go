package Server

import (
	"log"
	"net/http"
	"time"
	"web-token-server/docs"
)

//Run running server
func Run() {
	// Go-routine for finding expired token in tokens array and remove
	tokenStore()

	// Mux router
	mux := http.NewServeMux()
	docs.Include(mux)
	mux.HandleFunc("/new", NewToken)
	mux.HandleFunc("/info", TokenInfo)
	mux.HandleFunc("/check", CheckToken)
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
