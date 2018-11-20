package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sigtot/byggern-rest/serial"
	"log"
	"net/http"
)

type ValueInput struct {
	Value int `json:"value"`
}

type PIDParams struct {
	KP int `json:"KP"`
	KI int `json:"KI"`
	KD int `json:"KD"`
}

var conn serial.Connection

func SetSerialConnection(connection serial.Connection) {
	conn = connection
}

func HandleMotorPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	handleMotorBoxInput(w, r, "motor")
}

func HandleServoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	handleMotorBoxInput(w, r, "servo")
}

func HandlePIDParamsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var pidParams PIDParams
	if err := decoder.Decode(&pidParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err := conn.Write(fmt.Sprintf("{K_p=%d&K_i=%d&K_d=%d}", pidParams.KP, pidParams.KI, pidParams.KD))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func HandleStateGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := conn.Write("{GET}"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	uartResponse, err := conn.ReadLine()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(uartResponse))
}

func HandleSolenoidKick(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := conn.Write("{kick}"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
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

	err = conn.Write(fmt.Sprintf("{%s=%d}", inputKey, input.Value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
