package parser

import (
	"strconv"
)

func ParsePort(portStr string) (uint16, error) {
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return 0, &ParseError{
			Type:    ErrInvalidPort,
			Message: portStr,
		}
	}
	if port < 1 || port > 65535 {
		return 0, &ParseError{
			Type:    ErrInvalidPort,
			Message: portStr,
		}
	}
	return uint16(port), nil
}
