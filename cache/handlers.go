package cache

import (
	init_cache "github.com/Fall-Web-Course/HW3/cache/init"
	proto "github.com/Fall-Web-Course/HW3/proto"
)


var client = init_cache.ConnectToCache("localhost:50051")

func GetCacheClient() proto.CacherClient {
	return client
}