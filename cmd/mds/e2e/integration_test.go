package e2e

import (
	"mds/cmd/mds/e2e/helpers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Client = &http.Client{}
var Url = "http://localhost:8080"

func TestConnect(t *testing.T) {
	res := helpers.Connect(t, "connectTest")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "connectTest")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	user, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	if user.Name != "connectTest" {
		t.Errorf("expected %s, got %s", "user", user.Name)
	}
}

func TestDisconnect(t *testing.T) {
	res := helpers.Connect(t, "disconnectTest")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "disconnectTest")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	user, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	if user.Name != "disconnectTest" {
		t.Errorf("expected %s, got %s", "user", user.Name)
	}

	res = helpers.Disconnect(t, "disconnectTest")
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestIdentity(t *testing.T) {
	res := helpers.Connect(t, "identityTest")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "identityTest")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	user, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	if user.Name != "identityTest" {
		t.Errorf("expected %s, got %s", "user", user.Name)
	}

	res = helpers.SendIdentity(t, user.Name)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestList(t *testing.T) {
	res := helpers.Connect(t, "listTest")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "listTest")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	user, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	if user.Name != "listTest" {
		t.Errorf("expected %s, got %s", "user", user.Name)
	}

	res = helpers.SendList(t, user.Name)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestRelay(t *testing.T) {

	// create sender
	res := helpers.Connect(t, "sender")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "sender")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	sender, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	//create receiver
	res = helpers.Connect(t, "receiver")
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	res = helpers.GetUser(t, "receiver")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	receiver, err := helpers.ParseUser(res)
	if err != nil {
		t.Errorf("unable to get user from response due to %s", err)
	}

	if sender.Name != "sender" {
		t.Errorf("expected %s, got %s", "sender", sender.Name)
	}

	if receiver.Name != "receiver" {
		t.Errorf("expected %s, got %s", "receiver", receiver.Name)
	}

	//send relay
	res = helpers.SendRelay(t, sender.Name,
		helpers.RelayRequest{
			Sender:     sender.Name,
			Recipients: []string{receiver.Name},
			Message:    "test",
		},
	)
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
