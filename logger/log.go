package logger

import (
	"log"
)

//functiono taht displays errors
func Logging(log_target bool, log_port string, log_workers bool) {
	if log_target == false {
		log.Fatal("the target is invalid")
	} else if log_port == "Invalid port" {
		log.Fatal("Invalid port")
	} else if log_workers == false {
		log.Fatal("Invalid workers, It must be between 1-2000")
	} else {
		log.Fatal(log_port)
	}
}
