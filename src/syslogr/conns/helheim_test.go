// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package conns_test

import "time"

type mockReader struct {
	ReadCalled chan bool
	ReadInput  struct {
		Buf chan []byte
	}
	ReadOutput struct {
		Len chan int
		Err chan error
	}
}

func newMockReader() *mockReader {
	m := &mockReader{}
	m.ReadCalled = make(chan bool, 100)
	m.ReadInput.Buf = make(chan []byte, 100)
	m.ReadOutput.Len = make(chan int, 100)
	m.ReadOutput.Err = make(chan error, 100)
	return m
}
func (m *mockReader) Read(buf []byte) (len int, err error) {
	m.ReadCalled <- true
	m.ReadInput.Buf <- buf
	return <-m.ReadOutput.Len, <-m.ReadOutput.Err
}

type mockRanger struct {
	DelayRangeCalled chan bool
	DelayRangeOutput struct {
		Min, Max chan time.Duration
	}
}

func newMockRanger() *mockRanger {
	m := &mockRanger{}
	m.DelayRangeCalled = make(chan bool, 100)
	m.DelayRangeOutput.Min = make(chan time.Duration, 100)
	m.DelayRangeOutput.Max = make(chan time.Duration, 100)
	return m
}
func (m *mockRanger) DelayRange() (min, max time.Duration) {
	m.DelayRangeCalled <- true
	return <-m.DelayRangeOutput.Min, <-m.DelayRangeOutput.Max
}

type mockMetricBatcher struct {
	BatchAddCounterCalled chan bool
	BatchAddCounterInput  struct {
		Name  chan string
		Delta chan uint64
	}
}

func newMockMetricBatcher() *mockMetricBatcher {
	m := &mockMetricBatcher{}
	m.BatchAddCounterCalled = make(chan bool, 100)
	m.BatchAddCounterInput.Name = make(chan string, 100)
	m.BatchAddCounterInput.Delta = make(chan uint64, 100)
	return m
}
func (m *mockMetricBatcher) BatchAddCounter(name string, delta uint64) {
	m.BatchAddCounterCalled <- true
	m.BatchAddCounterInput.Name <- name
	m.BatchAddCounterInput.Delta <- delta
}