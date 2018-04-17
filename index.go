package main

import(
	"net/http"
	"log"
	"fmt"
	"api/test"
)



func parsar(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	err := test.ParserValue(r.Form)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}else{
		fmt.Fprintf(w,"success")
	}

}

func main() {
	http.HandleFunc("/",parsar)
	test.CreateTable()
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("ListenAndServe: ",err)
	}
}
