package main

import (
	"flag"
	"fmt"
	"go.bug.st/serial"
	"log"
	"os"
	"sync"
	"tmp/arduino"
	"tmp/console"
)

func main() {
	mode := &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}
	var wp sync.WaitGroup
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "INFO: ", log.Ldate|log.Ltime)
	comport := flag.String("port", "", "Settings arduino port")
	flag.Parse()

	if len(*comport) == 0 {
		ports, err := serial.GetPortsList()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Error. You need to select one of available ports using the flag \"--port\"\n %s", ports)
		os.Exit(1)
	}
	a := arduino.Connect(mode, logger, *comport)
	defer a.Port.Close()
	wp.Add(1)
	go func() {

		defer wp.Done()
		// this function listening COM port and logging data.
		err := a.Debug(logger)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wp.Add(1)
	go func() {

		defer wp.Done()
		// This function start interface for people.
		err := console.Start(logger, a)
		if err != nil {
			log.Fatal(err)
		}

	}()
	wp.Wait()
}
