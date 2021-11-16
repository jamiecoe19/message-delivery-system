package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var Client = &http.Client{}
var Url = "http://localhost:8080"

func Connect(t *testing.T, name string) *http.Response {
	req, _ := http.NewRequest("GET", Url+"/connect?name="+name, nil)
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func Disconnect(t *testing.T, name string) *http.Response {
	req, _ := http.NewRequest("DELETE", Url+"/disconnect?name="+name, nil)
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func GetUser(t *testing.T, name string) *http.Response {
	req, _ := http.NewRequest("GET", Url+"/get?name="+name, nil)
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func SendIdentity(t *testing.T, name string) *http.Response {
	req, _ := http.NewRequest("GET", Url+"/identity?name="+name, nil)
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func SendList(t *testing.T, name string) *http.Response {
	req, _ := http.NewRequest("GET", Url+"/list?name="+name, nil)
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func SendRelay(t *testing.T, name string, body RelayRequest) *http.Response {
	b := strings.NewReader(string(marshalReplayRequestJson(t, body)))
	req, _ := http.NewRequest("POST", Url+"/relay", b)
	req.Header.Add("Content-Type", "application/json")
	res, err := Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func marshalReplayRequestJson(t *testing.T, body RelayRequest) []byte {
	jsonString, err := json.Marshal(&body)
	if err != nil {
		t.Fatal(err)
	}
	return jsonString
}

func ParseUser(res *http.Response) (User, error) {
	body, err := ReadResponse(res)
	if err != nil {
		return User{}, err
	}
	var result User
	if err = json.Unmarshal(body, &result); err != nil {
		return User{}, err
	}

	return result, nil

}

func ReadResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
