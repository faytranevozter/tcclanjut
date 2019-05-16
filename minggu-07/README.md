# GO Web Example
## 3 Contoh Golang Web Basic dari [GoWebExamples](https://gowebexamples.com)

### Http Server

1. Buat file `terserah.go` yang isinya  
```go
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
```

2. Jalankan melalui terminal `go run terserah.go`  
![Step 1](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/httpserver/step1.png)  
3. Kemudian lihat melalui broser `http://localhost:9000`  
![Step 2](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/httpserver/step2.png)  
4. Kemudian lihat (lagi) melalui broser `http://localhost:9000/about`  
![Step 3](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/httpserver/step3.png)  


### Routing dengan Gorilla/mux

1. Menginstall package gorilla/mux dengan menjalankan perintah `go get -u github.com/gorilla/mux`  
![Step 1](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/routing/step1.png)  

2. Membuat file `gorilla.go` yang isinya  
```go
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
```

3. Jalankan melalui terminal `go run gorilla.go`  
![Step 2](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/routing/step2.png)  
4. Kemudian lihat melalui broser `http://localhost:9000/users/fahrur`  
![Step 3](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/routing/step3.png)  
5. Kemudian lihat melalui broser `http://localhost:9000/users/panjul`  
![Step 4](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/routing/step4.png)  


### Assets and Files

1. Membuat file `assets.go` yang isinya  
```go
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
```

2. Jalankan melalui terminal `go run assets.go`  
![Step 1](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/assetsnfiles/step1.png)  
3. Kemudian lihat melalui broser `http://localhost:9000/static`  
![Step 2](https://raw.githubusercontent.com/faytranevozter/tcclanjut/master/minggu-07/assetsnfiles/step2.png)  
