package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/varun053101/GO-TASK/internal/repository"
)

// handles user related http requests
type UserHandler struct {
	repo *repository.UserRepository
}

// create a new user handler
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// request body for creating user
type createUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"` // yyyy-mm-dd
}

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// parse date
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid dob format",
		})
	}

	user, err := h.repo.CreateUser(
		c.Context(),
		req.Name,
		pgtype.Date{Time: dob, Valid: true},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// calculate age
	age := int(time.Since(user.Dob.Time).Hours() / 24 / 365)

	return c.Status(201).JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  age,
	})
}
