/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
	"github.com/MyProject/verifed"
	"github.com/sirupsen/logrus"
	"time"
	"log"
	"github.com/MyProject/logger"
	"sync"
)


// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "allows to scan port",
	Long: `allows to scan the ports target`,
	Run: func(cmd *cobra.Command, args []string) {

		// retrieves the values entered by the user

		parsetarget, _ := cmd.Flags().GetString("target")
		parseport, _ := cmd.Flags().GetString("port")
		parseworkers, _ := cmd.Flags().GetString("workers")
		parsequiet, _ := cmd.Flags().GetBool("quiet")

		// verification of the argument
		if parsetarget == "" {
			log.Fatal("Target undefinited, use scan --help")
		}

		if parseport == "" {
			log.Fatal("Port undefinited, use scan --help")
		}
		results_port := verifed.Verified_Port(parseport)
		results_target := verifed.Target(parsetarget)
		results_workers := verifed.Workers(parseworkers)
		currenTime := time.Now()

		// If the port, the target and workers is good. we scan all the ports
		if results_target == true && results_port == "All" && results_workers == true {
			fmt.Printf("Starting scan %s at %s\n", parsetarget, currenTime.String())
			results := make(chan int)  //Create chan
			var wg sync.WaitGroup	   //Create waitgroup
			numThreads, _ := strconv.Atoi(parseworkers)

			//We calculate the number of ports to scan for each goroutine
			for i := 0; i < numThreads; i++ {
				start := (65535 / numThreads * i) + 1
				end := start + (65535 / numThreads) - 1
				wg.Add(1)
				go scanAll(start, end, &wg, results, parsetarget, parsequiet)
			}

			// function taht waits for the results
			go func() {
				wg.Wait()
				close(results)
			}()

			// show results
			for port := range results {
				fmt.Printf("Open ports: %d\n", port)
			}
		} else if results_target == true && results_port == "Simple" {
			// runs a port scan
                        fmt.Printf("Starting scan %s at %s\n", parsetarget, currenTime.String())
			scanSimple(parsetarget, parseport, parsequiet)
		} else if results_target == true && results_port == "Multiple" && results_workers == true {
			// runs a port range scan
			fmt.Printf("Starting scan %s at %s\n", parsetarget, currenTime.String())
			parts := strings.Split(parseport, "-")
			first, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			var wg sync.WaitGroup
			results := make(chan int)
			numThreads, _ := strconv.Atoi(parseworkers)
			portsPerThread := (end - first + 1) / numThreads
			for i := 0; i < numThreads; i++ {
				start := first + portsPerThread*i
				last := start + portsPerThread - 1
				check := numThreads - 1
				if i == check {
					last = end
				}
				wg.Add(1)
				go scanMultiple(start, last, results, &wg, parsetarget, parsequiet)
			}
			go func() {
				wg.Wait()
				close(results)
			}()
			for port := range results {
				fmt.Printf("Open port: %d\n", port)
			}
		}else{
			//displays the logs
			logger.Logging(results_target, results_port, results_workers)
		}

	},
}


//function that performs a single port scan
func scanSimple(ip string, port string, quiet bool) {

	address := ip + ":" + port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		if quiet == true {
			return
		}else {
			logrus.Info("Maybe close")
			return
		}
	}
	conn.Close()
	fmt.Printf("Port: %s Open\n", port)
}

//function that scans all ports
func scanAll(start, end int, wg *sync.WaitGroup, results chan int, ip string, quiet bool) {
	defer wg.Done()
	for i:= start; i <= end; i++ {
		address := ip + ":" + strconv.Itoa(i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			if quiet == true {
				continue
			} else {
				logrus.Info("Maybe close")
				continue
			}
		}
		conn.Close()
		results <- i
	}
}

//function that scans a ports range
func scanMultiple (startport, endport int, results chan int, wg *sync.WaitGroup, ip string, quiet bool) {
	defer wg.Done()
        for i := startport; i <= endport; i++ {
                address := ip + ":" + strconv.Itoa(i)
                conn, err := net.Dial("tcp", address)
                if err != nil {
			if quiet == true {
				continue
			} else {
				logrus.Info("Maybe close")
                        	continue
			}
                }
                conn.Close()
		results <- i
        }
}

// ours flags
func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.PersistentFlags().StringP("target", "t", "", "design the target")
	scanCmd.PersistentFlags().StringP("port", "p", "", "design the port target")
	scanCmd.PersistentFlags().StringP("workers", "w", "10", "indicates the number of workers")
	scanCmd.Flags().BoolP("quiet", "q", false, "do not log, only display the results")
}
