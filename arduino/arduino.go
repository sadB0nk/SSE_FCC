package arduino

import (
	"bufio"
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
	"strings"
)

type Arduino struct {
	Port serial.Port
}

func Connect(mode *serial.Mode, logger *log.Logger, port string) (a Arduino) {
	logger.SetPrefix("Arduino startup: ")
	logger.Printf("Starting arduino connect on Port: %s, with baudrate: %d\n", port, mode.BaudRate)
	fmt.Printf("Starting arduino connect on Port: %s, with baudrate: %d\n", port, mode.BaudRate)
	conn, err := serial.Open(port, mode)
	if err != nil {
		logger.Println(err)
		log.Fatal(err)
	}
	a.Port = conn
	return a
}
func (a Arduino) Debug(logger *log.Logger) error {
	reader := bufio.NewReader(a.Port)

	for {
		str, err := reader.ReadString('\n')
		if len(str) > 1 && str[len(str)-1] == '\n' && str[len(str)-2] == '\n' {
			str = str[:len(str)-2]
		}
		if err != nil {
			logger.SetPrefix("Error of reading ")
			logger.Println(err)
			continue
		}
		str = strings.Trim(str, "\n")

		if str[0] == '1' {
			fmt.Println(str[1:])
		}
		logger.SetPrefix("Arduino debug: ")
		logger.Println(str[1:])
	}

}
func (a Arduino) Dataload(logger *log.Logger, data string) error {
	n, err := io.WriteString(a.Port, data)
	if err != nil {
		return err
	}
	logger.SetPrefix("Arduino load: ")
	data = strings.Trim(data, "\n")
	logger.Printf("On Arduino was write %d bytes and message: %s\n", n, data)
	fmt.Printf("On Arduino was write %d bytes and message: %s\n", n, data)
	return nil
}
