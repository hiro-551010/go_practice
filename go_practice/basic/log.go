package main

import (
	"log"
	"os"
	"io"
)

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	multilogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multilogFile)
}

func main() {
	loggingSettings("test.log")
	// log.Println("logging")
	// log.Printf("%T %v", "test", "test")

	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalln("error!")
	_, err := os.Open("./aaa")
	if err != nil {
		log.Fatalln("Exit!!", err)
	}
}