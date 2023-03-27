package main

import (
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

	a := arduino.Connect(mode, logger, "COM3")
	defer a.Port.Close()

	go func() {
		wp.Add(1)
		defer wp.Done()
		err := a.Debug(logger)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		wp.Add(1)
		defer wp.Done()
		err := console.Start(logger, a)
		if err != nil {
			log.Fatal(err)
		}
	}()
	wp.Wait()
}
