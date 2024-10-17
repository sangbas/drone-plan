package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
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

func (s *Server) PostEstateIdTree(ctx echo.Context, id string) error {
	var payload generated.CreateTreeRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.BadRequestError)
	}

	if err := ctx.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.BadRequestError)
	}

	estateId, err := uuid.Parse(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, entity.BadRequestError)
	}

	estate, err := s.Repository.GetEstateById(ctx.Request().Context(), estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, entity.NewErrorResponse("estate id not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, entity.InternalServerError)
	}

	if payload.X > estate.Length {
		return ctx.JSON(http.StatusBadRequest, entity.NewErrorResponse("x-axis is out of bound"))
	}
	if payload.Y > estate.Width {
		return ctx.JSON(http.StatusBadRequest, entity.NewErrorResponse("y-axis is out of bound"))
	}

	trees, err := s.Repository.GetTreesByEstateId(ctx.Request().Context(), estateId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, entity.InternalServerError)
	}

	for _, tree := range trees {
		if tree.XAxis == payload.X && tree.YAxis == payload.Y {
			return ctx.JSON(http.StatusBadRequest, entity.NewErrorResponse("plot already has tree,"))
		}
	}

	treeId, err := s.Repository.CreateTree(ctx.Request().Context(), entity.NewTree(estateId, payload.X, payload.Y, payload.Height))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, entity.InternalServerError)
	}

	var resp generated.CreateTreeResponse
	resp.Id = treeId.String()
	return ctx.JSON(http.StatusCreated, resp)
}

func (s *Server) GetEstateIdStats(ctx echo.Context, estateID string) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %s", estateID)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, estateID string) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %s", estateID)
	return ctx.JSON(http.StatusOK, resp)
}
