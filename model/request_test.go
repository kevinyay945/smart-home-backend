package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("request model", func() {
	var r *MRequest
	var mock sqlmock.Sqlmock
	var gdb *gorm.DB

	BeforeEach(func() {
		gdb, mock = mockGorm()

		r = &MRequest{db: gdb}
	})
	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("should select all request in database", func() {
		const sqlSelectAll = `SELECT \* FROM "requests"`
		mock.ExpectQuery(sqlSelectAll).WillReturnRows(sqlmock.NewRows(nil))
		_, err := r.Get()
		Expect(err).ShouldNot(HaveOccurred())
	})
})

func mockGorm() (*gorm.DB, sqlmock.Sqlmock) {
	_db, mock, err := sqlmock.New() // mock sql.DB
	Expect(err).ShouldNot(HaveOccurred())
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: _db,
	}), &gorm.Config{})
	Expect(err).ShouldNot(HaveOccurred())
	return gdb, mock
}
