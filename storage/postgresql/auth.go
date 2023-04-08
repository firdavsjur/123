package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) Register(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		query string
		id    string
	)
	id = uuid.New().String()
	access_token := helper.MakeJWT(id, req.Login, req.Password)
	query = `
		INSERT INTO users(
			id, 
			first_name,
			last_name,
			login,
			password,
			access_token
		)
		VALUES (
			$1,$2,$3,$4,$5,$6)
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Login,
		req.Password,
		access_token,
	)
	if err != nil {
		log.Print("ok")
		return "", err
	}

	return access_token, nil
}

func (r *authRepo) Login(ctx context.Context, req *models.Login1) (string, error) {

	var (
		query    string
		password string
		token    string
		id       string
	)

	query = `
		SELECT
			id,
			password
		FROM users
		WHERE login = $1 
	`

	err := r.db.QueryRow(ctx, query, req.Login).Scan(
		&id,
		&password,
	)
	if err != nil {
		return "Login Topilmadi", errors.New("Login Not Found")
	}
	if password != req.Password {
		return "password inCorrect", errors.New("password Incorrect")
	}
	fmt.Print("salom")
	token = helper.MakeJWT(id, req.Login, req.Password)
	query = `
		UPDATE users SET access_token=$1 WHERE login=$2 and password=$3
	`
	_, err = r.db.Exec(ctx, query, token, req.Login, req.Password)
	if err != nil {
		log.Print("update bolmadi")
		return "", err
	}
	return token, nil
}
