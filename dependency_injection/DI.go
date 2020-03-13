package main

import (
	"fmt"
	"io"
	"net/http"
)


//Greet write formatted hellow world with the writer dependency which implemented io.Writer we injected in 
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//MyGreeterHandler http handler 
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
