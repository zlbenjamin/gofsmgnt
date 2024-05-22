package utils

import "runtime"

// Memory infomation
type MemInfo struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	HeapAlloc  uint64
	HeapSys    uint64
	NumGC      uint32
}

// Runtime memory info
func RuntimeMem() (ret MemInfo) {
	var rm runtime.MemStats
	runtime.ReadMemStats(&rm)

	// set ret
	ret.Alloc = rm.Alloc
	ret.TotalAlloc = rm.TotalAlloc
	ret.Sys = rm.Sys
	ret.HeapAlloc = rm.HeapAlloc
	ret.HeapSys = rm.HeapSys
	ret.NumGC = rm.NumGC

	return
}
