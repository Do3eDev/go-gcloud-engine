package go_gcloud_engine

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
	"time"
)

func MemCacheDelete(Env string, request *http.Request, key string) bool {
	if Env == "local" {
		return false
	}
	err := memcache.Delete(appengine.NewContext(request), key)
	if err != nil {
		return false
	}
	return true
}

func MemCacheAdd(Env string, request *http.Request, key string, value []byte, minute time.Duration) bool {
	if Env == "local" {
		return false
	}
	var err error
	if minute == 0 {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value})
	} else {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value, Expiration: time.Minute * minute})
	}
	if err == memcache.ErrNotStored {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func MemCacheAddRandom(Env string, request *http.Request, key string, value []byte, minute time.Duration) bool {
	if Env == "local" {
		return false
	}
	var err error
	if minute == 0 {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value})
	} else {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value, Expiration: minute})
	}
	if err == memcache.ErrNotStored {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func MemCacheAddSecond(Env string, request *http.Request, key string, value []byte, second time.Duration) bool {
	if Env == "local" {
		return false
	}
	var err error
	if second == 0 {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value})
	} else {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value, Expiration: time.Second * second})
	}
	if err == memcache.ErrNotStored {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func MemCacheSet(Env string, request *http.Request, key string, value []byte, minute time.Duration) bool {
	if Env == "local" {
		return false
	}
	var err error
	if minute == 0 {
		err = memcache.Set(appengine.NewContext(request), &memcache.Item{Key: key, Value: value})
	} else {
		err = memcache.Set(appengine.NewContext(request), &memcache.Item{Key: key, Value: value, Expiration: time.Minute * minute})
	}
	if err == memcache.ErrNotStored {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func MemCacheSetSecond(Env string, request *http.Request, key string, value []byte, second time.Duration) bool {
	if Env == "local" {
		return false
	}
	var err error
	if second == 0 {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value})
	} else {
		err = memcache.Add(appengine.NewContext(request), &memcache.Item{Key: key, Value: value, Expiration: time.Second * second})
	}
	if err == memcache.ErrNotStored {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func MemCacheGet(Env string, request *http.Request, key string) ([]byte, bool) {
	if Env == "local" {
		return nil, false
	}
	Item, err := memcache.Get(appengine.NewContext(request), key)
	if err != nil {
		return nil, false
	}
	return Item.Value, true
}
