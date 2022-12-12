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
	CommonLogFormat = "%s - \"%s %s %s\" %d %d"
	TimeLogFormat   = "2022-12-08T13:20:00Z"
)

var (
	created = time.Now()
	delay   = time.Second

	interval time.Duration
)

var (
	infoLog    = InfoLog()
	warningLog = WarningLog()
	errorLog   = ErrorLog()
	debugLog   = DebugLog()
)

func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
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

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

func InfoLog() string {
	return color.GreenString("INFO: " + color.WhiteString("This is Info an log message "))
}

func WarningLog() string {
	return color.YellowString("WARNING: " + color.WhiteString("This is warning an log message "))
}

func ErrorLog() string {
	return color.RedString("ERROR: " + color.WhiteString("This is error an log message "))
}

func DebugLog() string {
	return color.BlueString("DEBUG: " + color.WhiteString("This is debug an log message "))
}

func GenerateMsg(arr []string) string {
	l := len(arr)
	log := NewCommonLogFormat(created)

	for {
		time.Sleep(1 * delay)
		fmt.Println(created.Format(TimeLogFormat) + " " + arr[rand.Intn(l)] + color.HiWhiteString(log))
	}
}

func main() {
	arr := []string{infoLog, warningLog, errorLog, debugLog}

	GenerateMsg(arr)
}
