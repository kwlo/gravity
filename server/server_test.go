package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"net/http/httptest"
	"testing"
)

type MockLogger struct{}

func (*MockLogger) Infof(template string, args ...interface{}) {}
func (*MockLogger) Sync() error {
	return nil
}

type MockServer struct {
	addr   string
	server *httptest.Server
	logger MockLogger
}

func (srv *MockServer) serveFunc() func(addr string, handler http.Handler) error {
	return func(addr string, handler http.Handler) error {
		srv.server = httptest.NewServer(handler)
		srv.addr = srv.server.URL

		return nil
	}
}

func getFromURL(url string) (string, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		return string(bodyBytes), nil
	}

	return "", fmt.Errorf("Status code not OK: %v", resp.StatusCode)
}

func TestPingRoute(t *testing.T) {
	mockServer := MockServer{}

	NewServer(&mockServer.logger, "foo", mockServer.serveFunc()).Start()

	got, err := getFromURL(mockServer.addr + "/ping")

	if err != nil {
		t.Fatal(err)
	}

	want := "pong"

	if got != want {
		t.Fatalf(
			"Response body not matching. Got: %v, Want: %v",
			got,
			want,
		)
	}
}

func TestVersionRoute(t *testing.T) {
	mockServer := MockServer{}

	NewServer(&mockServer.logger, "foo", mockServer.serveFunc()).Start()

	got, err := getFromURL(mockServer.addr + "/version")

	if err != nil {
		t.Fatal(err)
	}

	want := "0.0"

	if !strings.Contains(got, want) {
		t.Fatalf(
			"Response body not matching. Got: %v, Should contains: %v",
			got,
			want,
		)
	}
}

func TestSimulationIDRoute(t *testing.T) {
	mockServer := MockServer{}

	NewServer(&mockServer.logger, "foo", mockServer.serveFunc()).Start()

	got, err := getFromURL(mockServer.addr + "/simulations/foobar")

	if err != nil {
		t.Fatal(err)
	}

	want := "ID: foobar\n"

	if got != want {
		t.Fatalf(
			"Response body not matching. Got: %s, Want: %s",
			got,
			want,
		)
	}
}
