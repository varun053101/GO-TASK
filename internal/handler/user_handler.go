package handler

import (
	"strconv"
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

type updateUserRequest struct {
	Name string `json:"name"`
	DOB  string `json:"dob"`
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

// GET /users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	user, err := h.repo.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	age := int(time.Since(user.Dob.Time).Hours() / 24 / 365)

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  age,
	})
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	var req updateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid dob format",
		})
	}

	user, err := h.repo.UpdateUser(
		c.Context(),
		int32(id),
		req.Name,
		pgtype.Date{Time: dob, Valid: true},
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
	})
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	err = h.repo.DeleteUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.repo.ListUsers(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	response := make([]fiber.Map, 0, len(users))
	for _, user := range users {
		age := int(time.Since(user.Dob.Time).Hours() / 24 / 365)

		response = append(response, fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Time.Format("2006-01-02"),
			"age":  age,
		})
	}

	return c.JSON(response)
}