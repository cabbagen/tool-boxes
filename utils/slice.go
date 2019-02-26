package utils

import (
	"log"
)

// interface 
type uslicer interface {
	Length() int
}

// int interface
type uintslicer interface {
	Value() []int
}


// int slice struct
// ============================================
// ============================================
type UIntSlice struct {
	uslice     []int
}

func NewUIntSlice(oslice []int) UIntSlice {
	return UIntSlice {oslice}
}

func (uis UIntSlice) Length() int {
	return len(uis.uslice)
}

func (uis UIntSlice) Value() []int {
	return uis.uslice
}

func (uis UIntSlice) Of(ints ...int) UIntSlice {
	intslice := []int{}
	
	for _, value := range ints {
		intslice = append(intslice, value)
	}
	
	return UIntSlice {intslice}
}

func (uis UIntSlice) Contact(ints ...int) UIntSlice {
	for _, value := range ints {
		uis.uslice = append(uis.uslice, value)
	}
	
	return uis
}

func (uis UIntSlice) Every(fn func(value int) bool) bool {
	rslice := []bool{}
	
	for _, value := range uis.uslice {
		if result := fn(value); result {
			rslice = append(rslice, result)
		}
	}
	
	return len(rslice) == len(uis.uslice)
}

func (uis UIntSlice) Fill(value, start, end int) UIntSlice {
	ulen := len(uis.uslice)
	
	if start >= end {
		log.Fatal("start can not > end")
	}
	
	if start > ulen {
		log.Fatal("start can not > len(uis.uslice)")
	}
	
	rend := ulen
	
	if ulen > end {
		rend = end
	}

	for k, v := range uis.uslice {
		uis.uslice[k] = v
		
		if k >= start && k < rend {
			uis.uslice[k] = value
		}
	}
	
	return uis
}









