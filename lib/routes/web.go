package routes

import (
	"github.com/chaseWilliams/family-map/lib/models"
	"log"
	"net/http"
	"time"
)

func FileServer(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	log.Println("got here")
	// try to open file
	d := http.Dir("./web")
	f, err := d.Open(r.URL.Path)
	if r.URL.Path == "/" {
		f, _ = d.Open("index.html")
	} else if err != nil {
		f, _ = d.Open("HTML/404.html")
		w.WriteHeader(http.StatusNotFound)
	}
	http.ServeContent(
		w,
		r,
		r.URL.Path,
		time.Now(),
		f,
	)
	return
}
