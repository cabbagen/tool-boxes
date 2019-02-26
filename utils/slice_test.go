package utils

import (
	"testing"
)

func Test_Int_Slice_Length(t *testing.T) {
	oslice := []int {12, 23, 23}
	
	uslice := NewUIntSlice(oslice)
	
	if uslice.Length() != 3 {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_Int_Slice_Value(t *testing.T) {	
	uslice := NewUIntSlice([]int {12, 23, 23}).Value()
	
	if uslice[0] != 12 {
		t.Error("no, failed !")
		return
	}
	
	if uslice[1] != 23 {
		t.Error("no, failed !")
		return
	}
	
	if uslice[2] != 23 {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_Int_Slice_Of(t *testing.T) {
	uslice := NewUIntSlice([]int {}).Of(12, 24).Value()

	if uslice[0] != 12 {
		t.Error("no, failed !")
		return
	}
	
	if uslice[1] != 24 {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_Int_Slice_Every(t *testing.T) {
	uslice := NewUIntSlice([]int {12, 23, 23})
	
	result := uslice.Every(func (value int) bool {
		return value > 15
	})
	
	if result {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}

func Test_Int_Slice_Fill(t *testing.T) {
	uslice := NewUIntSlice([]int {12, 23, 23}).Fill(0, 1, 10).Value()
	
	if uslice[0] != 12 {
		t.Error("no, failed !")
		return
	}
	
	if uslice[1] != 0 {
		t.Error("no, failed !")
		return
	}
	
	if uslice[2] != 0 {
		t.Error("no, failed !")
		return
	}
	
	t.Log("yeah, passed !")
}



































