package main

import (
	"cutelinks/config"
	"cutelinks/generator"
	"cutelinks/postgres"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	configuration:= config.GetConfig()
	fmt.Println("INMEMORYSTORAGE =",configuration.INMEMORYSTORAGE)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		url := strings.ToLower(r.FormValue("url"))
		if r.Method=="POST" {
			tinyUrl:=postgres.SaveUrl(generator.String(10), url)
        		fmt.Fprintf(w, tinyUrl)
		}
		if r.Method=="GET" {
			fullUrl:=postgres.GetUrl(url)
			fmt.Fprintf(w, fullUrl)
		}

	})
	http.ListenAndServe(":"+configuration.PORT,nil)
}
