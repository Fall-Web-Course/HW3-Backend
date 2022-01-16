package init

import (
	proto "github.com/Fall-Web-Course/HW3/proto"
	grpc "google.golang.org/grpc"
)

func ConnectToCache(address string) proto.CacherClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := proto.NewCacherClient(conn)
	return client
}