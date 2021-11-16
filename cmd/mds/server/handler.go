package server

import (
	"fmt"
	"mds/cmd/mds/messenger"
	"net/http"

	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

type Handler struct {
	service     messenger.Service
	stopChannel map[string]chan (bool)
}

func NewHandler(service messenger.Service) Handler {
	return Handler{
		service:     service,
		stopChannel: make(map[string]chan bool),
	}
}

func StartQueue(stop chan (bool), msgs <-chan amqp.Delivery, ID string, name string) {
	go func() {
		for d := range msgs {
			fmt.Printf("recipient: %s \nqueue id: %s \nmessage: %s \n", name, ID, d.Body)
		}
		<-stop
		fmt.Printf("sign out user: %s has signed out \n", name)
	}()
}

func (h Handler) Connect(c echo.Context) error {
	name := c.QueryParam("name")

	h.stopChannel[name] = make(chan (bool))

	err := h.service.Connect(name, h.stopChannel[name], StartQueue)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	return c.JSONPretty(http.StatusCreated, NewSuccessfulResponse(), " ")

}

func (h Handler) Disconnect(c echo.Context) error {
	name := c.QueryParam("name")
	err := h.service.Disconnect(name)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	h.stopChannel[name] <- true

	return c.JSONPretty(http.StatusOK, NewSuccessfulResponse(), " ")

}

func (h Handler) SendIdentiyMesasge(c echo.Context) error {
	user := c.QueryParam("name")

	err := h.service.SendIdentity(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	return c.JSONPretty(http.StatusOK, NewSuccessfulResponse(), " ")
}

func (h Handler) SendListMesasge(c echo.Context) error {
	user := c.QueryParam("name")

	if err := h.service.SendList(user); err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	return c.JSONPretty(http.StatusOK, NewSuccessfulResponse(), " ")
}

func (h Handler) SendRelay(c echo.Context) error {
	request := new(RelayRequest)
	if err := c.Bind(request); err != nil {
		return c.JSONPretty(http.StatusBadRequest, err, " ")
	}

	err := h.service.SendRelay(request.Sender, request.Recipients, request.Message)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	return c.JSONPretty(http.StatusOK, NewSuccessfulResponse(), " ")

}

// integration test handlers
func (h Handler) GetUser(c echo.Context) error {
	name := c.QueryParam("name")

	user, err := h.service.GetUser(name)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, err, " ")
	}

	return c.JSONPretty(http.StatusOK, user, " ")
}
