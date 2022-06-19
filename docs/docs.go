package docs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("docs/templates/*"))

func parseDocs() DocStruct {
	jsonFile, errOpenJson := os.Open("./docs/docs.json")
	if errOpenJson != nil {
		log.Println("Error open json file", errOpenJson)
	}

	defer func(jsonFile *os.File) {
		errCloseJson := jsonFile.Close()
		if errCloseJson != nil {
			log.Println("Error close config file", errCloseJson)
		}
	}(jsonFile)

	byteConfig, errByteConfig := ioutil.ReadAll(jsonFile)
	if errByteConfig != nil {
		log.Println("Err read byte", errByteConfig)
	}
	docs := DocStruct{}
	err := json.Unmarshal(byteConfig, &docs)
	if err != nil {
		log.Println("Error parse config", err)
	}
	return docs
}

func Docs(w http.ResponseWriter, r *http.Request) {

	errExecuteTmpl := tmpl.ExecuteTemplate(w, "Main", parseDocs())
	if errExecuteTmpl != nil {
		log.Println(errExecuteTmpl)
	}
}

func Include(mux *http.ServeMux) {
	mux.HandleFunc("/docs", Docs)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./docs/assets/"))))
}
