package go_gcloud_engine

import (
	"net/http"
)

func MemCacheDelete(_ string, _ *http.Request, key string) bool {
	if !MemcachePing {
		return false
	}
	return DeleteCache(MemcacheConnection, key) == nil
}

func MemCacheAdd(
	_ string,
	_ *http.Request,
	key string,
	value []byte,
	minute int,
) bool {
	if !MemcachePing {
		return false
	}
	if AddCache(MemcacheConnection, key, value) == nil {
		if minute > 0 {
			return TouchCache(MemcacheConnection, key, minute*60) == nil
		}
	}
	return true
}

func MemCacheAddRandom(
	Env string,
	request *http.Request,
	key string,
	value []byte,
	minute int,
) bool {
	return MemCacheAdd(Env, request, key, value, minute)
}

func MemCacheAddSecond(
	_ string,
	_ *http.Request,
	key string,
	value []byte,
	second int,
) bool {
	if !MemcachePing {
		return false
	}
	if AddCache(MemcacheConnection, key, value) == nil {
		if second > 0 {
			return TouchCache(MemcacheConnection, key, second) == nil
		}
	}
	return true
}

func MemCacheSet(
	_ string,
	_ *http.Request,
	key string,
	value []byte,
	minute int,
) bool {
	if !MemcachePing {
		return false
	}
	if SetCache(MemcacheConnection, key, value) == nil {
		if minute > 0 {
			return TouchCache(MemcacheConnection, key, minute*60) == nil
		}
	}
	return true
}

func MemCacheSetSecond(
	_ string,
	_ *http.Request,
	key string,
	value []byte,
	second int,
) bool {
	if !MemcachePing {
		return false
	}
	if SetCache(MemcacheConnection, key, value) == nil {
		if second > 0 {
			return TouchCache(MemcacheConnection, key, second) == nil
		}
	}
	return true
}

func MemCacheGet(_ string, _ *http.Request, key string) (value []byte, cache bool) {
	if !MemcachePing {
		return nil, false
	}
	cache, value = GetCache(MemcacheConnection, key)
	return
}
