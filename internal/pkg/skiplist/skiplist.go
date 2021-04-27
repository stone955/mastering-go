package skiplist

const (
	defaultP              = 0.25
	defaultMaxNodes int64 = 65535
)

// List 跳表结构体
type List struct {
	// 跳表的最大level
	maxLevel int
	// 当前跳表所有节点的最大level
	currentMaxLevel int
	// 跳表的最大元素个数
	maxNodes int64
	// 表头
	headNode Node
	// 表尾
	tailNode Node
	// 生成随机level的概率值
	p float64
}

// New 创建一个跳表
func New(opts ...Option) *List {
	// TODO

	l := &List{}

	for _, opt := range opts {
		opt(l)
	}
	return l
}

// Option ...
type Option func(skipList *List)

// WithP ...
func WithP(p float64) Option {
	return func(skipList *List) {
		if p > 0 && p < 100 {
			skipList.p = p
		}
	}
}

// WithMaxNodes ...
func WithMaxNodes(maxNodes int64) Option {
	return func(skipList *List) {
		if maxNodes > 0 {
			// 计算层数
			skipList.maxNodes = maxNodes
		}
	}
}

// Node 跳表节点
type Node struct {
	// TODO ...
}
