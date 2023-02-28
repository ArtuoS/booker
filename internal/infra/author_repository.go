package infra

import (
	"database/sql"
)

type AuthorRepository struct {
	Db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return AuthorRepository{
		Db: db,
	}
}

func (r *AuthorRepository) Execute(cmd string) error {
	stmt, err := r.Db.Prepare(cmd)
	if err != nil {
		println(cmd)
		println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}
