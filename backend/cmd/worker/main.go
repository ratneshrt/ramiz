package main

import (
	"context"
	"log"

	"github.com/ratneshrt/ramix/internals/queue"
)

func main() {
	q := queue.New("localhost:6379")

	for {
		job, err := q.PopJob(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Executing job: ", job.JobID)
	}
}
