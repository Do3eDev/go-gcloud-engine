package go_gcloud_engine

import (
	"net/http"
	"time"
)

func MemCacheDelete(Env string, request *http.Request, key string) bool {
	if !MemcachePing {
		return false
	}
	return DeleteCache(MemcacheConnection, key) == nil
}

func MemCacheAdd(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	minute time.Duration,
) bool {
	if !MemcachePing {
		return false
	}
	if AddCache(MemcacheConnection, key, value) == nil {
		if minute > 0 {
			return TouchCache(MemcacheConnection, key, int(minute*60)) == nil
		}
	}
	return true
}

func MemCacheAddRandom(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	minute time.Duration,
) bool {
	return MemCacheAdd(Env, request, key, value, minute)
}

func MemCacheAddSecond(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	second time.Duration,
) bool {
	if !MemcachePing {
		return false
	}
	if AddCache(MemcacheConnection, key, value) == nil {
		if second > 0 {
			return TouchCache(MemcacheConnection, key, int(second)) == nil
		}
	}
	return true
}

func MemCacheSet(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	minute time.Duration,
) bool {
	if !MemcachePing {
		return false
	}
	if SetCache(MemcacheConnection, key, value) == nil {
		if minute > 0 {
			return TouchCache(MemcacheConnection, key, int(minute*60)) == nil
		}
	}
	return true
}

func MemCacheSetSecond(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	second time.Duration,
) bool {
	if !MemcachePing {
		return false
	}
	if SetCache(MemcacheConnection, key, value) == nil {
		if second > 0 {
			return TouchCache(MemcacheConnection, key, int(second)) == nil
		}
	}
	return true
}

func MemCacheGet(Env string, request *http.Request, key string) (value []byte, cache bool) {
	if !MemcachePing {
		return nil, false
	}
	cache, value = GetCache(MemcacheConnection, key)
	return
}
