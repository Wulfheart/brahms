package viz

import (
	"fmt"
	"log"
	"net/http"
	"wulfheart/brahms/score"
)

var s score.Score

func Test(score score.Score) {
	s = score
	http.Handle("/", http.HandlerFunc(circle))
	fmt.Println("Starting Webserver")
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	Render(w, s)
}
