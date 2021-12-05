/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 15:48
 * @Motto: Keep thinking, keep coding!
 */

package bank

import "sync"

var rmu sync.RWMutex
var balance4 int

/**
RLock只能在临界区共享变量且在没有任何写入操作时可用，一般来说
我们不应该假设逻辑上的只读函数/方法，也不会去更新某一些变量。
比如一个方法功能是访问一个变量，但它也有可能会同时去给一个内部的计数器
或者去更新缓存--使即时的调用能够能够更快。

RWMutex只有当获得锁的大部分goroutine的都是读操作的时候，而锁在竞争条件下
即goroutine必须等待才能获取到锁的时候，RWMutex才是最能带来好处的，RWMutex需要更
复杂的内部记录，所以会让它比一般的无竞争锁的mutex慢一些
*/
func Balance4() int {
	rmu.RLock() // readers lock
	defer rmu.RUnlock()
	return balance4
}
