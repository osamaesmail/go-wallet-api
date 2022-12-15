package transaction_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go-wallet-api/pkg/transaction"
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

func setupRepo(t *testing.T) (transaction.Repository, repoMocks) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	pgc := postgres.New(postgres.Config{Conn: db})
	gDB, err := gorm.Open(pgc)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating grom database connection", err)
	}
	dbMock.MatchExpectationsInOrder(false)
	return transaction.NewRepository(gDB), repoMocks{db, dbMock}
}

func TestRepository_Create(t *testing.T) {
	t.Run(
		"success", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			selectRows := []string{"id", "user_id", "balance", "currency", "created_at", "updated_at", "deleted_at"}
			
			selectValues := []driver.Value{
				newModel.FromAccount,
				uuid.New(),
				decimal.New(1000, 0),
				"USD",
				time.Now(),
				time.Now(),
				nil,
			}
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			selectArgs := []driver.Value{
				req.Amount,
				newModel.FromAccount,
			}
			addArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				"USD",
				sqlmock.AnyArg(),
				newModel.ToAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			qSelect := `SELECT * FROM "accounts" WHERE balance - $1 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $2`
			qAdd := `UPDATE "accounts" SET "balance"=balance + $1,"updated_at"=$2 WHERE currency = $3 AND id != $4 AND "accounts"."deleted_at" IS NULL AND "id" = $5`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qSelect)).WithArgs(selectArgs...).
				WillReturnRows(sqlmock.NewRows(selectRows).AddRow(selectValues...))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qAdd)).WithArgs(addArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectCommit()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.NoError(t, err)
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - create", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnError(errors.New("err"))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "err", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - subtract balance error", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnError(errors.New("err"))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "err", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - Insufficient funds", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnResult(sqlmock.NewResult(0, 0))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "From account not found or insufficient funds", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - select funded account", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			selectArgs := []driver.Value{
				req.Amount,
				newModel.FromAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			qSelect := `SELECT * FROM "accounts" WHERE balance - $1 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $2`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qSelect)).WithArgs(selectArgs...).
				WillReturnError(errors.New("err"))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "err", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - add balance error", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			selectRows := []string{"id", "user_id", "balance", "currency", "created_at", "updated_at", "deleted_at"}
			
			selectValues := []driver.Value{
				newModel.FromAccount,
				uuid.New(),
				decimal.New(1000, 0),
				"USD",
				time.Now(),
				time.Now(),
				nil,
			}
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			selectArgs := []driver.Value{
				req.Amount,
				newModel.FromAccount,
			}
			addArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				"USD",
				sqlmock.AnyArg(),
				newModel.ToAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			qSelect := `SELECT * FROM "accounts" WHERE balance - $1 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $2`
			qAdd := `UPDATE "accounts" SET "balance"=balance + $1,"updated_at"=$2 WHERE currency = $3 AND id != $4 AND "accounts"."deleted_at" IS NULL AND "id" = $5`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qSelect)).WithArgs(selectArgs...).
				WillReturnRows(sqlmock.NewRows(selectRows).AddRow(selectValues...))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qAdd)).WithArgs(addArgs...).
				WillReturnError(errors.New("err"))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "err", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	t.Run(
		"error - to account not found", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.CreateRequest{
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
			}
			model := transaction.Transaction{
				FromAccount: req.FromAccount,
				ToAccount:   req.ToAccount,
				Amount:      req.Amount,
			}
			newModel := model
			newModel.ID = uuid.New()
			newModel.CreatedAt = time.Now()
			
			selectRows := []string{"id", "user_id", "balance", "currency", "created_at", "updated_at", "deleted_at"}
			
			selectValues := []driver.Value{
				newModel.FromAccount,
				uuid.New(),
				decimal.New(1000, 0),
				"USD",
				time.Now(),
				time.Now(),
				nil,
			}
			
			args := []driver.Value{
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
				newModel.ToAccount,
				sqlmock.AnyArg(),
			}
			subtractArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				newModel.Amount,
				newModel.FromAccount,
			}
			selectArgs := []driver.Value{
				req.Amount,
				newModel.FromAccount,
			}
			addArgs := []driver.Value{
				newModel.Amount,
				sqlmock.AnyArg(),
				"USD",
				sqlmock.AnyArg(),
				newModel.ToAccount,
			}
			
			// mocks
			q := `INSERT INTO "transactions" ("id","amount","from_account","to_account","created_at") VALUES ($1,$2,$3,$4,$5)`
			qSubtract := `UPDATE "accounts" SET "balance"=balance - $1,"updated_at"=$2 WHERE balance - $3 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $4`
			qSelect := `SELECT * FROM "accounts" WHERE balance - $1 >= 0 AND "accounts"."deleted_at" IS NULL AND "id" = $2`
			qAdd := `UPDATE "accounts" SET "balance"=balance + $1,"updated_at"=$2 WHERE currency = $3 AND id != $4 AND "accounts"."deleted_at" IS NULL AND "id" = $5`
			m.dbMock.ExpectBegin()
			m.dbMock.ExpectExec(regexp.QuoteMeta(q)).WithArgs(args...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qSubtract)).WithArgs(subtractArgs...).
				WillReturnResult(sqlmock.NewResult(0, 1))
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qSelect)).WithArgs(selectArgs...).
				WillReturnRows(sqlmock.NewRows(selectRows).AddRow(selectValues...))
			m.dbMock.ExpectExec(regexp.QuoteMeta(qAdd)).WithArgs(addArgs...).
				WillReturnResult(sqlmock.NewResult(0, 0))
			m.dbMock.ExpectRollback()
			
			// call method
			_, err := r.Create(model)
			
			// assert
			assert.Error(t, err)
			assert.Equal(t, "To Account not found or currency not matched", err.Error())
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
			model := transaction.DTO{
				ID:          uuid.New(),
				FromAccount: uuid.New(),
				ToAccount:   uuid.New(),
				Amount:      decimal.New(100, 0),
				CreatedAt:   time.Now(),
				Currency:    "USD",
			}
			models := []transaction.DTO{model, model}
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}
			
			rows := []string{"id", "from_account", "to_account", "amount", "created_at", "currency"}
			values := []driver.Value{
				model.ID, model.FromAccount, model.ToAccount, model.Amount, model.CreatedAt, model.Currency,
			}
			
			// mocks
			q := `SELECT transactions.*,a_from.currency currency FROM "transactions" JOIN accounts a_from ON transactions.from_account = a_from.id JOIN accounts a_to ON transactions.to_account = a_to.id WHERE "a_from"."user_id" = $1 OR "a_to"."user_id" = $2`
			qCount := `SELECT count(*) FROM "transactions" JOIN accounts a_from ON transactions.from_account = a_from.id JOIN accounts a_to ON transactions.to_account = a_to.id WHERE "a_from"."user_id" = $1 OR "a_to"."user_id" = $2`
			m.dbMock.ExpectQuery(regexp.QuoteMeta(q)).
				WithArgs(req.UserID, req.UserID).
				WillReturnRows(sqlmock.NewRows(rows).AddRow(values...).AddRow(values...))
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qCount)).
				WithArgs(req.UserID, req.UserID).
				WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(2))
			
			// call method
			resp, total, err := r.List(req)
			
			// assert
			assert.NoError(t, err)
			assert.Equal(t, int64(2), total)
			assert.Equal(t, models, resp)
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
	
	t.Run(
		"error", func(t *testing.T) {
			r, m := setupRepo(t)
			defer m.db.Close()
			
			// data
			req := transaction.ListRequest{
				UserID: uuid.New(),
			}
			
			// mocks
			qCount := `SELECT count(*) FROM "transactions" JOIN accounts a_from ON transactions.from_account = a_from.id JOIN accounts a_to ON transactions.to_account = a_to.id WHERE "a_from"."user_id" = $1 OR "a_to"."user_id" = $2`
			m.dbMock.ExpectQuery(regexp.QuoteMeta(qCount)).
				WithArgs(req.UserID, req.UserID).
				WillReturnError(errors.New("err"))
			
			// call method
			_, _, err := r.List(req)
			
			// assert
			assert.NotNil(t, err)
			assert.Equal(t, "err", err.Error())
			assert.NoError(t, m.dbMock.ExpectationsWereMet())
		},
	)
}
