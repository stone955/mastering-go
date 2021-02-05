package hash

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// Hash 抽象的哈希函数
type Hash func(data []byte) uint32

// Ring 抽象的哈希环
type Ring struct {
	mu       sync.RWMutex
	h        Hash           // 计算哈希的函数
	replicas int            // 副本数，影响虚拟节点的个数
	vKeys    []int          // 保存虚拟节点的有序列表，从小到大排序
	meta     map[int]string // 虚拟节点与物理节点对应关系，如果你有3个节点，replicas 设置成 3，那么就在环上有9个节点，9个元素
}

// Option 配置函数
type Option func(ring *Ring)

// WithReplicas 配置副本数
func WithReplicas(replicas int) Option {
	return func(ring *Ring) {
		ring.replicas = replicas
	}
}

// WithHash 配置哈希函数
func WithHash(h Hash) Option {
	return func(ring *Ring) {
		ring.h = h
	}
}

// New 初始化函数
func New(options ...Option) *Ring {
	r := Ring{
		h:        crc32.ChecksumIEEE,
		replicas: 3,
		vKeys:    []int{},
		meta:     make(map[int]string),
	}

	for _, option := range options {
		option(&r)
	}
	return &r
}

// Add 添加节点
func (r *Ring) Add(keys ...string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, key := range keys {
		for i := 0; i < r.replicas; i++ {
			// 计算每个虚拟节点的key
			vKey := int(r.h([]byte(strconv.Itoa(i) + key)))
			// 追加到哈希环的虚拟节点
			r.vKeys = append(r.vKeys, vKey)
			// 虚拟节点 <--> 物理节点
			r.meta[vKey] = key
		}
	}
	// 做完排序之后，其实就完成了值域的打散划分
	sort.Ints(r.vKeys)
}

// Get 返回key对应的物理节点
func (r *Ring) Get(key string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// 根据输入的key计算出哈希值
	hKey := int(r.h([]byte(key)))
	// 计算该哈希值落在哪个值域范围，选择到虚拟节点
	idx := sort.Search(len(r.vKeys), func(i int) bool {
		return r.vKeys[i] > hKey
	})
	if idx == len(r.vKeys) {
		idx = 0
	}
	// 选择到具体的物理节点
	return r.meta[r.vKeys[idx]]
}
