package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// этот обработчик принимает только запросы, отправленные методом GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
	// продолжаем обработку запроса
	// ...
}

func requestPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "URL ===============\r\n"
	body += fmt.Sprintf("URL: %v\r\nTrimmed: %v\r\n", req.URL, req.URL.Path[1:])

	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"
	err := req.ParseForm()
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	}
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func StaticFilesHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, r.URL.Path[1:])
	http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, `./net-http.go`)
}

func golangPage(res http.ResponseWriter, req *http.Request) {
	fs := http.FileServer(http.Dir("./"))
	h := http.StripPrefix(`/golang/`, fs)
	//fs.ServeHTTP(res, req)

	h.ServeHTTP(res, req)
	//fmt.Fprintf(res, "FS: %#v\r\n", fs)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(`/api/`, apiPage)
	mux.HandleFunc(`/request`, requestPage)
	mux.HandleFunc(`/json`, JSONHandler)
	mux.HandleFunc(`/static`, StaticFilesHandler)
	mux.HandleFunc(`/golang/`, golangPage)
	//  Содержимое файла net-http.go будет отображено на главной странице
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
