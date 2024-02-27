package geecache

// A ByteView holds an immutable view of bytes.
// 字节切片的简单封装
type ByteView struct {
	b []byte
}

// Len returns the view's length
// 实现了 Value 接口，作为主要存储类型，支持任意类型的存储
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes
func cloneBytes(b []byte) []byte {
	// 只读的关键实现，切片数组默认是返回引用的
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
