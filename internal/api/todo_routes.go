package api

import (
	"net/http"
	"strconv"

	"github.com/HtetLinMaung/todo/internal/middleware"
	"github.com/HtetLinMaung/todo/internal/model"
	"github.com/HtetLinMaung/todo/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoRoute struct {
	todoService *service.TodoService
}

func NewTodoRoute(todoService *service.TodoService) *TodoRoute {
	return &TodoRoute{todoService: todoService}
}

func (tr *TodoRoute) TodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/api/todos", middleware.AuthMiddleware())
	todoGroup.POST("/", tr.AddTodo)
	todoGroup.GET("/", tr.GetTodos)
	todoGroup.GET("/:todoId", tr.GetTodoByIdMiddleware(), tr.GetTodoById)
	todoGroup.PUT("/:todoId", tr.GetTodoByIdMiddleware(), tr.UpdateTodo)
	todoGroup.DELETE("/:todoId", tr.GetTodoByIdMiddleware(), tr.DeleteTodo)
}

func (tr *TodoRoute) AddTodo(c *gin.Context) {
	creatorId := c.GetInt64("user_id")
	var todoRequest *model.TodoRequest

	if err := c.ShouldBindJSON(&todoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if todoRequest.Label == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Label cannot be empty!",
		})
		return
	}

	todo_id, err := tr.todoService.AddTodo(todoRequest, creatorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error adding todo to database!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Todo added successfully.",
		"data":    todo_id,
	})
}

func (tr *TodoRoute) GetTodos(c *gin.Context) {
	var todoQuery model.TodoQuery

	if err := c.ShouldBindQuery(&todoQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	userId := c.GetInt64("user_id")
	role := c.GetString("role")

	var search string
	if todoQuery.Search != nil {
		search = *todoQuery.Search
	}

	var page uint
	if todoQuery.Page != nil {
		page = *todoQuery.Page
	}

	var perPage uint
	if todoQuery.PerPage != nil {
		perPage = *todoQuery.PerPage
	}

	result, err := tr.todoService.GetTodos(search, page, perPage, userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error fetching todos from database!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"message":     "Successful.",
		"data":        result.Data,
		"total":       result.Total,
		"page":        result.Page,
		"per_page":    result.PerPage,
		"page_counts": result.PageCounts,
	})
}

func (tr *TodoRoute) GetTodoByIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoId, err := strconv.ParseInt(c.Param("todoId"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
		c.Set("todoId", todoId)

		userId := c.GetInt64("user_id")
		role := c.GetString("role")
		todo, err := tr.todoService.GetTodoById(todoId, userId, role)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Error fetching todo from database!",
			})
			return
		}
		if todo == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Todo not found!",
			})
			return
		}
		c.Set("todo", todo)
		c.Next()
	}
}

func (tr *TodoRoute) GetTodoById(c *gin.Context) {
	value, exists := c.Get("todo")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Something went wrong!",
		})
		return
	}
	todo, ok := value.(*model.Todo)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Something went wrong!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Todo fetched successfully.",
		"data":    todo,
	})
}

func (tr *TodoRoute) UpdateTodo(c *gin.Context) {
	var todoRequest *model.TodoRequest

	if err := c.ShouldBindJSON(&todoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if todoRequest.Label == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Label cannot be empty!",
		})
		return
	}

	todoId := c.GetInt64("todoId")
	err := tr.todoService.UpdateTodo(todoRequest, todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error updating todo!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Todo updated successfully.",
	})
}

func (tr *TodoRoute) DeleteTodo(c *gin.Context) {
	todoId := c.GetInt64("todoId")
	err := tr.todoService.DeleteTodo(todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error deleteing todo!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Todo deleted successfully.",
	})
}
