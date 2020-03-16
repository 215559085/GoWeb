package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
var handledCount int = 0  //because of goroutine,they have
var handledIcoRequest = 0
func httpHandler(writer http.ResponseWriter,request *http.Request){
	_ = request.ParseForm();
	fmt.Println("Handled HttpRequest: "+strconv.Itoa(handledCount))
	/*for k,v := range request.Form{
		fmt.Println("key: ",k)
		fmt.Println("value: ",strings.Join(v,""))
	}*/
	handledCount+=1
	_, _ = fmt.Fprintf(writer, "<h1>Hello GO web,You are the: ")
	_, _ = fmt.Fprintf(writer, strconv.Itoa(handledCount)+" quest</h1>")
	return
}
func icoHandler(writer http.ResponseWriter,r *http.Request){
	handledIcoRequest++
	fmt.Println("Handled IcoRequest: "+strconv.Itoa(handledIcoRequest))
	return
}
func main() {
    fmt.Println("Hello GoLang!")
	http.HandleFunc("/",httpHandler)
	http.HandleFunc("/favicon.ico",icoHandler);//not all browser need this icon
	err := http.ListenAndServe(":8080",nil);
	if err != nil{
       log.Fatal("error",err)
	}
}