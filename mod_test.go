package punit

import (
	"testing"
	"time"
	"math/rand"
	"sort"
)

const (
	samplec = 1000000
	testc = 10
)

func TestFmtConcat(t *testing.T) {
	samples := make([]int64, samplec)
	for i := 0; i < samplec; i++ {
		samples[i] = int64(rand.Int()) % int64(time.Hour*24*365)
	}
	sort.Slice(samples, func(i, j int) bool {
		return samples[i] < samples[j]
	})
	for i := 0; i < testc; i++ {
		x := samples[i * samplec / testc]
		t.Logf("%s\n", TimeScale.Format(int64(x), int64(time.Second), true))
	}
}

func TestFmtPlain(t *testing.T) {
	samples := make([]int64, samplec)
	for i := 0; i < samplec; i++ {
		samples[i] = int64(rand.Int()) % int64(time.Hour*24*365)
	}
	sort.Slice(samples, func(i, j int) bool {
		return samples[i] < samples[j]
	})
	for i := 0; i < testc; i++ {
		x := samples[i * samplec / testc]
		t.Logf("%s\n", TimeScale.Format(int64(x), int64(time.Second), false))
	}
}
