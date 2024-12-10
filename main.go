// Package main is must for any go lang project
package main

import (
	"fmt"
	"net/http"
)

//import the required packages
func main(){
	//create a handler function 
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"Hello World")
	})

	//Start the Server
	fmt.Println("Starting server on localhost:8081")
	http.ListenAndServe(":8081",nil)

}