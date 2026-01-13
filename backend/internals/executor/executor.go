package executor

import "github.com/ratneshrt/ramix/internals/models"

type Executor interface {
	Run(job models.ExecuteJob) (*models.ExecuteResult, error)
}
