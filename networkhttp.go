package main
import (
	"fmt" 
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , this is a simple HTTP server in Go!")
	})
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	
}