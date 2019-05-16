package main

// import package "net/http"
import "net/http"

func main() {
	// membuat folder assets menjadi bisa diakses via http://localhost:9000/static
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// port yg digunakan, yang nantinya diakses di http://localhost:9000
	http.ListenAndServe(":9000", nil)
}