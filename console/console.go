package console

import (
	"fmt"
	"strconv"
	"strings"

	"log"

	"tmp/arduino"
)

var s string

func Start(logger *log.Logger, a arduino.Arduino) error {
	fmt.Println("Message to arduino must be: Number of team should be done -\"1 10\" or \"1-3 10\" <- This message will be sent to team 1,2,3")
	for {
		logger.SetPrefix("Console input: ")
		output := ""
		fmt.Scan(&s)
		if len(s) == 1 {
			output = s + " "
			i1, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error of parsing, num of team.")
				continue
			}
			if i1 < 1 || i1 > 6 {
				fmt.Println("Team number should be from 1 to 6")
				continue
			}
			fmt.Scan(&s)
			output += s

			a.Dataload(logger, output)
			continue
		}

		if len(s) == 3 {
			tmp := strings.Split(s, "-")
			fmt.Scan(&s)
			i1, err := strconv.Atoi(tmp[0])
			if err != nil {
				fmt.Println("Error of parsing first int\t", err)
				continue
			}
			i2, err := strconv.Atoi(tmp[1])
			if err != nil {
				fmt.Println("Error of parsing second int\t", err)
				continue
			}
			if i2 < 1 || i2 > 6 || i1 < 1 || i1 > 6 {
				fmt.Println("Team number should be from 1 to 6")
				continue
			}
			for i := i1; i <= i2; i++ {
				tmp := strconv.Itoa(i)
				output = tmp + " "
				output += s
				a.Dataload(logger, output+"\n")
			}
			continue
		}
		fmt.Println("Team number should be from 1 to 6")
	}
}
