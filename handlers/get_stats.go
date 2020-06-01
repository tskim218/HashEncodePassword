package handlers

import (
	"encoding/json"
	"github.com/tskim218/HashEncodePassword/datasource"
	"log"
	"net/http"
)

type ResponseStats struct {
	Total int64		`json:"total"`
	Average int64  	`json:"average"`
}

func GetStats(stats datasource.Stats) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reStats := ResponseStats{}

		log.Printf("----- request Statistics End-Point ------\n\n")

		reStats.Total = stats.GetTotal()
		totalProcessedTime := stats.GetTotalTime()

		reStats.Average = totalProcessedTime / reStats.Total

		statsJson, err := json.Marshal(reStats)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(statsJson)
		w.Write([]byte("\n"))

		log.Printf("total: %d, average: %d\n", reStats.Total, reStats.Average)
	})
}
