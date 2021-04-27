package hash

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRing_Add(t *testing.T) {
	ring := New(WithReplicas(3))
	var nodes []string
	for i := 0; i < 10; i++ {
		nodes = append(nodes, "Node"+strconv.Itoa(i))
	}
	ring.Add(nodes...)

	t.Logf("VirtualKeys = %v, \nRelations = %v\n",
		ring.vKeys, ring.meta)
}

func TestRing_Get(t *testing.T) {
	ring := New(WithReplicas(3))
	var nodes []string
	for i := 0; i < 10; i++ {
		nodes = append(nodes, "Node"+strconv.Itoa(i))
	}
	ring.Add(nodes...)

	r := rand.NewSource(time.Now().Unix())
	for i := 0; i < 10; i++ {
		key := strconv.FormatInt(r.Int63(), 10)
		t.Logf("Node = %v\n", ring.Get(key))
	}
}
