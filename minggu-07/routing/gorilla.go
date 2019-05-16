package main

// import package, yg penting adalah "net/http" dan "gorilla/mux"
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// menginisiasi router dengan mux
	r := mux.NewRouter()
	// membuat routing dengan segment dinamis
	r.HandleFunc("/users/{nama}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // mengambil variabel-variabel dari url
		nama := vars["nama"] // mengambil variabel nama
		// menampilkan tulisan
		fmt.Fprintf(w, "Halo, %s, ini halaman privasi anda.", nama)
	})

	// port yg digunakan, mirip seperti http server biasa, hanya berbeda pada parameter kedua
	http.ListenAndServe(":80", r)
}