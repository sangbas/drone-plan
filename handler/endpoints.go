package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstate(ctx echo.Context) error {
	var payload generated.CreateEstateRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.BadRequestError)
	}

	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.BadRequestError)
	}

	id, err := s.Repository.CreateEstate(ctx.Request().Context(), entity.NewEstate(payload.Length, payload.Width))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, entity.InternalServerError)
	}

	var resp generated.CreateEstateResponse
	resp.Id = id.String()
	return ctx.JSON(http.StatusCreated, resp)
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, estateID string) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %s", estateID)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateIdStats(ctx echo.Context, estateID string) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %s", estateID)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstateIdTree(ctx echo.Context, estateID string) error {
	var resp generated.CreateTreeResponse
	resp.Id = ""
	return ctx.JSON(http.StatusCreated, resp)
}
