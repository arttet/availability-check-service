package entrypoint

import (
	"context"
	"log"
	"time"

	"github.com/arttet/validator-service/internal/app/validator-service/repository"
	"github.com/arttet/validator-service/internal/app/validator-service/validator"
	"github.com/arttet/validator-service/internal/database"

	"go.uber.org/zap"
)

func EntryPoint(dsn string, timeout time.Duration) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("failied logger initialization")
	}
	defer logger.Sync() // nolint:errcheck
	zap.ReplaceGlobals(logger)

	db, err := database.NewConnection(dsn, "mysql")
	if err != nil {
		logger.Fatal("failied database initialization", zap.Error(err))
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	repo := repository.NewRepository(db)
	validator.NewValidator(repo).VerifyHosts(ctx) // nolint:errcheck
}
