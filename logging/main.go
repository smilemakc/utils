package logging

import (
	"os"
	"strconv"
	"log"
)

var (
	debugEnv    = os.Getenv("DEBUG")
	_debug bool

)
func init() {
	if debugEnv == "" {
		_debug = false
	} else {
		d, err := strconv.ParseBool(debugEnv)
		if err != nil {
			log.Println("error parse DEBUG", err)
			_debug = false
		} else {
			_debug = d
		}
	}

	logInit()
}