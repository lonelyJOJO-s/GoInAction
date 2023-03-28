package main

import "log"

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

}

func test1() {
	// write into standard log
	log.Println("message")

	//println -> os.Exit(1)
	log.Fatalln("fatal message")

	// println -> panic
	log.Panicln("panic message")
}
