package routes

import (
	"log"
	"net/http"
)

func FileServer(w http.ResponseWriter, r *http.Request) (err error) {
	log.Println("got here")
	// try to open file
	/*
		d := http.Dir("./web")
		f, err := d.Open(r.URL.Path)
		if err != nil {
			util.WriteNotFound(w)
			return
		}
		contents, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		w.Write(contents)
	*/
	http.ServeFile(w, r, "./web")

	return
}
