// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type WeatherReport byte

const (
	WeatherReportCurrent WeatherReport = 0
	// WeatherReportForecast WeatherReport = 1  // [DUPLICATE!] Deprecated use hourly_forecast instead
	WeatherReportHourlyForecast WeatherReport = 1
	WeatherReportDailyForecast  WeatherReport = 2
	WeatherReportInvalid        WeatherReport = 0xFF
)

func (w WeatherReport) Byte() byte { return byte(w) }

func (w WeatherReport) String() string {
	switch w {
	case WeatherReportCurrent:
		return "current"
	case WeatherReportHourlyForecast:
		return "hourly_forecast"
	case WeatherReportDailyForecast:
		return "daily_forecast"
	default:
		return "WeatherReportInvalid(" + strconv.Itoa(int(w)) + ")"
	}
}

// FromString parse string into WeatherReport constant it's represent, return WeatherReportInvalid if not found.
func WeatherReportFromString(s string) WeatherReport {
	switch s {
	case "current":
		return WeatherReportCurrent
	case "hourly_forecast":
		return WeatherReportHourlyForecast
	case "daily_forecast":
		return WeatherReportDailyForecast
	default:
		return WeatherReportInvalid
	}
}

// List returns all constants.
func ListWeatherReport() []WeatherReport {
	return []WeatherReport{
		WeatherReportCurrent,
		// WeatherReportForecast,
		WeatherReportHourlyForecast,
		WeatherReportDailyForecast,
	}
}