package main

import (
	"chemi123/reward/proto/rewardapipb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := rewardapipb.NewRewardApiClient(conn)
	request := &rewardapipb.Request{
		RequestId: 100,
	}
	res, err := client.GetRewards(context.TODO(), request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
