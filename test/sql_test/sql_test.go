package sql_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		return
	}
	return
}

func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	//验证测试完所有sql
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSqlQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery("SELECT id,name,address FROM users WHERE id = ?").WithArgs(100).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "address"}).
			AddRow(100, "tom", "beijing").
			AddRow(200, "jame", "shanghai"))

	rows, err := db.Query("SELECT id,name,address FROM users WHERE id = ?", 100)
	if err != nil {
		t.Fatalf("error was not expected while updating to database: %s", err)
	}
	for rows.Next() {
		var (
			id      int
			name    string
			address string
		)
		rows.Scan(&id, &name, &address)
		t.Logf("id: %d, name: %s, address: %s", id, name, address)
	}
}
