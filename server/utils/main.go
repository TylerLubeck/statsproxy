package utils

import (
	"fmt"
)

func FormatServer(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
