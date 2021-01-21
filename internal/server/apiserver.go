package apiserver

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/quizardapp/auth-api/internal/store/sqlstore"
)

func Start() error {
	db, err := newDB()
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)

	srv := newServer(store)

	return http.ListenAndServe(":"+os.Getenv("PORT"), srv)
}

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("CLEARDB_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
