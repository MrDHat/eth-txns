package logger

import (
	"fmt"
	"os"
)

func Log(message ...interface{}) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(message...)
	}
}
