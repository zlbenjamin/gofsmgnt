package utils

import (
	"encoding/json"
	"runtime"
)

// Memory infomation
type MemInfo struct {
	// Allocated but not released
	Alloc uint64 `json:"alloc"`
	// Total allocated since started
	TotalAlloc uint64 `json:"totalAlloc"`
	// Total allocated from OS
	Sys uint64 `json:"sys"`
	// Allocated for heap but not released
	HeapAlloc uint64 `json:"heapAlloc"`
	// Heap memory
	HeapSys uint64 `json:"heapSys"`
	// Times of GC
	NumGC uint32 `json:"numGc"`
}

// MemInfo: String()
func (m MemInfo) String() string {
	bts, err := json.Marshal(m)
	if err == nil {
		return string(bts)
	}

	return ""
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

// System information
type SystemInfo struct {
	// sys architecture
	Arch string `json:"arch"`
	// operation system
	OS string `json:"os"`
	// The number of logical CPUs usable by the current process
	NumCPU int `json:"numCpu"`
	// The number of goroutines that currently exist
	NumGoroutine int `json:"numGoroutine"`
}

// SystemInfo: String()
func (m SystemInfo) String() string {
	bts, err := json.Marshal(m)
	if err == nil {
		return string(bts)
	}

	return ""
}

// Get system information
func GetSystemInfo() (ret SystemInfo) {
	ret.Arch = runtime.GOARCH
	ret.OS = runtime.GOOS
	ret.NumCPU = runtime.NumCPU()
	ret.NumGoroutine = runtime.NumGoroutine()

	return
}
