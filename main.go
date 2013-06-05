package main

import (
  "code.google.com/p/go.net/websocket"
  eventsource "github.com/antage/eventsource/http"
  "flag"
  "log"
  "net/http"
  "text/template"
  "os"
  "strconv"
  "time"
  "fmt"
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

  es := eventsource.New(nil)
  defer es.Close()
  http.Handle("/events", es)
  go func() {
    fmt.Println("listening...")
    id := 1
    for {
        es.SendMessage("tick", "tick-event", strconv.Itoa(id))
        id++
        time.Sleep(2 * time.Second)
        fmt.Println("Sending id...")
    }
  }()


  if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }

}