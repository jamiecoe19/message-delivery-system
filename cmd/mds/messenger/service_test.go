package messenger_test

import (
	"mds/cmd/mds/messenger"
	"mds/internal/message"
	mockInternal "mds/internal/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/streadway/amqp"
)

func TestConnect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ID := uint64(1234567890)
	IDString := "1234567890"
	name := "test"

	startQueueCalled := false

	startQueue := func(chan (bool), <-chan amqp.Delivery, string, string) {
		startQueueCalled = true
	}

	var delivery <-chan amqp.Delivery

	mockRepo := mockInternal.NewMockUserRepository(mockCtrl)
	mockRabbit := mockInternal.NewMockRabbitMQ(mockCtrl)
	mockRepo.EXPECT().Create(message.User{Name: name}).Return(ID, nil)
	mockRabbit.EXPECT().CreateQueue(IDString).Return(nil)
	mockRabbit.EXPECT().Consume(IDString).Return(delivery, nil)

	srv := messenger.NewService(mockRepo, mockRabbit)

	err := srv.Connect(name, make(chan bool), startQueue)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if !startQueueCalled {
		t.Errorf("startQueue was never used")
	}
}

func TestDisconnect(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ID := uint64(1234567890)

	IDString := "1234567890"
	name := "test"

	user := message.User{ID, name}

	mockRepo := mockInternal.NewMockUserRepository(mockCtrl)
	mockRabbit := mockInternal.NewMockRabbitMQ(mockCtrl)
	mockRepo.EXPECT().Get(name).Return(user, nil)
	mockRepo.EXPECT().Delete(user).Return(nil)
	mockRabbit.EXPECT().DeleteQueue(IDString).Return(nil)

	srv := messenger.NewService(mockRepo, mockRabbit)

	err := srv.Disconnect(name)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestSendIdentity(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ID := uint64(1234567890)

	IDString := "1234567890"
	name := "test"

	user := message.User{ID, name}
	identityMessage := message.IndentityResponse{
		UserID: user.UserID,
	}

	mockRepo := mockInternal.NewMockUserRepository(mockCtrl)
	mockRabbit := mockInternal.NewMockRabbitMQ(mockCtrl)
	mockRepo.EXPECT().Get(name).Return(user, nil)
	mockRabbit.EXPECT().Publish(IDString, identityMessage).Return(nil)

	srv := messenger.NewService(mockRepo, mockRabbit)

	err := srv.SendIdentity(name)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestSendList(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userList := message.Users{
		message.User{
			UserID: 1223344,
			Name:   "test1",
		},
		message.User{
			UserID: 121241231231,
			Name:   "test1",
		},
		message.User{
			UserID: 29128309128,
			Name:   "myself",
		},
	}

	IDString := "29128309128"

	list := []uint64{1223344, 121241231231}
	listMessage := message.ListResponse{
		Users: list,
	}

	mockRepo := mockInternal.NewMockUserRepository(mockCtrl)
	mockRabbit := mockInternal.NewMockRabbitMQ(mockCtrl)
	mockRepo.EXPECT().GetAll().Return(userList, nil)
	mockRabbit.EXPECT().Publish(IDString, listMessage).Return(nil)

	srv := messenger.NewService(mockRepo, mockRabbit)

	err := srv.SendList("myself")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestSendRelay(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userList := message.Users{
		message.User{
			UserID: 1223344,
			Name:   "test1",
		},
		message.User{
			UserID: 121241231231,
			Name:   "test2",
		},
		message.User{
			UserID: 29128309128,
			Name:   "myself",
		},
	}

	user1IDString := "1223344"
	user2IDString := "121241231231"

	relayMessage := message.RelayResponse{
		Message: "a test message",
	}

	mockRepo := mockInternal.NewMockUserRepository(mockCtrl)
	mockRabbit := mockInternal.NewMockRabbitMQ(mockCtrl)
	mockRepo.EXPECT().GetAll().Return(userList, nil)
	mockRabbit.EXPECT().Publish(user1IDString, relayMessage).Return(nil)
	mockRabbit.EXPECT().Publish(user2IDString, relayMessage).Return(nil)

	srv := messenger.NewService(mockRepo, mockRabbit)

	err := srv.SendRelay("myself", []string{"test1", "test2"}, "a test message")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}
