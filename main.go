package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro"

	pb "github.com/finalsatan/shiiip-consignment/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(
		micro.Name("shiiip.cli"),
	)
	service.Init()

	client := pb.NewShippingServiceClient("shiiip.consignment", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Create consignment error: %v", err)
	}
	log.Printf("Created: %v", r.Created)

	r, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Get all consignments error: %v", err)
	}

	formatConsignments, _ := json.MarshalIndent(&r.Consignments, "", "\t")

	log.Printf("All consignments: %v", string(formatConsignments))

}
