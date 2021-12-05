/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 16:14
 * @Motto: Keep thinking, keep coding!
 */

package bank

import (
	"fmt"
	"sync"
	"testing"
)
// 这里会产生死锁，因为不可重入，如果是不同的锁，不会造成死锁
// 但这里的锁是全局的，所以会死锁
var tmu sync.RWMutex
func TestRace(t *testing.T) {
	tmu.Lock()
	fmt.Println("in main test for lock race")
	AddRace()
	tmu.Unlock()
}

func AddRace() {
	tmu.Lock()
	defer tmu.Unlock()
	fmt.Println("test for lock race")
}
