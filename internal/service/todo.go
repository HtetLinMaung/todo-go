package service

import (
	"database/sql"
	"fmt"

	"github.com/HtetLinMaung/todo/internal/model"
	"github.com/HtetLinMaung/todo/internal/utils"
)

type TodoService struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) AddTodo(todoRequest *model.TodoRequest, creatorId int64) (int64, error) {
	var todoId int64

	stmt, err := s.db.Prepare("insert into todos (label, description, is_done, creator_id) values ($1, $2, $3, $4) returning todo_id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(todoRequest.Label, todoRequest.Description, todoRequest.IsDone, creatorId).Scan(&todoId)
	if err != nil {
		return 0, err
	}

	return todoId, nil
}

func (s *TodoService) GetTodos(search string, page uint, perPage uint, userId int64, role string) (*utils.PaginationResult[model.Todo], error) {
	var total uint = 0
	var args []interface{}
	baseQuery := "from todos t inner join users u on u.user_id = t.creator_id where t.deleted_at is null"

	if role != "admin" {
		args = append(args, userId)
		baseQuery = fmt.Sprintf("%s and creator_id = $%d", baseQuery, len(args))
	}

	result := utils.GeneratePaginationQuery(&utils.PaginationOptions{
		SelectColumns: "t.todo_id, t.label, t.description, t.is_done, t.created_at, u.name",
		BaseQuery:     baseQuery,
		SearchColumns: []string{"t.todo_id::text", "t.label", "t.description", "u.name"},
		Search:        search,
		OrderOptions:  "t.created_at desc",
		Page:          page,
		PerPage:       perPage,
	})

	stmt, err := s.db.Prepare(result.CountQuery)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(args...).Scan(&total)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := s.db.Query(result.Query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]model.Todo, 0)
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.TodoID, &todo.Label, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.CreatorName); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var PageCounts uint
	if perPage != 0 {
		PageCounts = (total + (perPage - 1)) / perPage
	}

	return &utils.PaginationResult[model.Todo]{
		Data:       todos,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		PageCounts: PageCounts,
	}, nil
}

func (s *TodoService) GetTodoById(todoId int64, creatorId int64, role string) (*model.Todo, error) {
	var todo model.Todo

	var args []interface{}
	args = append(args, todoId)
	query := "select t.todo_id, t.label, t.description, t.is_done, t.created_at, u.name from todos t inner join users u on u.user_id = t.creator_id where t.deleted_at is null and t.todo_id = $1"
	if role != "admin" {
		args = append(args, creatorId)
		query = fmt.Sprintf("%s and creator_id = $2", query)
	}

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&todo.TodoID, &todo.Label, &todo.Description, &todo.IsDone, &todo.CreatedAt, &todo.CreatorName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) UpdateTodo(todoRequest *model.TodoRequest, todoId int64) error {
	stmt, err := s.db.Prepare("update todos set label = $1, description = $2, is_done = $3 where todo_id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todoRequest.Label, todoRequest.Description, todoRequest.IsDone, todoId)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) DeleteTodo(todoId int64) error {
	stmt, err := s.db.Prepare("update todos set deleted_at = CURRENT_TIMESTAMP where todo_id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todoId)
	if err != nil {
		return err
	}
	return nil
}
