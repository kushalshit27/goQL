package goQL

import (
	"log"
	"os"
)

var L *log.Logger

func init() {
	L = log.New(os.Stdout, "GO-QL: ", log.Ldate|log.Ltime|log.Lshortfile)
}
