package activity_group

type ActivityGroupRequest struct {
	Title string `mapstructure:"title" validate:"required"`
	Email string `mapstructure:"email" validate:"required,email"`
}
