package others

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/model"
	"github.com/fyk7/go-clean-arch-demo-v3/app/domain/repository"
	"github.com/sirupsen/logrus"
)

// transaction系のコードもこちらに書く必要がある

// SQLHandlerとしてConnを抽象化するのもあり
type sqlUserRepository struct {
	Conn *sql.DB
}

func NewSqlUserRepository(db *sql.DB) repository.UserRepository {
	return &sqlUserRepository{
		Conn: db,
	}
}

// ポインタを返すのであればポインタで、実体を返すのであれば実体で統一する。
func (s *sqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []*model.User, err error) {
	rows, err := s.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("failed query to db: %w", err)
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]*model.User, 0)
	for rows.Next() {
		u := model.User{}
		err = rows.Scan(
			&u.ID,
			&u.Email,
		)
		result = append(result, &u)
	}
	return result, nil
}

func (s *sqlUserRepository) fetchOne(ctx context.Context, query string, args ...interface{}) (result *model.User, err error) {
	// protect from SQL Injection
	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return &model.User{}, err
	}
	// single row
	row := stmt.QueryRowContext(ctx, args...)

	u := model.User{}
	err = row.Scan(&u.ID, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *sqlUserRepository) FindAll(ctx context.Context) (res []*model.User, err error) {
	query := `SELECT * FROM user`
	return s.fetch(ctx, query)
}

func (s *sqlUserRepository) FindByEmail(ctx context.Context, email string) (res *model.User, err error) {
	query := `SELECT * FROM user WHERE email=?`
	return s.fetchOne(ctx, query, email)
}

func (s *sqlUserRepository) Save(ctx context.Context, user *model.User) (res *model.User, err error) {
	query := `insert into user (id, email) values (?, ?)`
	// prevent from SQL Injection
	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(ctx, user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	// 引数でもらったuserをそのまま返していいのか??
	return user, nil
}

func (s *sqlUserRepository) Update(ctx context.Context, user *model.User) error {
	query := `update user set email = ? where id = ?`
	stmt, err := s.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}
