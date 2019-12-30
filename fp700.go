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

//func read() {
//
//	// fmt.Println(strconv.ParseInt("0047", 16, 0))
//
//	// fmt.Println(strings.TrimLeft("003", "0"))
//
//	command := utils.EncodeMessage(125, "20\t2\t18\t")
//
//	port, _ := OpenPort()
//
//	defer port.Close()
//
//	port.Write(command)
//
//	reader := bufio.NewReader(port)
//
//	for {
//
//		readCommand := utils.EncodeMessage(125, "21\t\t\t")
//
//		port.Write(readCommand)
//
//		// 	fmt.Println("read command writtern ", bytesWritten)
//		// 	// fmt.Println("read command error is", err)
//
//		time.Sleep(time.Millisecond * 100)
//		reply, _ := reader.ReadBytes(0x03)
//
//		// 	// fmt.Println("readint error reader is", eror)
//		// 	// fmt.Println("this is how much I read", reply)
//		// 	// fmt.Println("I can read this amount", reader.Buffered())
//		indexOfByte := utils.IndexOf(reply, 0x22)
//		if indexOfByte > -1 {
//
//		}
//		if len(reply) > 2 {
//			msg := utils.DecodeMessage(reply[1:])
//			split := strings.Split(msg, "\t")
//			if len(split) > 1 {
//				fmt.Println(split[1])
//			}
//
//		}
//		// 	// fmt.Println("this is data bytes", reply)
//
//		// 	// response = append(response, reply...)
//		// 	// if len(reply) < 0 {
//		// 	// 	break
//		// 	// }
//	}
//
//	// fmt.Println("full response", response)
//
//}