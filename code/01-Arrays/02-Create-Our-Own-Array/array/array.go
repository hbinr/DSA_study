package array

type Array struct {
	data []int
	size int
}

// NewArray 构造函数，传入数组的capacity创建数组
func NewArray(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
		size: 0,
	}
}

// NewDefaultArray 创建一个默认容量为10的数组
func NewDefaultArray() *Array {
	return &Array{
		data: make([]int, 10),
		size: 0,
	}
}

// GetSize 获取数组大小
func (a *Array) GetSize() int {
	return a.size
}

// GetCapacity 获取数组容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// IsEmpty 判断数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
