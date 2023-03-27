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

	for {
		logger.SetPrefix("Console input: ")
		output := ""
		fmt.Scan(&s)
		if len(s) == 1 {
			output = s + " "
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

			for i := i1; i <= i2; i++ {
				tmp := strconv.Itoa(i)
				output = tmp + " "
				output += s
				a.Dataload(logger, output)
			}
		}
	}
}
