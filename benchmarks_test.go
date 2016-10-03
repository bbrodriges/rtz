/*
Copyright (c) 2016 bender.rodriges

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package rtz

import (
	"math/rand"
	"testing"
	"time"
)

// Big square in YEKT time zone
const (
	minLat  = 54
	minLong = 79
	maxLat  = 70
	maxLong = 124

	loops = 100000
)

var coordinates map[int][]float64

func genCoords(n int) {
	coordinates = make(map[int][]float64, n)

	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		lat := float64(rand.Intn(maxLat-minLat)+minLat) + rand.Float64()
		long := float64(rand.Intn(maxLong-minLong)+minLong) + rand.Float64()
		coordinates[i] = []float64{lat, long}
	}
}

// BenchmarkLookupLatLong measures impact of rtz
func BenchmarkLookupLatLong(b *testing.B) {
	genCoords(b.N)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		pair := coordinates[i]
		LookupZoneName(pair[0], pair[1])
	}
	b.StopTimer()
}
