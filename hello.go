package main

import (
	"fmt"
	"log"
	"time"
)
func timeZone () (loctime time.Time) {
	loc, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		log.Panic(err)
	}
	t := time.Now().In(loc)
	return t
}
func main(){
	fmt.Println(timeZone())
}
