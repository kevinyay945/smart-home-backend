package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"smart-home-backend/lib/pg"
)

var _ = Describe("request model", func() {
	var r *MRequest
	var mock sqlmock.Sqlmock
	var gdb *gorm.DB

	BeforeEach(func() {
		gdb, mock = pg.MockGorm()

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
