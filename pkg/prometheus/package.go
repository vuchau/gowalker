// Copyright 2018 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Unknwon/gowalker/models"
)

var (
	totalPackagesGaugeFunc = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "gowalker",
		Subsystem: "package",
		Name:      "total",
		Help:      "Number of total packages",
	}, func() float64 {
		return float64(models.NumTotalPackages())
	})
	monthlyActivePackagesGaugeFunc = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "gowalker",
		Subsystem: "package",
		Name:      "monthly_active",
		Help:      "Number of monthly active packages",
	}, func() float64 {
		return float64(models.NumMonthlyActivePackages())
	})
	weeklyActivePackagesGaugeFunc = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "gowalker",
		Subsystem: "package",
		Name:      "weekly_active",
		Help:      "Number of weekly active packages",
	}, func() float64 {
		return float64(models.NumWeeklyActivePackages())
	})
	dailyActivePackagesGaugeFunc = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "gowalker",
		Subsystem: "package",
		Name:      "daily_active",
		Help:      "Number of daily active packages",
	}, func() float64 {
		return float64(models.NumDailyActivePackages())
	})
)

func init() {
	prometheus.MustRegister(
		totalPackagesGaugeFunc,
		monthlyActivePackagesGaugeFunc,
		weeklyActivePackagesGaugeFunc,
		dailyActivePackagesGaugeFunc,
	)
}
