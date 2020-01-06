package fp700

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/chyupa/fp700/utils"
	"go.bug.st/serial.v1"
	"log"
)

var Port string

func OpenPort() (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(Port, mode)

	if err != nil {
		fmt.Println("open error is", err)
		log.Fatal(err)
		return nil, errors.New("nu am putut deschide portul")
	}

	return port, nil
}

func SendCommand(code int, payload string) ([]byte, error) {
	command := utils.EncodeMessage(code, payload)

	port, err := OpenPort()
	if err != nil {
		return []byte(""), err
	}

	defer port.Close()

	port.Write(command)

	reader := bufio.NewReader(port)

	reply, err := reader.ReadBytes(0x03)
	if err != nil {
		return []byte(""), errors.New("Nu am putut citi raspunsul de la imprimanta")
	}

	startingIndex := utils.IndexOf(reply, 1)
	if startingIndex != -1 {
		reply = append(reply[startingIndex:])
	}

	return reply, nil
}
