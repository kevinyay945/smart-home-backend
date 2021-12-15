package pg

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockGorm() (*gorm.DB, sqlmock.Sqlmock) {
	_db, mock, err := sqlmock.New() // mock sql.DB
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: _db,
	}), &gorm.Config{})
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	return gdb, mock
}
