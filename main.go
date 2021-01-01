package main

import (
  "github.com/gorilla/mux"
  "log"
)

// データの型となるモデル（struct）
type Book struct {
  ID      string `json:"id"`
  Title   string `json:"title"`
  Author  *Author `json:"author"`
}

type Author struct {
  FirstName string `json:"firstname"`
  LastName  string `json:"lastname"`
}

// main処理
func main() {
  //ルーターのイニシャライズ
  r := mux.NewRouter()

  //エンドポイント
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", r))
}
