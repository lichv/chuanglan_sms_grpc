package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sms.lichv.com/pb"
	"sms.lichv.com/utils"
)

type server struct{
	*pb.UnimplementedSmsServer
}

func (s *server) Send(ctx context.Context, in *pb.ChuanglanRequest) (*pb.ChuanglanReply,error) {
	res, err := utils.CreateChuanglan("xxxx", "XXXXXX", "XXXX").Send(in.Mobile, in.Content)
	if err != nil {
		return nil,err
	}
	return &pb.ChuanglanReply{Message: *res},nil
}

func main(){
	lis,err := net.Listen("tcp",":50001")
	if err != nil {
		log.Fatalf("failed to listen:%v",err)
	}
	s := grpc.NewServer()
	pb.RegisterSmsServer(s,&server{})
	reflection.Register(s)
	log.Println("短信服务开启")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v",err)
	}
}

