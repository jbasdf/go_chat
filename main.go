package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}

func serveError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Internal Server Error")
	io.WriteString(w, err.Error())
}

func main() {
	log.Println("Starting Server")

	flag.Parse()

	// API Routes
	r := mux.NewRouter()
	r.HandleFunc("/api/geo_jots", GeoJotsHandler).Methods("GET")
	r.HandleFunc("/api/geo_jots/{id}", ShowGeoJotHandler).Methods("GET")
	r.HandleFunc("/api/geo_jots", CreateGeoJotHandler).Methods("POST")
	r.HandleFunc("/api/geo_jots/{id}", UpdateGeoJotHandler).Methods("PUT")
	r.HandleFunc("/api/geo_jots/{id}", DeleteGeoJotHandler).Methods("DELETE")
	http.Handle("/api/", r)

	// Serve Static Files
	http.Handle("/", http.FileServer(http.Dir("./public/")))

	// Handle websockets
	go h.run()
	http.Handle("/ws", websocket.Handler(wsHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Listening on: " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//"time"
//"fmt"
//eventsource "github.com/antage/eventsource/http"
// func main() {
//   // es := eventsource.New(nil)
//   // defer es.Close()
//   // http.Handle("/events", es)
//   // go func() {
//   //   fmt.Println("listening...")
//   //   id := 1
//   //   for {
//   //       es.SendMessage("tick", "tick-event", strconv.Itoa(id))
//   //       id++
//   //       time.Sleep(2 * time.Second)
//   //       fmt.Println("Sending id...")
//   //   }
//   // }()

// }
