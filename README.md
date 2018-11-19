# REST API for TTK4155 Byggern
Restful api for reading and updating the state of a physical pinball game

## Dependencies
```bash
go get https://github.com/tarm/serial
```

## Usage
### Run it directly
```bash
go run main.go
```
### Or compile and run the binary
```bash
go build
./byggern-rest
```

## API Docs
The API has the api root /api/ and available endpoints 
* /api/servo (POST)
* /api/motor (POST)
* /api/solenoid (POST)
* /api/pid (POST)
* /api/state (GET)

### The following endpoints have been implemented so far:
#### Servo value (POST)
```bash
curl localhost:8000/api/servo -X POST -d '{"value":80}'
```

#### Motor reference (POST)
```bash
curl localhost:8000/api/motor -X POST -d '{"value":80}'
```

#### Solenoid kick (POST)
```bash
curl localhost:8000/api/solenoid -X POST
```

#### PID controller parameters (POST)
**Note**:
Control parameters are divided by 1000 and converted to float on the server.
Thus, sending `"Kp": 500` corresponds to setting `Kp` to 0.5 on the machine. 
```bash
curl localhost:8000/api/pid -X POST -d '{"KP": 500, "KI": 10, "KD": 100}'
```

#### State (GET)
```bash
curl localhost:8000/api/state
```
