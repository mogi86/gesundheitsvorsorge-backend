package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase/usecase_mock"
)

func TestController_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()
	m := usecase_mock.NewMockInterface(ctrl)
	m.EXPECT().Sample().Return("test response!")

	controller := NewController(m)

	server := httptest.NewServer(controller)
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("test request failed. %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid http status code. %v", res.StatusCode)
	}

	responseMessage, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("read response failed. %v", err)
	}

	if string(responseMessage) != "test response!" {
		t.Errorf("response message is wrong. %s", responseMessage)
	}
}
