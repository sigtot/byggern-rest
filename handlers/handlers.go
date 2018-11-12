package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sigtot/byggern-rest/serial"
	"log"
	"net/http"
)

const serialName = "/dev/ttyACM0"
const serialBaud = 9600
const serialStopBits = 2

type ValueInput struct {
	Value int `json:"value"`
}

func HandleMotorInput(w http.ResponseWriter, r *http.Request) {
	handleMotorBoxInput(w, r, "motor")
}

func HandleServoInput(w http.ResponseWriter, r *http.Request) {
	handleMotorBoxInput(w, r, "servo")
}

func HandleSolenoidKick(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement solenoid kick endpoint
}

func handleMotorBoxInput(w http.ResponseWriter, r *http.Request, inputKey string) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var input ValueInput
	err := decoder.Decode(&input)
	if err != nil {
		log.Println(err)
	}

	conn, err := serial.CreateConnection(
		serialName,
		serialBaud,
		serialStopBits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = conn.Write(fmt.Sprintf("{%s=%d}", inputKey, input.Value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}

	err = conn.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("error" + err.Error())
	}
}
