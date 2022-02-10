package v1

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"smart-home-backend/lib/pg/schema"
	modelMock "smart-home-backend/model/mock"
	"smart-home-backend/utils"
	"strings"
	"time"
)

var _ = Describe("Request", func() {
	var mockCtrl *gomock.Controller
	var mockRequestModel *modelMock.MockIRequest
	var e *echo.Echo
	BeforeEach(func() {
		e = echo.New()
		e.Validator = &utils.CustomValidator{Validator: validator.New()}
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequestModel = modelMock.NewMockIRequest(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Test For GetRequests", func() {

		It("should be success", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			requestRoute := &requestRoute{
				Request: mockRequestModel,
			}
			mockRequestModel.EXPECT().Get().Return(nil, nil)

			err := requestRoute.GetRequests(c)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(rec.Code).Should(Equal(http.StatusOK))
		})

		It("should be fail", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			mockRequestModel.EXPECT().Get().Return(nil, errors.New("fail to get data"))
			requestRoute := &requestRoute{
				Request: mockRequestModel,
			}

			err := requestRoute.GetRequests(c)
			Expect(err).Should(HaveOccurred())
			err2, ok := err.(*echo.HTTPError)
			Expect(ok).Should(BeTrue())
			Expect(err2.Code).Should(Equal(http.StatusInternalServerError))
		})

	})
	Describe("Test For CreateRequest", func() {
		createUuid := "19ef9082-436c-4f4c-9145-6ba06819df84"
		nowTime := time.Now()
		BeforeEach(func() {
			newUuidV4 = func() (uuid.UUID, error) {
				byteUuid, _ := uuid.Parse(createUuid)
				return uuid.UUID(byteUuid), nil
			}
			getNow = func() time.Time {
				return nowTime
			}
		})

		It("correct data format and return 201", func() {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`
{
	"name": "correct name"
}
`))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			requestRoute := &requestRoute{
				Request: mockRequestModel,
			}
			mockRequestModel.EXPECT().Save(&schema.Request{
				Uuid:     createUuid,
				CreateAt: nowTime,
				UpdateAt: nowTime,
				Name:     "correct name",
			}).Return(schema.Request{}, nil)

			err := requestRoute.CreateRequest(c)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(rec.Code).Should(Equal(http.StatusCreated))
		})

		It("wrong data format and return 400 and reason", func() {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`
{}
`))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			requestRoute := &requestRoute{
				Request: mockRequestModel,
			}

			err := requestRoute.CreateRequest(c)

			Expect(err).Should(HaveOccurred())
			err2, ok := err.(*echo.HTTPError)
			Expect(ok).Should(BeTrue())
			Expect(err2.Code).Should(Equal(http.StatusBadRequest))
			validatorErr, ok2 := err2.Message.(validator.ValidationErrors)
			Expect(ok2).Should(BeTrue())
			Expect(len(validatorErr)).Should(Equal(1))
			for _, fieldError := range validatorErr {
				if fieldError.Field() == "Name" {
					Expect(fieldError.Error()).Should(ContainSubstring("required"))
				}
			}

		})

	})
})
