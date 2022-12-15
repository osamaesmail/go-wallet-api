package account_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go-api-grpc/pkg/account"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

type repoMocks struct {
	db     *sql.DB
	dbMock sqlmock.Sqlmock
}

func setupRepo(t *testing.T) (account.Repository, repoMocks) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	pgc := postgres.New(postgres.Config{Conn: db})
	gDB, err := gorm.Open(pgc)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating grom database connection", err)
	}
	return account.NewRepository(gDB), repoMocks{db, dbMock}
}

func TestRepository_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "EGP",
			}
			model := account.Account{
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.UserID,
				newModel.Balance,
				newModel.Currency,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			}
			
			// mocks
			q := `INSERT INTO "accounts" ("id","user_id","balance","currency","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectCommit()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.NoError(t, err)
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	
	t.Run(
		"error", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := account.CreateRequest{
				UserID:   uuid.New(),
				Balance:  decimal.New(100, 0),
				Currency: "EGP",
			}
			model := account.Account{
				UserID:   req.UserID,
				Balance:  req.Balance,
				Currency: req.Currency,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.UserID,
				newModel.Balance,
				newModel.Currency,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			}
			
			// mocks
			q := `INSERT INTO "accounts" ("id","user_id","balance","currency","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).WillReturnError(errors.New("err"))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.NotNil(t, err)
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
}

func TestRepository_List(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			model := account.Account{
				ID:        uuid.New(),
				UserID:    uuid.New(),
				Balance:   decimal.New(100, 0),
				Currency:  "EGP",
				CreatedAt: time.Now(),
			}
			models := []account.Account{model, model}
			req := account.ListRequest{
				UserID: model.UserID,
			}
			
			rows := []string{"id", "user_id", "balance", "currency", "created_at"}
			values := []driver.Value{model.ID, model.UserID, model.Balance, model.Currency, model.CreatedAt}
			
			// mocks
			q := `SELECT * FROM "accounts" WHERE "user_id" = $1 AND "accounts"."deleted_at" IS NULL`
			m.dbMock.ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(model.UserID).
				WillReturnRows(sqlmock.NewRows(rows).AddRow(values...).AddRow(values...))
			
			// call method
			resp, err := r.List(req)
			
			// assert
			assert.NoError(t, err)
			assert.Len(t, resp, 2)
			assert.Equal(t, models, resp)
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	
	t.Run(
		"error - not found", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			model := account.Account{
				ID:        uuid.Nil,
				UserID:    uuid.New(),
				Balance:   decimal.New(100, 0),
				Currency:  "EGP",
				CreatedAt: time.Now(),
			}
			req := account.ListRequest{
				UserID: model.UserID,
			}
			
			rows := []string{"id", "user_id", "balance", "currency", "created_at"}
			
			// mocks
			q := `SELECT * FROM "accounts" WHERE "user_id" = $1 AND "accounts"."deleted_at" IS NULL`
			m.dbMock.ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(model.UserID).
				WillReturnRows(sqlmock.NewRows(rows))
			
			// call method
			_, err := r.List(req)
			
			// assert
			assert.NotNil(t, err)
			assert.Equal(t, "User not found", err.Error())
			assert.Nil(t, m.dbMock.ExpectationsWereMet())
		},
	)
}
