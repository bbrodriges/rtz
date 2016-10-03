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

import "testing"

func TestUnknownLocation(t *testing.T) {
	lat, long := 59.921212, 10.769582 // Oslo

	_, err := LookupZoneName(lat, long)

	if err == nil {
		t.Errorf("Oslo is obviously not in Russia")
	}
}

func TestLookupLatLong(t *testing.T) {
	cases := []struct {
		lat, long float64
		want      string
	}{
		{54.707323, 20.475165, usz1TZ},
		{55.751387, 37.598896, mskTZ},
		{53.195238, 50.157211, samtTZ},
		{56.828723, 60.600648, yektTZ},
		{54.984014, 73.389416, omstTZ},
		{56.029345, 92.874770, kratTZ},
		{52.284193, 104.278421, irktTZ},
		{62.025044, 129.729086, yaktTZ},
		{43.112873, 131.878365, vlatTZ},
		{59.567632, 150.808302, magtTZ},
		{53.037063, 158.654980, pettTZ},
	}
	for _, tt := range cases {
		got, err := LookupZoneName(tt.lat, tt.long)
		if err != nil {
			t.Error(err)
		}

		if got.Name != tt.want {
			t.Errorf("LookupZoneName(%v, %v) = %q; want %q", tt.lat, tt.long, got.Name, tt.want)
		}
	}
}

func TestLookupAlias(t *testing.T) {
	cases := []struct {
		lat, long float64
		want      string
	}{
		/* moscow */
		{48.720828, 44.459613, mskTZ}, // "Europe/Volgograd" alias
		{58.601318, 49.658060, mskTZ}, // "Europe/Kirov" alias
		/* samara */
		{54.308207, 48.390901, samtTZ}, // "Europe/Ulyanovsk" alias
		{46.346593, 48.033217, samtTZ}, // "Europe/Astrakhan" alias
		/* krasnoyarsk */
		{56.482138, 84.969445, kratTZ}, // "Asia/Tomsk" alias
		{55.040269, 82.948042, kratTZ}, // "Asia/Novosibirsk" alias
		{53.757215, 87.130338, kratTZ}, // "Asia/Novokuznetsk" alias
		{53.340574, 83.756804, kratTZ}, // "Asia/Barnaul" alias
		/* yakutsk */
		{52.033567, 113.483689, yaktTZ}, // "Asia/Chita" alias
		{63.772107, 131.651701, yaktTZ}, // "Asia/Khandyga" alias
		/* magadan */
		{64.566189, 143.227966, magtTZ}, // "Asia/Ust-Nera" alias
		{67.459319, 153.707491, magtTZ}, // "Asia/Srednekolymsk" alias
		{46.949921, 142.730061, magtTZ}, // "Asia/Sakhalin" alias
		/* Kamchatka Time aliases */
		{64.733475, 177.512505, pettTZ}, // "Asia/Anadyr" alias
	}
	for _, tt := range cases {
		got, err := LookupZoneName(tt.lat, tt.long)
		if err != nil {
			t.Error(err)
		}

		if got.Name != tt.want {
			t.Errorf("LookupZoneName(%v, %v) = %q; want %q", tt.lat, tt.long, got.Name, tt.want)
		}
	}
}
