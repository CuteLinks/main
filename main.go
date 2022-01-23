package main

import (
	"cutelinks/config"
	"cutelinks/generator"
	"cutelinks/postgres"
	"cutelinks/redis"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	configuration:= config.GetConfig()
	fmt.Println("INMEMORYSTORAGE =",configuration.INMEMORYSTORAGE)
	randomizer := generator.Randomizer{}
	gen := generator.New(&randomizer)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.Method=="POST" {
			url := strings.ToLower(r.FormValue("url"))
			var tinyUrl string
			if configuration.INMEMORYSTORAGE{
				tinyUrl=redis.SaveUrl(gen.String(10), url)
			} else {
				tinyUrl=postgres.SaveUrl(gen.String(10), url)
			}
			fmt.Fprintf(w, tinyUrl)
		}
		if r.Method=="GET" {
			url := r.FormValue("url")
			var fullUrl string
			if configuration.INMEMORYSTORAGE{
				fullUrl=redis.GetUrl(url)
			} else {
				fullUrl=postgres.GetUrl(url)
			}
			fmt.Fprintf(w, fullUrl)
		}

	})
	http.ListenAndServe(":"+configuration.PORT,nil)
}