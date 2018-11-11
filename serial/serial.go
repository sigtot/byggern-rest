package serial

import (
	"errors"
	"github.com/tarm/serial"
	"log"
)

const MaxNumBytes = 2048

type Connection struct {
	port     *serial.Port
	byteChan chan byte
}

func CreateConnection(name string, baudRate int, stopBits serial.StopBits) (Connection, error) {
	c := &serial.Config{Name: name, Baud: baudRate, StopBits: stopBits}
	port, err := serial.OpenPort(c)
	byteChan := make(chan byte)
	conn := Connection{port, byteChan}
	if err != nil {
		return conn, err
	}
	go conn.receiveBytes()
	return conn, err
}

func (c *Connection) ReadLine() (line string, err error) {
	var bytes []byte
	for i := 0; i < MaxNumBytes; i++ {
		b := <-c.byteChan
		if b == byte(13) {
			return string(bytes[:i]), nil
		} else {
			bytes = append(bytes, b)
		}
	}
	return "", errors.New("max line length reached")
}

func (c *Connection) Write(str string) error {
	_, err := c.port.Write([]byte(str))
	return err
}

func (c *Connection) WriteLine(line string) error {
	return c.Write(line)
}

func (c *Connection) Close() {
	c.port.Close()
}

func (c *Connection) receiveBytes() {
	for {
		buf := make([]byte, 128)
		n, err := c.port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < n; i++ {
			c.byteChan <- buf[i]
		}
	}
}
