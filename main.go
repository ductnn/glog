package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
)

const (
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	CommonLog       = "02/Jan/2006:15:04:05 -0700"
	TimeLogFormat   = "2022-12-08T13:20:00Z"
)

func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

func InfoLog() string {
	return color.GreenString("INFO: " + color.WhiteString("This is Info an log message"))
}

func WarningLog() string {
	return color.YellowString("WARNING: " + color.WhiteString("This is warning an log message"))
}

func ErrorLog() string {
	return color.RedString("ERROR: " + color.WhiteString("This is error an log message"))
}

func DebugLog() string {
	return color.BlueString("DEBUG: " + color.WhiteString("This is debug an log message"))
}

func main() {
	var (
		created = time.Now()
		delay   = time.Second

		interval time.Duration
	)

	for {
		time.Sleep(2 * delay)
		log := NewCommonLogFormat(created)
		// fmt.Println(log + "\n")
		color.HiWhite(log + "\n")
		time.Sleep(1 * delay)
		fmt.Printf(created.Format(TimeLogFormat) + " " + InfoLog() + "\n")
		time.Sleep(1 * delay)
		fmt.Printf(created.Format(TimeLogFormat) + " " + WarningLog() + "\n")
		time.Sleep(1 * delay)
		fmt.Printf(created.Format(TimeLogFormat) + " " + ErrorLog() + "\n")
		time.Sleep(1 * delay)
		fmt.Printf(created.Format(TimeLogFormat) + " " + DebugLog())
		fmt.Print("\n")
		created = created.Add(interval)
	}
}
