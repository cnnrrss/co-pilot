package copilot

import "encoding/json"

// Cache defines the interface methods for accessing the cached contacts
type Cache interface {
	GetContactByID(id string) ([]byte, error)
	SetContactByID(id string, contact string) error
	CleanCache() error
}

// GetContactByID ...
func (r *RedisCache) GetContactByID(id string) ([]byte, error) {
	response, err := r.Client.Get(id).Result()
	if err != nil {
		return nil, err
	}
	return []byte(response), nil
}

// SetContactByID updates the cache
func (r *RedisCache) SetContactByID(id string, contact string) error {
	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}
	return r.Client.Set(id, string(data), *r.ttl).Err()
}

// CleanCache to flush all k,v pairs from the cache
func (r *RedisCache) CleanCache() error {
	sts := r.FlushDBAsync()
	return sts.Err()
}
