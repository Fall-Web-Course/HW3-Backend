package init

import (
	"fmt"

	cache "github.com/Fall-Web-Course/HW3/cache"
	proto "github.com/Fall-Web-Course/HW3/proto"
	utils "github.com/Fall-Web-Course/HW3/utils"

	"google.golang.org/grpc"
)

func InitCache() {
	CACHE_ADDRESS := utils.Getenv("CACHE_ADDRESS", "127.0.0.1")
	CACHE_PORT := utils.Getenv("CACHE_PORT", "50051")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", CACHE_ADDRESS, CACHE_PORT), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	client := proto.NewCacherClient(conn)
	cache.SetCacheClient(client)
}