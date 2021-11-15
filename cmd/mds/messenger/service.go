package messenger

import (
	"fmt"
	"mds/internal"
	"mds/internal/message"
)

type Service interface {
	Connect(name string) error
	Disconnect(name string) error
	SendIdentity(name string) error
	SendList(name string) error
	SendRelay(name string, message interface{}) error
	GetMessage(name string) (string, error)
}

type service struct {
	repo   internal.UserRepository
	rabbit internal.RabbitMQ
}

func NewService(repo internal.UserRepository, rabbit internal.RabbitMQ) Service {
	return service{
		repo:   repo,
		rabbit: rabbit,
	}
}

func (service service) Connect(name string) error {
	id, err := service.repo.Create(message.User{Name: name})
	if err != nil {
		return err
	}

	if id == 0 {
		return fmt.Errorf("no id assoicated with user")
	}

	err = service.rabbit.CreateQueue(fmt.Sprintf("%d", id))
	if err != nil {
		return err
	}

	return nil
}

func (service service) Disconnect(name string) error {

	user, err := service.repo.Get(name)
	if err != nil {
		return err
	}

	err = service.repo.Delete(user)
	if err != nil {
		return err
	}

	err = service.rabbit.DeleteQueue(fmt.Sprintf("%d", user.UserID))
	if err != nil {
		return err
	}

	return nil
}

func (service service) SendIdentity(name string) error {
	user, err := service.repo.Get(name)
	if err != nil {
		return err
	}

	if user.UserID == 0 {
		return fmt.Errorf("no id assoicated with user")
	}

	err = service.rabbit.Publish(fmt.Sprintf("%d", user.UserID), message.IndentityResponse{UserID: user.UserID})
	if err != nil {
		return err
	}

	return nil
}

func (service service) SendList(name string) error {
	users, err := service.repo.GetAll()
	if err != nil {
		return err
	}

	var otherUsers []uint64
	var myUserId uint64
	for _, user := range users {
		if user.Name != name {
			otherUsers = append(otherUsers, user.UserID)
		} else {
			myUserId = user.UserID
		}
	}

	err = service.rabbit.Publish(fmt.Sprintf("%d", myUserId), message.ListResponse{Users: otherUsers})
	if err != nil {
		return err
	}

	return nil
}

func (service service) SendRelay(name string, messageBody interface{}) error {
	users, err := service.repo.GetAll()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name != name {
			err = service.rabbit.Publish(fmt.Sprintf("%d", user.UserID), message.RelayResponse{Message: messageBody})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (service service) GetMessage(name string) (string, error) {

	user, err := service.repo.Get(name)
	if err != nil {
		return "", err
	}

	msg, err := service.rabbit.Consume(fmt.Sprintf("%d", user.UserID))
	if err != nil {
		return "", err
	}

	return msg, nil
}
