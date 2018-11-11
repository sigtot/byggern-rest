package handlers

import (
	"fmt"
	"github.com/sigtot/byggern-rest/serial"
	"log"
	"net/http"
)

const serialName = "/dev/serial/by-id/usb-Arduino__www.arduino.cc__0042_95334323430351A00182-if00"
const serialBaud = 9600
const serialStopBits = 2

func HandleMotorInput(w http.ResponseWriter, r *http.Request) {
	conn, err := serial.CreateConnection(
		serialName,
		serialBaud,
		serialStopBits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	err = conn.Write(fmt.Sprintf("{motor=%d}", 120)) // TODO: Get motor value from url
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func HandleServoInput(w http.ResponseWriter, r *http.Request) {
	conn, err := serial.CreateConnection(
		serialName,
		serialBaud,
		serialStopBits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer conn.Close()

	err = conn.Write(fmt.Sprintf("{servo=%d}", 30)) // TODO: Get servo value from url
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func HandleSolenoidKick(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement solenoid kick endpoint
}
