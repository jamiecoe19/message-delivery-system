package internal

import (
	"mds/internal/message"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAll() (message.Users, error)
	Get(string) (message.User, error)
	Create(message.User) (uint64, error)
	Delete(message.User) error
}

type userRepository struct {
	db *gorm.DB
}

type UserDTOs []UserDTO

func (UserDTO) TableName() string {
	return "users"
}

type UserDTO struct {
	UserID uint64 `gorm:"user_id"`
	Name   string `gorm:"name"`
}

func (dto UserDTO) toDomain() message.User {
	return message.User{
		UserID: dto.UserID,
		Name:   dto.Name,
	}
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (repo userRepository) GetAll() (message.Users, error) {
	var dtos UserDTOs

	err := repo.db.Find(&dtos).Error
	if err != nil {
		return message.Users{}, err
	}

	var users message.Users
	for _, dto := range dtos {
		users = append(users, dto.toDomain())
	}
	return users, nil
}

func (repo userRepository) Get(name string) (message.User, error) {
	var dto UserDTO

	err := repo.db.Where("name = ?", name).First(&dto).Error
	if err != nil {
		return message.User{}, err
	}
	return dto.toDomain(), nil
}

func (repo userRepository) Create(user message.User) (uint64, error) {

	ID := GenerateID()

	dto := message.User{
		UserID: ID,
		Name:   user.Name,
	}

	err := repo.db.Create(&dto).Error
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (repo userRepository) Delete(user message.User) error {
	userDTO := UserDTO{
		UserID: user.UserID,
		Name:   user.Name,
	}

	err := repo.db.Where("user_id = ?", userDTO.UserID).Delete(&userDTO).Error
	if err != nil {
		return err
	}

	return nil
}
