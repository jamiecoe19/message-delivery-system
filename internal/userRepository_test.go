package internal_test

import (
	"mds/internal"
	"mds/internal/message"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm db", err)
	}

	userRepo := internal.NewUserRepository(gdb)

	user := message.User{UserID: 12312312, Name: "test"}
	mock.ExpectBegin()
	mock.ExpectExec(
		"INSERT INTO `users` (`user_id`,`name`) VALUES (?,?)").WithArgs(sqlmock.AnyArg(), user.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	_, err = userRepo.Create(user)
	if err != nil {
		t.Errorf("expected null, got %s", err)
	}

}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm db", err)
	}

	userRepo := internal.NewUserRepository(gdb)

	ID := uint64(2312312121212)
	name := "test"
	mock.ExpectQuery("SELECT * FROM `users` WHERE (name = ?) LIMIT 1").
		WithArgs(name).
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "name"}).
				AddRow(ID, name))

	user, err := userRepo.Get(name)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if user.UserID != ID {
		assert.Equal(t, ID, user.UserID)
	}
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening gorm db", err)
	}

	userRepo := internal.NewUserRepository(gdb)
	test1ID := uint64(2312312121212)
	test2ID := uint64(14234234234)

	mock.ExpectQuery("SELECT * FROM `users`").
		WillReturnRows(
			sqlmock.NewRows([]string{"user_id", "name"}).
				AddRow(test1ID, "test1").
				AddRow(test2ID, "test2"))

	users, err := userRepo.GetAll()
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if users[0].UserID != test1ID {
		assert.Equal(t, test1ID, users[0].UserID)
	}

	if users[1].UserID != test2ID {
		assert.Equal(t, test2ID, users[1].UserID)
	}
}
