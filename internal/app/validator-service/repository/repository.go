package repository

import (
	"context"
	"database/sql"

	"github.com/arttet/validator-service/internal/database"
	"github.com/arttet/validator-service/internal/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository interface {
	ListChecks(ctx context.Context) (model.Checks, error)
	UpdateStatus(ctx context.Context, check *model.Check) (bool, error)
}

// NewRepository creates a new instance of Repository Service.
func NewRepository(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

const TableName = "checks"

var (
	SelectColumns = []string{
		"id",
		"host",
		"port",
		"status",
		"timeout",
	}
)

type repo struct {
	db *sqlx.DB
}

func (r *repo) ListChecks(ctx context.Context) (model.Checks, error) {
	var checks model.Checks

	err := database.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
		sb := squirrel.StatementBuilder.
			Select(SelectColumns...).
			From(TableName).
			OrderBy("id ASC").
			RunWith(r.db)

		query, args, err := sb.ToSql()
		if err != nil {
			return err
		}

		return r.db.SelectContext(ctx, &checks, query, args...)
	})

	if err != nil {
		return nil, errors.Wrap(err, "repo.ListChecks")
	}

	if len(checks) == 0 {
		return nil, errors.Wrap(sql.ErrNoRows, "repo.ListChecks")
	}

	return checks, nil
}

func (r *repo) UpdateStatus(ctx context.Context, check *model.Check) (bool, error) {
	var updated bool

	err := database.WithTx(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) error {
		sb := squirrel.StatementBuilder.
			Update(TableName).
			Set("status", check.Status).
			Set("fail_message", check.FailMessage).
			Where(squirrel.Eq{"id": check.ID}).
			RunWith(r.db)

		result, err := sb.ExecContext(ctx)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			updated = true
		}

		return err
	})

	if err != nil {
		return false, errors.Wrap(err, "repo.UpdateStatus")
	}

	return updated, nil
}
