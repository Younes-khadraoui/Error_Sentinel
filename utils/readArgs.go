package utils

import (
	"log"
	"regexp"
	"strconv"
)

func GetPort(args []string) string {
	if len(args) < 2 {
		log.Panic("Please Provide a PORT number")
	}
	port := args[1]
	var re = regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(port) {
		log.Panic("PORT must contain only numbers.")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Panic("Error converting Port to INT ", err)
	}
	if portInt > 65353 {
		log.Panic("Plese Provide a PORT number < 65353")
	}

	return port
}
