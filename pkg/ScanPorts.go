package PortScanner

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

func Scan(IP string) error {
	if IP == "" {
		return errors.New("no ip given")
	}

	for i := 1; i < 65535; i++ {
		timeout := time.Duration(19 * time.Millisecond)
		port := strconv.Itoa(i)

		_, err := net.DialTimeout("tcp", IP+":"+port, timeout)
		if err != nil {
			//fmt.Printf("%s %s %s\n", IP, "not responding", err.Error())
		} else {
			fmt.Printf("%s %s %s\n", IP, "responding on port:", port)
			PortParagraphTxt = PortParagraphTxt + "\n" + "Found Port " + IP + ":" + port
		}
	}

	PortParagraphTxt = PortParagraphTxt + "\nPort scan finished"

	return nil
}
