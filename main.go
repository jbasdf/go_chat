package main

import (
  "code.google.com/p/go.net/websocket"
  "flag"
  "log"
  "net/http"
  "text/template"
  "os"
)

//var addr = flag.String("addr", ":80", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTempl.Execute(c, req.Host)
}

func main() {
  flag.Parse()
  go h.run()
  http.HandleFunc("/", homeHandler)
  http.Handle("/ws", websocket.Handler(wsHandler))
  if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}