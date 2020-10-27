package main

import (
   "fmt"
   "net/http"
   "strings"
   "log"
)


func sayHelloName(w http.ResponseWriter, r *http.Request) {
   r.ParseForm()     // parse arguments
   fmt.Println(r.Form)     // print form information to output
   fmt.Println("path", r.URL.Path)
   fmt.Println("scheme", r.URL.Scheme)
   fmt.Println(r.Form["url_long"])
   for k, v := range r.Form {
      fmt.Println("key:", k)
      fmt.Println("val:", strings.Join(v, ""))
   }
   fmt.Fprintf(w, "Hello Joe Boy!")
}

func main() {
   http.HandleFunc("/", sayHelloName)     // router
   err := http.ListenAndServe(":9090", nil)  // set listen port
   if err != nil {
      log.Fatal("ListenAndServer: ", err)
   }
}

