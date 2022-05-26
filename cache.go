package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	note 		map[string]data
}
type data struct{
	value 		string
	deadline  	*time.Time
}



func NewCache() Cache {
	return  Cache{
		note: map[string]data{}}
}


func (cache *Cache) Get(key string) (string, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if _, ok := cache.note[key]; !ok {
		return "", false
	}
	if cache.note[key].deadline!=nil && cache.note[key].deadline.Before(time.Now()) {
		return "", false
	}
	return cache.note[key].value, true


}


func (cache *Cache) Put(key, value string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.note[key] = data{value,nil}
}

func (cache *Cache) Keys() []string {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var sliceOfKeys []string
	for k,v:=range cache.note{
		if v.deadline!=nil && v.deadline.Before(time.Now()){
			continue
		}
		sliceOfKeys=append(sliceOfKeys, k)
	}
	return sliceOfKeys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.note[key] = data{value, &deadline}
}







