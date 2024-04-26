package interfaces_to_python

import (
	"context"
	"log"

	pb "air_driver/smart_air_conditioner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getClient() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	maxAttempts := 3
	log.Println("creating client connection...")
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		conn, err = grpc.Dial(
			"localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		log.Println("client connection created")
		if err == nil {
			return conn, nil
		}
		log.Printf("Attempt %d failed to connect: %v\n", attempt, err)
	}

	return nil, err
}
func GetTemperature() (error, float32) {
	conn, error := getClient()
	if error != nil {
		return error, -1
	}
	defer conn.Close()
	c := pb.NewTemperatureServiceClient(conn)

	// GetTemperature call
	temp, err := c.GetTemperature(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get temperature: %v", err)
	}
	return err, temp.Temperature
}
