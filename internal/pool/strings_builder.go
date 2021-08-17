package pool

import (
	"strings"
	"sync"
)

var stringsBuilderPool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}

// GetStringsBuilder gives a reference to a fresh strings.Builder from the pool.
func GetStringsBuilder() *strings.Builder {
	return stringsBuilderPool.Get().(*strings.Builder)
}

// PutStringsBuilder resets and puts the given strings.Builder back to the pool.
func PutStringsBuilder(sb *strings.Builder) {
	sb.Reset()
	stringsBuilderPool.Put(sb)
}
