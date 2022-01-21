package main

import (
	"cutelinks/generator"
	"fmt"
	"net/http"
)

func main() {
	configuration:=GetConfig()
	fmt.Println("INMEMORYSTORAGE =",configuration.INMEMORYSTORAGE)
	s := generator.String(10)
	fmt.Println(s)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		if r.Method=="POST" {
        	fmt.Fprintf(w, "post url: "+url)}
		if r.Method=="GET" {
			fmt.Fprintf(w, "get url: "+url)}

	})
	http.ListenAndServe(":82",nil)
}
