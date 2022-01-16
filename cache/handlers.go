package cache

import (
	"context"
	"strconv"
	"time"

	init_cache "github.com/Fall-Web-Course/HW3/cache/init"
	proto "github.com/Fall-Web-Course/HW3/proto"
)


var client = init_cache.ConnectToCache("localhost:50051")

func GetCacheClient() proto.CacherClient {
	return client
}

func SetKey(key string, value string) (*proto.KeyPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	int_key, _ := strconv.ParseInt(key, 10, 64)
	r, err := client.SetKey(ctx, &proto.KeyPair{Key: int_key, Value: value})
	return r, err
}

func GetKey(key string) (*proto.Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	int_key, _ := strconv.ParseInt(key, 10, 64)
	r, err := client.GetKey(ctx, &proto.Key{Key: int_key})
	return r, err
}