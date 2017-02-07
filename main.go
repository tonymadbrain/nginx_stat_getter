package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var (
	stat = flag.String("s", "", "Some stat from range: active, accepted, handled, requests, reading, writing, waiting")
	host = flag.String("h", "", "Host of Nginx server")
	port = flag.Int("p", 80, "Port of Nginx server")
)

func main() {
	flag.Parse()

	if flag.NFlag() < 3 {
		printHelp()
		os.Exit(1)
	}

	url := fmt.Sprintf("http://%s:%d/nginx_status/", *host, *port)

	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s", string(bytes))
		// Format:
		// Active connections: 2
		// server accepts handled requests
		// 246 246 950
		// Reading: 0 Writing: 1 Waiting: 1

		re := regexp.MustCompile("\\d+")
		res_array := re.FindAllString(string(bytes), -1)

		switch *stat {
		case "active":
			fmt.Println(res_array[0])
		case "accepted":
			fmt.Println(res_array[1])
		case "handled":
			fmt.Println(res_array[2])
		case "requests":
			fmt.Println(res_array[3])
		case "reading":
			fmt.Println(res_array[4])
		case "writing":
			fmt.Println(res_array[5])
		case "waiting":
			fmt.Println(res_array[6])
		case "pure":
			fmt.Println(string(bytes))
		case "array":
			fmt.Println(res_array)
		default:
			log.Fatal("Unrecognized stat requested")
		}
	}
}

func printHelp() {
	fmt.Println("Need arguments:\n  h - Host of Nginx server\n  p - Port of Nginx server\n  s - Some stat from range: active, accepted, handled, requests, reading, writing, waiting")
}
