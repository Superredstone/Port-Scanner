package PortScanner

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

func Scan(IP string) error {
	if net.ParseIP(IP) == nil {
		PortParagraphTxt = PortParagraphTxt + "\nInvalid IP address"

		return errors.New("invalid ip address")
	}

	if Timeout == "" {
		Timeout = "19"
	}

	t, err := strconv.ParseInt(Timeout, 10, 64)
	if err != nil {
		PortParagraphTxt = PortParagraphTxt + "\nError, invalid timeout time"
		return err
	}

	for i := 1; i < 65535; i++ {
		progress++

		timeout := time.Duration(time.Duration(t) * time.Millisecond)
		port := strconv.Itoa(i)

		_, err = net.DialTimeout("tcp", IP+":"+port, timeout)
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
