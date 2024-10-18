package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetHello(t *testing.T) {

}

func TestPostEstate(t *testing.T) {
	e := echo.New()
	mockRepo := repository.NewMockRepositoryInterface(gomock.NewController(t))
	server := &Server{Repository: mockRepo}
	e.Validator = &CustomValidator{validator: validator.New()}

	t.Run("success", func(t *testing.T) {
		estate := entity.Estate{Length: 10, Width: 5}
		id := uuid.MustParse("ee972400-fa78-41a6-a3b6-8953df7f5a60")
		payload, _ := json.Marshal(estate)

		req := httptest.NewRequest(http.MethodPost, "/estate", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockRepo.EXPECT().CreateEstate(c.Request().Context(), gomock.Any()).Return(id, nil).Times(1)

		// Mock validation method
		c.Set("validate", func(v interface{}) error { return nil })

		if assert.NoError(t, server.PostEstate(c)) {
			var resp generated.CreateEstateResponse
			json.NewDecoder(rec.Body).Decode(&resp)
			assert.NotEmpty(t, resp.Id)
		}
	})
}

func TestPostEstateIdTree(t *testing.T) {
	e := echo.New()
	mockRepo := repository.NewMockRepositoryInterface(gomock.NewController(t))
	server := &Server{Repository: mockRepo}
	e.Validator = &CustomValidator{validator: validator.New()}

	t.Run("success", func(t *testing.T) {
		estateId := uuid.MustParse("ee972400-fa78-41a6-a3b6-8953df7f5a60")
		treeId := uuid.MustParse("39800709-8a33-4d53-bb82-f96c40018f84")
		// tree := entity.Tree{Id: treeId, EstateId: estateId, XAxis: 1, YAxis: 4, Height: 5}
		estate := &entity.Estate{Id: estateId, Length: 10, Width: 5}
		reqPayload := generated.CreateTreeRequest{Height: 5, X: 1, Y: 4}
		payload, _ := json.Marshal(reqPayload)

		req := httptest.NewRequest(http.MethodPost, "/estate/{id}/tree", bytes.NewBuffer(payload))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockRepo.EXPECT().GetEstateById(c.Request().Context(), gomock.Any()).Return(estate, nil).Times(1)
		mockRepo.EXPECT().GetTreesByEstateId(c.Request().Context(), gomock.Any()).Return([]entity.Tree{}, nil).Times(1)
		mockRepo.EXPECT().CreateTree(c.Request().Context(), gomock.Any()).Return(treeId, nil).Times(1)

		// Mock validation method
		c.Set("validate", func(v interface{}) error { return nil })

		if assert.NoError(t, server.PostEstateIdTree(c, estateId.String())) {
			var resp generated.CreateTreeResponse
			json.NewDecoder(rec.Body).Decode(&resp)
			assert.NotEmpty(t, resp.Id)
		}
	})
}

func TestGetEstateIdStats(t *testing.T) {
	e := echo.New()
	mockRepo := repository.NewMockRepositoryInterface(gomock.NewController(t))
	server := &Server{Repository: mockRepo}
	e.Validator = &CustomValidator{validator: validator.New()}

	t.Run("success", func(t *testing.T) {
		estateId := uuid.MustParse("ee972400-fa78-41a6-a3b6-8953df7f5a60")
		estateStat := repository.EstateStat{Count: 1, Max: 2, Min: 1, Median: 2}

		req := httptest.NewRequest(http.MethodGet, "/estate/{id}/stats", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockRepo.EXPECT().GetEstateStat(c.Request().Context(), estateId).Return(estateStat, nil).Times(1)

		// Mock validation method
		c.Set("validate", func(v interface{}) error { return nil })

		if assert.NoError(t, server.GetEstateIdStats(c, estateId.String())) {
			var resp generated.EstateStatResponse
			json.NewDecoder(rec.Body).Decode(&resp)
			assert.NotEmpty(t, resp)
		}
	})
}

func TestGetEstateIdDronePlan(t *testing.T) {
	e := echo.New()
	mockRepo := repository.NewMockRepositoryInterface(gomock.NewController(t))
	server := &Server{Repository: mockRepo}
	e.Validator = &CustomValidator{validator: validator.New()}

	t.Run("success", func(t *testing.T) {
		estateId := uuid.MustParse("ee972400-fa78-41a6-a3b6-8953df7f5a60")
		estate := &entity.Estate{Id: estateId, Length: 10, Width: 5}

		req := httptest.NewRequest(http.MethodGet, "/estate/{id}/drone-plan", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		mockRepo.EXPECT().GetEstateById(c.Request().Context(), estateId).Return(estate, nil).Times(1)
		mockRepo.EXPECT().GetTreesByEstateId(c.Request().Context(), estateId).Return([]entity.Tree{}, nil).Times(1)

		// Mock validation method
		c.Set("validate", func(v interface{}) error { return nil })

		if assert.NoError(t, server.GetEstateIdDronePlan(c, estateId.String(), generated.GetEstateIdDronePlanParams{})) {
			var resp generated.EstateDronePlanResponse
			json.NewDecoder(rec.Body).Decode(&resp)
			assert.NotEmpty(t, resp)
		}
	})
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
