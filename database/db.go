package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("error can't load PORT fromm env")
		port = 27010

	}
	return fmt.Sprintf("")
}
