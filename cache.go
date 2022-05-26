package cache

import (
	"time"
)

type Cache struct {
	note 		map[string]data
}
type data struct{
	value 		string
	deadline  	time.Time
}


func NewCache() Cache {
	return  Cache{
		note: map[string]data{}}
}


func (cache *Cache) Get(key string) (string, bool) {
	if _, ok := cache.note[key]; !ok {
		return "", false
	}
	if !cache.note[key].deadline.IsZero() && cache.note[key].deadline.Before(time.Now()) {
		return "", false
	}
	return cache.note[key].value, true


}


func (cache *Cache) Put(key, value string) {
	cache.note[key] = data{value: value}
}

func (cache *Cache) Keys() []string {
	var sliceOfKeys []string
	for k,v:=range cache.note{
		if !v.deadline.IsZero() && v.deadline.Before(time.Now()){
			continue
		}
		sliceOfKeys=append(sliceOfKeys, k)
	}
	return sliceOfKeys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.note[key] = data{value, deadline}
}







