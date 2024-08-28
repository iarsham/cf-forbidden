package cf_forbidden

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"net/http"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	var err error
	client, err = New()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestNewClient(t *testing.T) {
	if client == nil {
		t.Error("client is nil")
	}
}

func TestHTTPGet(t *testing.T) {
	expectedResp := cycletls.Response{
		Status: http.StatusOK,
	}
	resp, err := client.Get("https://indeed.com", nil)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != expectedResp.Status {
		t.Errorf("expected %d, got %d", expectedResp.Status, resp.Status)
	}
}

func TestHTTPPost(t *testing.T) {
	expectedResp := cycletls.Response{
		Status: http.StatusOK,
	}
	resp, err := client.Post("https://indeed.com", nil, "test")
	if err != nil {
		t.Error(err)
	}
	if resp.Status != expectedResp.Status {
		t.Errorf("expected %d, got %d", expectedResp.Status, resp.Status)
	}
}
