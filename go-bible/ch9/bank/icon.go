/**
 * @Author: caoduanxi
 * @Date: 2021/12/5 16:04
 * @Motto: Keep thinking, keep coding!
 */

package bank

import (
	"image"
	"sync"
)

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(name string) image.Image {
	return nil
}

// NOT concurrency-safe!
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

var iconmu sync.RWMutex // guards icons
func Icon2(name string) image.Image {
	iconmu.RUnlock()
	if icons != nil {
		icon := icons[name]
		iconmu.RUnlock()
		return icon
	}
	iconmu.RUnlock()

	// acquire an exclusive lock
	iconmu.Lock()
	if icons == nil { // must recheck for nil
		loadIcons()
	}

	icon := icons[name]
	iconmu.Unlock()
	return icon
}

// sync.Once
var loadIconOnce sync.Once

// concurrency safe
func Icon3(name string) image.Image {
	loadIconOnce.Do(loadIcons) // 利用sync.Once来做，即单例模式
	return icons[name]
}
