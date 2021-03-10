// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, TZtime string) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func FlagIn() (string, int) {
	tz := os.Getenv("TZ")
	var port = flag.Int("port", 1234, "Port")
	flag.Parse()
	return tz, *port
}

func Time(tz string) string {
	t, err := TimeIn(time.Now(), tz)
	if err == nil {
		return fmt.Sprintf("%v\t: %v\n", t.Location(), t.Format("15:04:05"))
	} else {
		return fmt.Sprintf("%v\t: <timezone unknown>\n", t.Location())
	}
}

func main() {
	//Get variables
	tz, port := FlagIn()
	//get time
	TZtime := Time(tz)
	server := fmt.Sprintf("localhost:%v", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, TZtime) // handle connections concurrently
	}
}
