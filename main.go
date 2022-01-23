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
	redis.SaveUrl("1234567890","wegiuwehgiuehriguheriuhergregw.ru")
	fmt.Println(redis.GetUrl("1234567890"))
	randomizer := generator.Randomizer{}
	gen := generator.New(&randomizer)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		url := strings.ToLower(r.FormValue("url"))
		if r.Method=="POST" {
			var tinyUrl string
			if configuration.INMEMORYSTORAGE{
				tinyUrl=redis.SaveUrl(gen.String(10), url)
			} else {
				tinyUrl=postgres.SaveUrl(gen.String(10), url)
			}
			fmt.Fprintf(w, tinyUrl)
		}
		if r.Method=="GET" {
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