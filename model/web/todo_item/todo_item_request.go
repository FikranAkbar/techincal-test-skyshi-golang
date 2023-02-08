package todo_item

type TodoItemRequest struct {
	Title           *string `mapstructure:"title" json:"title"`
	ActivityGroupId *uint   `mapstructure:"activity_group_id" json:"activity_group_id"`
	IsActive        *bool   `mapstructure:"is_active" json:"is_active"`
}
