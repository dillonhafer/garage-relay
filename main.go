package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
)

func High() (err error) {
	err = rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	pin := rpio.Pin(25)
	pin.Output()
	pin.High()
	return nil
}

func Low() (err error) {
	err = rpio.Open()
	if err != nil {
		return err
	}
	defer rpio.Close()

	pin := rpio.Pin(25)
	pin.Output()
	pin.Low()
	return nil
}

func main() {
	var err error
	command := os.Args[1]
	if command == "high" {
		println("Setting PIN high")
		err = High()
	}

	if command == "low" {
		println("Setting PIN low")
		err = Low()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	println("You must specify 'high' or 'low'")
	println("(i.e. garage-relay high)")
	os.Exit(1)
}
