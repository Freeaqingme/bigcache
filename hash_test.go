package bigcache

type hashStub uint64

func (stub hashStub) sum(_ string) uint64 {
	return uint64(stub)
}
