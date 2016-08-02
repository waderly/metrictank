package offsetfence

// OffsetFence allows to put restrictions on offset partitions for given partitions
// when using kafka for metrics and clustering, we use it to assure the metrics consumption is
// sufficiently in sync, before processing cluster messages
type OffsetFence struct {
	sync.Mutex
	data map[int]int64
}

func New() OffsetFence {
	return OffsetFence{
		sync.Mutex{},
		make(map[int]int64),
	}

}

func (o *OffsetFence) Set(part int, offset int64) {
	o.Lock()
	o.data[part] = offset
	o.Unlock()
}
