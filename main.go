package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
)

const (
	CommonLogFormat = "%s - \"%s %s %s\" %d %d"
	TimeLogFormat   = "2006-01-02T15:04:05-0700"
)

var (
	infoLog    = InfoLog()
	warningLog = WarningLog()
	errorLog   = ErrorLog()
	debugLog   = DebugLog()

	created = time.Now()
	delay   = time.Second

	interval time.Duration
)

func NewCommonLogFormat(t time.Time) string {
	return color.HiWhiteString(
		CommonLogFormat,
		// gofakeit.IPv4Address(),
		randIP(),
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
	return (color.HiGreenString("INFO:    "))
}

func WarningLog() string {
	return (color.HiYellowString("WARNING: "))
}

func ErrorLog() string {
	return (color.HiRedString("ERROR:   "))
}

func DebugLog() string {
	return (color.HiBlueString("DEBUG:   "))
}

func randIP() string {
	ips := []string{}
	min := 1
	max := 255

	for j := 0; j < 4; j++ {
		random := strconv.Itoa(rand.Intn(max-min) + min)
		ips = append(ips, random)
	}

	return strings.Join(ips, ".")
}

func GenerateMsg(arr []string) string {
	l := len(arr)

	return (color.HiWhiteString(created.Format(TimeLogFormat)) + " " + arr[rand.Intn(l)])
}

func main() {
	arr := []string{infoLog, warningLog, errorLog, debugLog}
	log := NewCommonLogFormat(created)

	for {
		time.Sleep(1 * delay)
		fmt.Println(GenerateMsg(arr) + " " + log)
		created = created.Add(interval)
	}
}
