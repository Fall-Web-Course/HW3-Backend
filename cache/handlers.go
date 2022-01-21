package cache

import (
	"context"
	"strconv"
	"time"

	proto "github.com/Fall-Web-Course/HW3/proto"
)

var client proto.CacherClient

func GetCacheClient() proto.CacherClient {
	return client
}

func SetCacheClient(cacher_client proto.CacherClient) {
	client = cacher_client
}

func SetKey(key string, value string) (*proto.KeyPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	int_key, _ := strconv.ParseInt(key, 10, 64)
	r, err := client.SetKey(ctx, &proto.KeyPair{Key: int_key, Value: value})
	return r, err
}

func GetKey(key string) (r *proto.Value, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	int_key, _ := strconv.ParseInt(key, 10, 64)
	r, err = client.GetKey(ctx, &proto.Key{Key: int_key})
	return r, err
}

func SetKeyWithDeadline(key string, value string, dead_line time.Time) (*proto.KeyPair, error) {
	ctx, cancel := context.WithDeadline(context.Background(), dead_line)
	defer cancel()
	int_key, _ := strconv.ParseInt(key, 10, 64)
	r, err := client.SetKey(ctx, &proto.KeyPair{Key: int_key, Value: value})
	return r, err
}
