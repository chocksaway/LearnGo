package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
	"strconv"
)

/*
	Check for exactly two arguments
*/
func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_MAIL, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	if len(os.Args) == 1 {
		io.WriteString(os.Stderr, "Please give one or more arguments\n")
		os.Exit(1)
	} else if len(os.Args) != 3 {
		io.WriteString(os.Stderr, "Two arguments expected\n")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	err = returnError(min, max)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		if err.Error() == "min and max are the same" {
			fmt.Println("This is not good.....")
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	log.Println("LOG_MAIL: ", programName, " finished")
}

func returnError(a, b float64) error {
	if a == b {
		err := errors.New("min and max are the same")
		return err
	} else {
		return nil
	}
}
