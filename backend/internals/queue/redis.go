package queue

import (
	"context"
	"encoding/json"

	"github.com/ratneshrt/ramix/internals/models"
	"github.com/redis/go-redis/v9"
)

type Queue struct {
	client *redis.Client
}

func New(addr string) *Queue {
	return &Queue{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (q *Queue) PushJob(ctx context.Context, job models.ExecuteJob) error {
	data, _ := json.Marshal((job))
	return q.client.RPush(ctx, "jobs", data).Err()
}

func (q *Queue) PopJob(ctx context.Context) (*models.ExecuteJob, error) {
	res, err := q.client.BLPop(ctx, 0, "jobs").Result()
	if err != nil {
		return nil, err
	}

	var job models.ExecuteJob
	json.Unmarshal([]byte(res[1]), &job)
	return &job, nil
}
