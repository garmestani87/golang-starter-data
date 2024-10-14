package dto

type UserDto struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

func NewUserDto() *UserDto {
	return &UserDto{}
}
