package go_gcloud_engine

import (
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
)

type AppMainConfig struct {
	Memcache struct {
		Servers []string `json:"servers"` // ["10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212"]
	} `json:"memcache"`
	ListenAndServe string `json:"listen_and_serve"` // ":8080"
	QueueService   string `json:"queue_service"`
}

var MemcacheConnection *memcache.Client
var MemcachePing bool
var QueueService string

func AppEngineMain(config AppMainConfig) {
	MemcacheConnection = memcache.New(config.Memcache.Servers...)
	MemcachePing = MemcacheConnection.Ping() == nil
	QueueService = config.QueueService

	if config.ListenAndServe == "" {
		config.ListenAndServe = ":8080"
	}
	_ = http.ListenAndServe(config.ListenAndServe, nil)
}

func AddCache(mc *memcache.Client, Key string, Value []byte) error {
	return mc.Add(&memcache.Item{Key: Key, Value: Value})
}

func SetCache(mc *memcache.Client, Key string, Value []byte) error {
	return mc.Set(&memcache.Item{Key: Key, Value: Value})
}

func DeleteCache(mc *memcache.Client, Key string) error {
	return mc.Delete(Key)
}

func TouchCache(mc *memcache.Client, Key string, Seconds int) error {
	return mc.Touch(Key, int32(Seconds))
}

func GetCache(mc *memcache.Client, Key string) (checkCache bool, Value []byte) {
	it, err := mc.Get(Key)
	if err == nil && it != nil {
		checkCache = it.Key == Key
		Value = it.Value
	}
	return
}
