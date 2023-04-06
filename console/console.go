package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tmp/arduino"
)

var s string
var massive_message_int [16]int

func Start(logger *log.Logger, a arduino.Arduino) error {
	//data structure: [team number], [command number], [len of payload of elem massive], [payload massive]
	fmt.Println("Message to arduino must be: Number of team should be done -\"1 10\" or \"1-3 10\" <- This message will be sent to team 1,2,3")
	for {
		logger.SetPrefix("Console input: ")
		buff := bufio.NewReader(os.Stdin)
		message, err := buff.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		message = strings.Trim(message, "\n\t\r")
		fmt.Println(message)
		massive_message := strings.Split(message, " ")
		fmt.Println(massive_message)
		for i := 0; i < 2; i++ {
			tmp, err := strconv.Atoi(massive_message[i])

			if err != nil {
				fmt.Println(massive_message[i])
				fmt.Println("Error of parsing ")
				continue
			}
			massive_message_int[i] = tmp
		}
		if massive_message_int[0] < 1 && massive_message_int[0] > 6 {
			fmt.Println("Team number must be from 1 to 6")
			continue
		}
		if len(massive_message) == 2 {
			massive_message_int[3] = 0
			for i := 0; i < 3; i++ {
				s := strconv.Itoa(massive_message_int[i])
				a.Dataload(logger, s+"\n")
			}
		}
	}
}
