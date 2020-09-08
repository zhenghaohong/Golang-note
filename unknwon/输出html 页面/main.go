package main

import (
	"fmt"
	"net/http"
)

const AddForm = `
	<form method="POST" action="/add">
	URL: <input type="text" name="url">
	<input type="submit" value="Add">
	</form>
`

func main()  {
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)

}


func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, AddForm)
	return
	//url := r.FormValue("url")
	//if url == "" {
	//
	//}
	//key := store.Put(url)
	//fmt.Fprintf(w, "http://localhost:8080/%s", key)
}
