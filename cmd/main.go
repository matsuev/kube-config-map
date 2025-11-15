package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var bindPort string
var bindPortInt int

func init() {
	var err error

	bindPort, _ = os.LookupEnv("BIND_PORT")
	name, _ := os.LookupEnv("APP_NAME")
	level, _ := os.LookupEnv("LOG_LEVEL")
	fmt.Println("APP_NAME:", name)
	fmt.Println("LOG_LEVEL:", level)
	fmt.Println("BIND_PORT:", bindPort)

	if bindPort == "" {
		bindPortInt = 8081
	} else {
		if bindPortInt, err = strconv.Atoi(bindPort); err != nil {
			bindPortInt = 8081
			log.Printf("config error: BIND_PORT: %s\n", bindPort)
			log.Printf("used default port :%d", bindPortInt)
		} else {
			if bindPortInt < 1024 || bindPortInt > 65000 {
				bindPortInt = 8081
				log.Printf("config error: BIND_PORT: %s\n", bindPort)
				log.Printf("used default port :%d", bindPortInt)
			}
		}
	}
}

func main() {

	http.HandleFunc("/api", rootHandler())

	log.Printf("service started on port :%d\n", bindPortInt)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", bindPortInt), nil); err != nil {
		log.Fatalln(err)
	}
}

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API"))
	}
}
