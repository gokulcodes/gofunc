package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main(){
	r := httprouter.New()
	fmt.Println(r)
	r.GET("/", HomeHandler)
	fmt.Println("Hosted on 8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	fmt.Fprintf(w, "httprouting")
}
