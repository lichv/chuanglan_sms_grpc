package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"sms.lichv.com/pb"
	"time"
)

func main() {
	conn,err := grpc.Dial("localhost:50001",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v",err)
	}
	defer conn.Close()

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	c := pb.NewSmsClient(conn)
	r, err := c.Send(ctx, &pb.ChuanglanRequest{Mobile: "15812345678", Content: "测试一个短信，使用grpc"})
	if err != nil {
		log.Fatalf("could not greet:%v",err)
	}
	log.Printf("Success:%s",r.Message)
}
