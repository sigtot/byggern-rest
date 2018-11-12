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
* /api/servo
* /api/motor
* /api/solenoid

### The following endpoints have been implemented so far:
#### Servo value (POST)
```bash
curl localhost:8000/api/servo -X POST -d '{"value":80}'
```

#### Motor value (POST)
```bash
curl localhost:8000/api/motor -X POST -d '{"value":80}'
```
