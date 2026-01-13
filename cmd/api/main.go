package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ratneshrt/ramix/internals/models"
	"github.com/ratneshrt/ramix/internals/queue"
)

func main() {
	q := queue.New("localhost:6379")

	http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request) {
		var req models.ExecuteRequest
		json.NewDecoder(r.Body).Decode(&req)

		job := models.ExecuteJob{
			JobID:    uuid.New().String(),
			Language: req.Language,
			Code:     req.Code,
			Input:    req.Input,
		}

		q.PushJob(context.Background(), job)

		json.NewEncoder(w).Encode(map[string]string{
			"job_id": job.JobID,
		})
	})
	http.ListenAndServe(":8080", nil)
}
