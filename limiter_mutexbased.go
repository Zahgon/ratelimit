// Copyright (c) 2016,2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ratelimit // import "go.uber.org/ratelimit"

import (
	"sync"
	"time"
)

type mutexLimiter struct {
	sync.Mutex
	last       time.Time
	sleepFor   time.Duration
	perRequest time.Duration
	maxSlack   time.Duration
	clock      Clock
}

// newMutexBased returns a new mutex based limiter.
func newMutexBased(rate int, opts ...Option) *mutexLimiter {
	_ = "STUB: not implemented"
	// TODO consider moving config building to the implementation
	// independent code.
	return nil
}

// Take blocks to ensure that the time spent between multiple
// Take calls is on average per/rate.
func (t *mutexLimiter) Take() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// If this is our first request, then we allow it.

// sleepFor calculates how much time we should sleep based on
// the perRequest budget and how long the last request took.
// Since the request may take longer than the budget, this number
// can get negative, and is summed across requests.

// We shouldn't allow sleepFor to get too negative, since it would mean that
// a service that slowed down a lot for a short period of time would get
// a much higher RPS following that.

// If sleepFor is positive, then we should sleep now.
