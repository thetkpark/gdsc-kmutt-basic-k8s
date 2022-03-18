package todo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

type ReqBodyTodo struct {
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}

func NewHandler(db *gorm.DB, l *zap.SugaredLogger) *Handler {
	return &Handler{
		logger: l,
		db:     db,
	}
}

// ListTodos listTodo
// @Summary List todo
// @Tags todo
// @Produce  json
// @Success 200 {array} todo.Todo "List of all todos"
// @Router /api/todos [get]
func (h *Handler) ListTodos(ctx *fiber.Ctx) error {
	var todos []Todo
	if tx := h.db.Find(&todos); tx.Error != nil {
		return h.logAndReturnError(ctx, "Failed to query todos", tx.Error, fiber.StatusInternalServerError)
	}
	return ctx.JSON(todos)
}

// CreateTodo createTodo
// @Summary Create todo
// @Tags todo
// @Accept json
// @Produce  json
// @Param body body todo.ReqBodyTodo true "Todo to create"
// @Success 201 {object} todo.Todo "The created todo"
// @Router /api/todo [post]
func (h *Handler) CreateTodo(ctx *fiber.Ctx) error {
	reqTodo := new(ReqBodyTodo)
	if err := ctx.BodyParser(reqTodo); err != nil {
		return h.logAndReturnError(ctx, "Failed to parse request body", err, fiber.StatusBadRequest)
	}

	todo := &Todo{
		Name:     reqTodo.Name,
		Finished: reqTodo.Finished,
	}

	if tx := h.db.Create(todo); tx.Error != nil {
		return h.logAndReturnError(ctx, "Failed to save todo", tx.Error, fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

// FinishedTodo finishTodo
// @Summary Finish todo
// @Tags todo
// @Produce  json
// @Param id path integer true "id of todo to finish"
// @Success 201 {object} todo.Todo "The updated todo"
// @Router /api/todo/{id} [patch]
func (h *Handler) FinishedTodo(ctx *fiber.Ctx) error {
	todoId, err := ctx.ParamsInt("id")
	if err != nil || todoId == 0 {
		return h.logAndReturnError(ctx, "Failed to parse todo ID", err, fiber.StatusBadRequest)
	}

	var todo Todo
	if err := h.db.First(&todo, todoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.logAndReturnError(ctx, "Todo not found", errors.New("not found"), fiber.StatusNotFound)
		}
		return h.logAndReturnError(ctx, "Failed to query todo with id", err, fiber.StatusInternalServerError)
	}

	todo.Finished = true

	if err := h.db.Save(&todo).Error; err != nil {
		return h.logAndReturnError(ctx, "Failed update finished", err, fiber.StatusInternalServerError)
	}

	return ctx.JSON(todo)
}

// DeleteTodo deleteTodo
// @Summary Delete todo
// @Tags todo
// @Produce  json
// @Param id path integer true "id of todo to delete"
// @Success 201 {object} todo.Todo "The deleted todo"
// @Router /api/todo/{id} [delete]
func (h *Handler) DeleteTodo(ctx *fiber.Ctx) error {
	todoId, err := ctx.ParamsInt("id")
	if err != nil || todoId == 0 {
		return h.logAndReturnError(ctx, "Failed to parse todo ID", err, fiber.StatusBadRequest)
	}

	var todo Todo
	if err := h.db.First(&todo, todoId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return h.logAndReturnError(ctx, "Todo not found", errors.New("not found"), fiber.StatusNotFound)
		}
		return h.logAndReturnError(ctx, "Failed to query todo with id", err, fiber.StatusInternalServerError)
	}

	todo.Finished = true

	if err := h.db.Delete(&todo).Error; err != nil {
		return h.logAndReturnError(ctx, "Failed delete todo", err, fiber.StatusInternalServerError)
	}

	return ctx.JSON(todo)
}

func (h *Handler) logAndReturnError(ctx *fiber.Ctx, msg string, err error, status int) error {
	h.logger.Errorw(msg, "error", err.Error())
	return ctx.Status(status).JSON(fiber.Map{
		"message": msg,
		"error":   err.Error(),
	})
}
