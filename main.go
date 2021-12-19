package main

import (
	"fmt"
	"github.com/dhikaroofi/go/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// func BubbleSort(array[] float32)[]float32 {
// 	arrayLength :=  len(array)-1
// 	for i:=0; i < arrayLength; i++ {
// 	   for j:=0; j < arrayLength-i; j++ {
// 		  if (array[j] < array[j+1]) {
// 			 array[j], array[j+1] = array[j+1], array[j]
// 		  }
// 	   }
// 	}
// 	return array
//  }
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, "Hello world!")
}

func main() {

	route := mux.NewRouter()
	app := app.Mojoo{}
	app.Initialize(route)
	log.Print("App is running on ")
	log.Fatal(http.ListenAndServe(":8080", route))

}
