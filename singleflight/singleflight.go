package singleflight

import "sync"

// call 一次请求的抽象
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group 是 singleflight 的主数据结构，管理不同 key 的请求
type Group struct {
	mu sync.Mutex // protects m, 确保
	m  map[string]*call
}

// Do 保证相同的 key 在一次 Call 中只会执行一次 fn，这里像是一个简单的工具封装
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil { // 懒汉式初始化
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok { // 如果已经有人在请求了
		g.mu.Unlock() // 这里是为了保证在调用 fn 之前，不会有其他的请求进来
		c.wg.Wait()   // 等别人 fn 请求的结果
		return c.val, c.err
	}

	// 存储调用状态
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	// 删除调用状态
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
