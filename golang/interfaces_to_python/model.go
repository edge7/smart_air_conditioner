package interfaces_to_python

import (
	"context"

	pb "air_driver/smart_air_conditioner"
)

func GetModelResult(inputPath string) (error, string) {
	conn, error := getClient()
	if error != nil {
		return error, ""
	}
	defer conn.Close()
	c := pb.NewTemperatureServiceClient(conn)

	modelResult, err := c.GetModelResult(
		context.Background(), &pb.ModelRequest{Input: inputPath},
	)
	return err, modelResult.GetResult()
}
