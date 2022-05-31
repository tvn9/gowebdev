package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=uft-8")
	io.WriteString(w, `
	<form method="POST">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}
