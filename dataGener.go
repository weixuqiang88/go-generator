package easy

import (
	_ "errors"
	_ "fmt"
	"math/rand"
	_ "strings"
	"time"
)

type DataSlice []string

type IterData struct {
	StrList DataSlice
	index   int
}

func NewIterData(str_arr DataSlice) *IterData {
	iter := new(IterData)
	iter.StrList = str_arr
	iter.index = 0
	return iter
}

func (iter *IterData) Next() string {
	var rs_p string
	for {
		if iter.index < len(iter.StrList) {
			rs_p = iter.StrList[iter.index]
			iter.index++
			return rs_p
		} else {
			iter.index = 1
			return iter.StrList[0]
		}
	}
}

type RateIter struct {
	IterData
	Rate      time.Duration
	StartTime time.Time
}

func NewRateIter(str_slice DataSlice, rate time.Duration) *RateIter {
	ri := new(RateIter)
	ri.StrList = str_slice
	ri.index = 0
	ri.Rate = rate
	ri.StartTime = time.Now()
	return ri
}

func (ri *RateIter) Next() string {
	var rs string
	for {
		var d time.Duration
		now := time.Now()
		d = now.Sub(ri.StartTime)
		if d >= ri.Rate {
			ri.StartTime = now
			ri.index++
		}
		if ri.index < len(ri.StrList) {
			rs = ri.StrList[ri.index]
			return rs
		} else {
			ri.index = 0
			rs = ri.StrList[0]
			return rs
		}
	}
}

type GenerateRandom struct {
	RandPtr *rand.Rand
}

func NewGenerateRandom() *GenerateRandom {
	gr := new(GenerateRandom)
	gr.RandPtr = rand.New(rand.NewSource(time.Now().UnixNano()))
	return gr
}

func (gr *GenerateRandom) GenerateInt(start int, end int) int {
	return gr.RandPtr.Intn((end - start)) + start
}

func (gr *GenerateRandom) GenerateSliceItem(sl DataSlice) string {
	return sl[gr.RandPtr.Intn(len(sl))]
}
