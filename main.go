package main 

import ( 
    "net/http"
    "io/ioutil"
    "fmt"
)

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/save", writeSome)
    http.ListenAndServe(":8088", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("helloooooo"))
}

func writeSome( w http.ResponseWriter, r *http.Request) { 
    w.Write([]byte("wrote some"))
    err := ioutil.WriteFile("/tmp/test12.out",[]byte("wrote some to this file"), 0644)
    if err != nil {
        fmt.Println(err)
    }
}
    

