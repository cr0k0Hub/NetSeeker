package domain

import (
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
