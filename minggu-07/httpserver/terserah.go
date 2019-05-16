package main

// import package, yg penting adalah "net/http"
import (
	"fmt"
	"net/http"
)

func main() {
	// tanda "/" berarti halaman utama
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// menampilkan tulisan 
		fmt.Fprintf(w, "Hello, my name is faytranevozter!")
	})

	// tanda "/about" berarti halaman http://localhost/about
	http.HandleFunc("/about", func (w http.ResponseWriter, r *http.Request) {
		// menampilkan tulisan (yg lain)
		fmt.Fprintf(w, "About someone here!")
	})

	// .. dan lain-lain

	// port yg digunakan, yang nantinya diakses di http://localhost:9000
	http.ListenAndServe(":9000", nil)
}