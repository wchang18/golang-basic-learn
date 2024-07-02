package sql

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	ID   int64  `gorm:"id"`
	Name string `gorm:"name"`
}

func TestGorm(t *testing.T) {
	sqlDb, sqlMock, _ := sqlmock.New()
	defer sqlDb.Close()

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlMock.ExpectQuery("SELECT id,name FROM `user` WHERE id = ?").WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "jack").AddRow(2, "tony"))

	var users []User
	gormDb.Table("user").Select("id,name").Where("id = ?", 1).Scan(&users)

	for _, user := range users {
		t.Log(user.ID, user.Name)
	}
}
