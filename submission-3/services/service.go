package services

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

var (
	PATH_JSON = path.Join("services", "data.json")
)

func StartServer(port string) {
	http.HandleFunc("/", routeHandler)
	fmt.Printf("Server started at http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	data := ReadJSONFile()
	response := ParseResponse(data)
	tpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tpl.Execute(w, response)

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
