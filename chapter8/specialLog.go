package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("fail to open file:", err)
	}

	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, file), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func test2() {
	Trace.Println("I have something records")
	Info.Println("specific information")
	Warning.Println("something to know")
	Error.Println("something has failed")
}
