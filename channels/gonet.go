package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
//	"time"
	"log"
	"os"
//	"strconv"
)

func main(){
	container, err := http.Get("http://infomatte.herokuapp.com")
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	fmt.Printf("%+v\n", container)
//	time.Sleep(1 * time.Second)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	msg, errr := ioutil.ReadFile("./gonet.go")
	if errr != nil {
		log.Fatal(errr)	
	}
	fmt.Println(string(msg))
	crt, capture := ioutil.TempFile(".", "random")
	if capture != nil {
		log.Fatal(capture)	
	}
	fmt.Println("Succeeded", crt)
	defer os.Remove(crt.Name())
}
