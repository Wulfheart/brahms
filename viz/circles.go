package viz

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	"log"
	"net/http"
)

func Test() {
	http.Handle("/", http.HandlerFunc(circle))
	fmt.Println("Starting Webserver")
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(500, 500)
	s.Circle(250, 250, 50, "stroke:rgb(255,0,0); fill: rgb(255,0,0); fill-opacity: 0.5")
	s.Circle(300, 240, 50, "fill: rgb(255,0,0); fill-opacity: 0.5")
	s.Circle(250, 250, 20, "fill: rgb(255,255,0); fill-opacity: 0.5")
	s.End()
}
