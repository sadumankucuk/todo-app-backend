package model

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

type TodoRequest struct {
	Task string `json:"task"`
}

type TodoResponse []Todo
