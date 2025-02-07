package metricsql

import (
	"strings"
)

var aggrFuncs = map[string]bool{
	// See https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
	"sum":          true,
	"min":          true,
	"max":          true,
	"avg":          true,
	"stddev":       true,
	"stdvar":       true,
	"count":        true,
	"count_values": true,
	"bottomk":      true,
	"topk":         true,
	"quantile":     true,
	"group":        true,

	// MetricsQL extension funcs
	"median":         true,
	"limitk":         true,
	"limit_offset":   true,
	"distinct":       true,
	"sum2":           true,
	"geomean":        true,
	"histogram":      true,
	"topk_min":       true,
	"topk_max":       true,
	"topk_avg":       true,
	"topk_median":    true,
	"topk_last":      true,
	"bottomk_min":    true,
	"bottomk_max":    true,
	"bottomk_avg":    true,
	"bottomk_median": true,
	"bottomk_last":   true,
	"any":            true,
	"mad":            true,
	"outliers_mad":   true,
	"outliersk":      true,
	"mode":           true,
	"zscore":         true,
	"quantiles":      true,
}

func isAggrFunc(s string) bool {
	s = strings.ToLower(s)
	return aggrFuncs[s]
}

func isAggrFuncModifier(s string) bool {
	s = strings.ToLower(s)
	switch s {
	case "by", "without":
		return true
	default:
		return false
	}
}
