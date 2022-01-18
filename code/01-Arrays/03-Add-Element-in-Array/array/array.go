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

// Add 在第 index 个位置插入一个新元素 e
func (a *Array) Add(index, e int) {

	if a.size == len(a.data) {
		panic("add faild, array is full")
	}

	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}

	// 将index之后的元素都向后移一位，从末尾开始移动
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	// 添加元素
	a.data[index] = e
	a.size++
}

// AddLast 在末尾追加元素
func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

// AddFirst 在首部追加元素
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}
