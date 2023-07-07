package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DeneesK/go-mart/internal/api-go-mart/entities"
	"github.com/DeneesK/go-mart/internal/storage-go-mart"
	"golang.org/x/crypto/bcrypt"
)

type PostgresRepo struct {
	db *storagegomart.Postgres
}

func NewRepository(db *storagegomart.Postgres) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (repo *PostgresRepo) Register(ctx context.Context, user *entities.User) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("this password is not allowed: %w", err)
	}

	user.CryptPassword = cryptPassword

	var id int64

	insertUser := `INSERT INTO users (login, password) VALUES($1, $2)
                      RETURNING id`

	err = tx.QueryRowContext(ctx, insertUser, user.Login, user.CryptPassword).Scan(&id)
	if err != nil {
		return err
	}

	user.ID = id

	return tx.Commit()
}

func (repo *PostgresRepo) Login(ctx context.Context, user *entities.User) error {
	selectUser := `SELECT id, password
				   FROM users
				   WHERE login = $1`

	err := repo.db.QueryRowContext(ctx, selectUser, user.Login).Scan(&user.ID, &user.CryptPassword)
	if !errors.Is(err, nil) && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}

func (repo *PostgresRepo) GetUserOrders(ctx context.Context, login string) ([]byte, error) {
	var result []byte
	return result, nil
}

func (repo *PostgresRepo) LoadOrder(ctx context.Context, login string, orderID string) error {
	return nil
}

func (repo *PostgresRepo) GetBalanceInfo(login string) ([]byte, error) {
	var result []byte
	return result, nil
}

func (repo *PostgresRepo) Withdraw(login string, orderID string, sum float64) error {
	return nil
}

func (repo *PostgresRepo) GetWithdrawals(login string) ([]byte, error) {
	var result []byte
	return result, nil
}
