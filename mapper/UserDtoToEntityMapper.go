package mapper

import (
	"golang-starter-data/domain/dto"
	"golang-starter-data/domain/entity"
)

func MapToEntity(dto *dto.UserDto) (e *entity.User) {
	e = entity.NewUserEntity()
	e.Name = dto.Name
	e.Age = dto.Age
	e.Lastname = dto.Lastname
	return
}
