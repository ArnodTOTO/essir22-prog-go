package verifed

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
)

//function that checks the target
func Target(target string) bool{
	if target == "localhost"{
		return true
	}else{
		match, _ := regexp.MatchString("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$", target)
		if match {
			return true
		}else {
			return false
		}
	}
}

//function taht checks the port
func Verified_Port(parseport string) string {
	results := ""
	for {
		if parseport == "all" {
			return fmt.Sprintf("All")
			break
		}else {
			port, err := strconv.Atoi(parseport)
			if err != nil {
				counter := strings.Contains(parseport, "-")
				if counter == true {
					parts:= strings.Split(parseport, "-")
					startport, err := strconv.Atoi(parts[0])
					if err != nil {
						return fmt.Sprintf("Invalid port")
						break
					}
					endport, err := strconv.Atoi(parts[1])
					if err != nil {
						return fmt.Sprintf("Invalid port")
						break
					}
					if startport > endport {
						return fmt.Sprintf("Your port range %d-%d is backwards. Did you mean %d-%d?", startport, endport, endport, startport)
						break
					}
					if startport < 1 || startport > 65535 {
						return fmt.Sprintf("Invalid port")
						break
					}
					if endport < 1 || endport > 65535 {
						return fmt.Sprintf("Invalid port")
						break
					}
					return fmt.Sprintf("Multiple")

				}else{
					return fmt.Sprintf("Invalid port")
				}
				break
			}else{
				if port < 1 || port > 65535 {
					return fmt.Sprintf("Invalid port")
					break
				} else {
					return fmt.Sprintf("Simple")
					break
				}
			}
		}
	}
	return results
}

//function that checks the number of workers
func Workers(parseworkers string) bool{
	workers, err := strconv.Atoi(parseworkers)
	if err != nil {
		return false
	}
	if workers < 1 || workers > 20000 {
		return false
	}
	return true
}
