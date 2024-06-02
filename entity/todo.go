package entity

type TodoList struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	ListId int64 `json:"listId"`
}

type TodoItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int64 `json:"id"`
	ListId int64 `json:"listId"`
	ItemId int64 `json:"itemId"`
}