package circularbuffer

import (
	"github.com/AMekss/assert"
	"testing"
)

func getFilledBuffer(t *testing.T) *CircularBuffer {
	v := New(5, 10)
	for i := 1; i < 12; i++ {
		v.Insert(float64(i))
	}
	return v
}

func TestInsert(t *testing.T) {
	v := getFilledBuffer(t)

	wantArray := []float64{11, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range wantArray {
		want := wantArray[i]
		if want != v.records[i] {
			t.Errorf("TestInsert: wantArray differs: index=%d want=%.1f got=%.1f", i, want, v.records[i])
		}
	}
}

func TestMedian(t *testing.T) {
	v := getFilledBuffer(t)

	perf, err := v.Median()
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, 6, perf)
}

func TestAverage(t *testing.T) {
	v := getFilledBuffer(t)

	perf, err := v.Average()
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, 6.5, perf)
}

func TestQuantile(t *testing.T) {
	v := getFilledBuffer(t)

	perf, err := v.Quantile(0)
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, v.sortedRecords[0], perf)

	perf, err = v.Quantile(1)
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, v.sortedRecords[len(v.sortedRecords)-1], perf)
}

func TestMin(t *testing.T) {
	v := getFilledBuffer(t)

	perf, err := v.Min()
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, 2, perf)
}

func TestMax(t *testing.T) {
	v := getFilledBuffer(t)

	perf, err := v.Max()
	assert.NoError(t.Fatalf, err)
	assert.EqualFloat64(t, 11, perf)
}
