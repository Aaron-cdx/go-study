/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 17:03
 * @Motto: Keep thinking, keep coding!
 */

package memo5

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// this is the first request for this key
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

// 拿到值的话关闭e.ready
func (e *entry) call(f Func, key string) {
	// Evaluate the function
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition
	close(e.ready)
}

// 拿不到值的话会一直阻塞在这里，直到获取到值为止
func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition
	<-e.ready
	// Send the result to the client
	response <- e.res
}
