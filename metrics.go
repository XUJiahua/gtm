package gtm

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CurrentOpLogTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "monstache_current_op_log_ts",
		Help: "Current op log timestamp",
	})
)
