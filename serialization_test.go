package roaring

// to run just these tests: go test -run TestSerialization*

import (
	"testing"
	"bytes"
)

func TestSerializationBasic(t *testing.T) {
	rb := BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	buf := new(bytes.Buffer)
	_,err:=rb.WriteTo(buf)
	if err != nil {
		t.Errorf("Failed writing")
	}
	newrb:= NewRoaringBitmap()
	_,err=newrb.ReadFrom(buf)
	if err != nil {
		t.Errorf("Failed reading")
	}
	if ! rb.Equals(newrb) {
		t.Errorf("Cannot retrieve serialized version")
	}
}



func TestSerializationBasic2(t *testing.T) {
	rb := BitmapOf(1, 2, 3, 4, 5, 100, 1000,10000,100000,1000000)
	buf := new(bytes.Buffer)
	_,err:=rb.WriteTo(buf)
	if err != nil {
		t.Errorf("Failed writing")
	}
	newrb:= NewRoaringBitmap()
	_,err=newrb.ReadFrom(buf)
	if err != nil {
		t.Errorf("Failed reading")
	}
	if ! rb.Equals(newrb) {
		t.Errorf("Cannot retrieve serialized version")
	}
}

func TestSerializationBasic3(t *testing.T) {
	rb := BitmapOf(1, 2, 3, 4, 5, 100, 1000,10000,100000,1000000)
	for i:=5000000 ; i< 5000000+2*(1<<16); i++ {
		rb.Add(i)
	}
	buf := new(bytes.Buffer)
	_,err:=rb.WriteTo(buf)
	if err != nil {
		t.Errorf("Failed writing")
	}
	newrb:= NewRoaringBitmap()
	_,err=newrb.ReadFrom(buf)
	if err != nil {
		t.Errorf("Failed reading")
	}
	if ! rb.Equals(newrb) {
		t.Errorf("Cannot retrieve serialized version")
	}
}