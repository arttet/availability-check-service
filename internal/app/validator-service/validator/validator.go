package validator

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/arttet/validator-service/internal/app/validator-service/repository"
	"github.com/arttet/validator-service/internal/model"

	"go.uber.org/zap"
)

type Validator interface {
	VerifyHosts(ctx context.Context) error
}

func NewValidator(repo repository.Repository) Validator {
	return &validator{
		repo: repo,
	}
}

type validator struct {
	repo repository.Repository
}

func (v *validator) VerifyHosts(ctx context.Context) error {
	checks, err := v.repo.ListChecks(ctx)
	if err != nil {
		zap.L().Error("failed to list the data", zap.Error(err))
		return err
	}

	wg := &sync.WaitGroup{}

	for _, check := range checks {
		wg.Add(1)
		go func(ctx context.Context, check *model.Check, wg *sync.WaitGroup) {
			defer wg.Done()
			v.repo.UpdateStatus(ctx, getStatus(check)) // nolint:errcheck
		}(ctx, check, wg)
	}

	wg.Wait()

	return nil
}

func getStatus(check *model.Check) *model.Check {
	status := [3]bool{}
	errors := [3]error{}

	address := fmt.Sprintf("%s:%d", check.Host, check.Port)
	timeout := time.Duration(check.Timeout) * time.Millisecond

	decisionIndex := -1
	for i := 0; i < 3; i++ {
		if i == 2 {
			if status[0] == status[1] {
				decisionIndex = 1
			} else {
				decisionIndex = 2
			}
		}

		_, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			errors[i] = err
		} else {
			status[i] = true
		}
	}

	check.Status = model.GetStatus(status[decisionIndex])
	if err := errors[decisionIndex]; err != nil {
		check.FailMessage = err.Error()
	}

	return check
}
