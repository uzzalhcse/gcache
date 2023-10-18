package cache

type Cache interface {
	Set(key, value interface{}, ttl int64)
	Get(key interface{}) (interface{}, error)
}
