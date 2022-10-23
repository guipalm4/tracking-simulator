package main

import (
	"fmt"
	"log"
  "github.com/guipalm4/tracking-simulator/simulator/infra/kafka"
	//"github.com/guipalm4/simulator-go/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[1])

	
}
