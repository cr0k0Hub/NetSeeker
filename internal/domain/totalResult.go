package domain

import (
	"fmt"
	"sync/atomic"
)

type TotalResult struct {
	Scanned     atomic.Int64
	Opened      atomic.Int64
	Closed      atomic.Int64
	Timeout     atomic.Int64
	Unreachable atomic.Int64
	Errors      atomic.Int64
}

func InitTotalResult() TotalResult {
	return TotalResult{}
}

func (t *TotalResult) Print() {
	fmt.Printf("Scanned:     %d\n", t.Scanned.Load())
	fmt.Printf("Opened:      %d\n", t.Opened.Load())
	fmt.Printf("Closed:      %d\n", t.Closed.Load())
	fmt.Printf("Timeout:     %d\n", t.Timeout.Load())
	fmt.Printf("Unreachable: %d\n", t.Unreachable.Load())
	fmt.Printf("Errors:      %d\n", t.Errors.Load())
}
