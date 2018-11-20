package main

import (
	"github.com/sigtot/byggern-rest/handlers"
	"github.com/sigtot/byggern-rest/serial"
	"net/http"
)

const apiRoot = "/api/"
const port = ":8000"

const serialName = "/dev/ttyACM0"
const serialBaud = 9600
const serialStopBits = 2

func main() {
	serialConn, err := serial.CreateConnection(
		serialName,
		serialBaud,
		serialStopBits)
	if err != nil {
		panic(err)
	}
	defer serialConn.Close()
	handlers.SetSerialConnection(serialConn)

	http.HandleFunc(apiRoot+"motor", handlers.HandleMotorPost)
	http.HandleFunc(apiRoot+"servo", handlers.HandleServoPost)
	http.HandleFunc(apiRoot+"state", handlers.HandleStateGet)
	http.HandleFunc(apiRoot+"solenoid", handlers.HandleSolenoidKick)
	http.HandleFunc(apiRoot+"pid", handlers.HandlePIDParamsPost)

        http.Handle("/", http.FileServer(http.Dir("./byggern-frontend")))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
