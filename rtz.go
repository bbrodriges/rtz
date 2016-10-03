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

// Package rtz maps from a latitude and longitude to a timezone of Russian Federation.
//
// It is based on data provided by github.com/bradfitz/latlong package (forked and patched).
package rtz

import (
	"fmt"

	"github.com/bbrodriges/latlong"
)

const (
	usz1TZ = "Europe/Kaliningrad"
	mskTZ  = "Europe/Moscow"
	samtTZ = "Europe/Samara"
	yektTZ = "Asia/Yekaterinburg"
	omstTZ = "Asia/Omsk"
	kratTZ = "Asia/Krasnoyarsk"
	irktTZ = "Asia/Irkutsk"
	yaktTZ = "Asia/Yakutsk"
	vlatTZ = "Asia/Vladivostok"
	magtTZ = "Asia/Magadan"
	pettTZ = "Asia/Kamchatka"
)

// Predefined russian timezones data
var russianTimezones = map[string]Timezone{
	usz1TZ: Timezone{
		Name:         usz1TZ,
		Abbreviation: "USZ1",
		Offset:       2.0,
		RFCOffset:    "+02:00",
		MskOffset:    -1.0,
	},
	mskTZ: Timezone{
		Name:         mskTZ,
		Abbreviation: "MSK",
		Offset:       3.0,
		RFCOffset:    "+03:00",
		MskOffset:    0.0,
	},
	samtTZ: Timezone{
		Name:         samtTZ,
		Abbreviation: "SAMT",
		Offset:       4.0,
		RFCOffset:    "+04:00",
		MskOffset:    1.0,
	},
	yektTZ: Timezone{
		Name:         yektTZ,
		Abbreviation: "YEKT",
		Offset:       5.0,
		RFCOffset:    "+05:00",
		MskOffset:    2.0,
	},
	omstTZ: Timezone{
		Name:         omstTZ,
		Abbreviation: "OMST",
		Offset:       6.0,
		RFCOffset:    "+06:00",
		MskOffset:    3.0,
	},
	kratTZ: Timezone{
		Name:         kratTZ,
		Abbreviation: "KRAT",
		Offset:       7.0,
		RFCOffset:    "+07:00",
		MskOffset:    4.0,
	},
	irktTZ: Timezone{
		Name:         irktTZ,
		Abbreviation: "IRKT",
		Offset:       8.0,
		RFCOffset:    "+08:00",
		MskOffset:    5.0,
	},
	yaktTZ: Timezone{
		Name:         yaktTZ,
		Abbreviation: "YAKT",
		Offset:       9.0,
		RFCOffset:    "+09:00",
		MskOffset:    6.0,
	},
	vlatTZ: Timezone{
		Name:         vlatTZ,
		Abbreviation: "VLAT",
		Offset:       10.0,
		RFCOffset:    "+10:00",
		MskOffset:    7.0,
	},
	magtTZ: Timezone{
		Name:         magtTZ,
		Abbreviation: "MAGT",
		Offset:       11.0,
		RFCOffset:    "+11:00",
		MskOffset:    8.0,
	},
	pettTZ: Timezone{
		Name:         pettTZ,
		Abbreviation: "PETT",
		Offset:       12.0,
		RFCOffset:    "+12:00",
		MskOffset:    9.0,
	},
}

// Data used by github.com/bradfitz/latlong from efele.net (http://efele.net/maps/tz/russia/ in particular)
// has standalone timezones inside larger official russin timezone (see https://en.wikipedia.org/wiki/Time_in_Russia).
// Thats why we need to alias some standalone Go timezones with official timezones.
var timezoneAliases = map[string]string{
	/* Moscow Time aliases */
	"Europe/Volgograd": mskTZ,
	"Europe/Kirov":     mskTZ,
	/* Samara Time aliases*/
	"Europe/Ulyanovsk": samtTZ,
	"Europe/Astrakhan": samtTZ,
	/* Krasnoyarsk Time aliases */
	"Asia/Tomsk":        kratTZ,
	"Asia/Novosibirsk":  kratTZ,
	"Asia/Novokuznetsk": kratTZ,
	"Asia/Barnaul":      kratTZ,
	/* Yakutsk Time aliases */
	"Asia/Chita":    yaktTZ,
	"Asia/Khandyga": yaktTZ,
	/* Magadan Time aliases */
	"Asia/Ust-Nera":      magtTZ,
	"Asia/Srednekolymsk": magtTZ,
	"Asia/Sakhalin":      magtTZ,
	/* Kamchatka Time aliases */
	"Asia/Anadyr": pettTZ,
}

// Timezone holds information about russian timezone
type Timezone struct {
	Name         string  `json:"name"`         // timezone name, e.g. "Europe/Moscow"
	Abbreviation string  `json:"abbreviation"` // timezone abbreviation, e.g. "MSK"
	Offset       float64 `json:"offset"`       // signed UTC offset, e.g. 3.0
	RFCOffset    string  `json:"rfcOffset"`    // signed UTC offset in RFC3399 format, e.g. "+03:00"
	MskOffset    float64 `json:"mskOffset"`    // signed time offset from Europe/Moscow time, e.g. 0.0
}

// LookupZoneName returns the Timezone struct and optional error
// at the given latitude and longitude.
func LookupZoneName(lat, long float64) (Timezone, error) {
	zoneName := latlong.LookupZoneName(lat, long)

	// search timezone
	if rtz, ok := russianTimezones[zoneName]; ok {
		return rtz, nil
	}

	// search by alias
	if rtz, ok := russianTimezones[timezoneAliases[zoneName]]; ok {
		return rtz, nil
	}

	// resistance is futile
	return Timezone{}, fmt.Errorf("Non-russian timezone: %s", zoneName)
}
