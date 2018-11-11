package main

import (
	"github.com/sigtot/byggern-rest/handlers"
	"net/http"
)

const apiRoot = "/api/"
const port = ":8000"

func main() {
	http.HandleFunc(apiRoot+"motor", handlers.HandleMotorInput)
	http.HandleFunc(apiRoot+"servo", handlers.HandleServoInput)
	http.HandleFunc(apiRoot+"solenoid", handlers.HandleSolenoidKick)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}