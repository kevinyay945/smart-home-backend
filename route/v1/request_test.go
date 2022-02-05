package v1

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"smart-home-backend/model"
)

var _ = Describe("Request", func() {
	var mockCtrl *gomock.Controller
	var mockRequestModel *model.MockIRequest
	var e *echo.Echo
	BeforeEach(func() {
		e = echo.New()
		mockCtrl = gomock.NewController(GinkgoT())
		mockRequestModel = model.NewMockIRequest(mockCtrl)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
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
