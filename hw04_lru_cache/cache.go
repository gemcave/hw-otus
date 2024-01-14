package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	node, exist := cache.items[key]

	if exist {
		cache.queue.MoveToFront(node)
		node.Value.(*cacheItem).value = value
		return true
	}

	newCacheItem := &cacheItem{key, value}
	li := cache.queue.PushFront(newCacheItem)
	if cache.queue.Len() > cache.capacity {
		back := cache.queue.Back()
		backCacheItem := back.Value.(*cacheItem)

		cache.queue.Remove(back)
		delete(cache.items, backCacheItem.key)
	}

	cache.items[key] = li
	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	node, exist := cache.items[key]
	if exist {
		cache.queue.MoveToFront(node)
		return node.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
