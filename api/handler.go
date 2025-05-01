package api

import (
	"encoding/json"
	"net/http"
	"go-monitoring-system/metrics"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	stats := metrics.GetSystemStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
