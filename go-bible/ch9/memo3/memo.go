/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 16:35
 * @Motto: Keep thinking, keep coding!
 */

package memo2

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Get is concurrency-safe
// 锁粒度太大了导致并发也和没有并发一样
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

