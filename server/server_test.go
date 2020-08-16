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
	mockServer := MockServer{}
	NewServer(
		&mockServer.logger,
		"foo",
		"./mocks",
		mockServer.serveFunc(),
	).Start()
	defer mockServer.server.Close()

	httpClient := &http.Client{}

	resp, err := httpClient.Get(mockServer.addr + url)

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
	got, err := getFromURL("/ping")
	want := "pong"

	if err != nil || got != want {
		t.Fatalf(
			"Response body not matching. Err: %v, Got: %v, Want: %v",
			err,
			got,
			want,
		)
	}
}

func TestVersionRoute(t *testing.T) {
	got, err := getFromURL("/version")
	want := "0.0"

	if err != nil || !strings.Contains(got, want) {
		t.Fatalf(
			"Response body not matching. Err: %v, Got: %v, Should contains: %v",
			err,
			got,
			want,
		)
	}
}

func TestSimulationIDRoute(t *testing.T) {
	got, err := getFromURL("/simulations/foobar")
	want := "ID: foobar\n"

	if err != nil || got != want {
		t.Fatalf(
			"Response body not matching. Err: %v, Got: %v, Want: %v",
			err,
			got,
			want,
		)
	}
}

func TestStaticFile(t *testing.T) {
	got, err := getFromURL("/static_page.txt")
	want := "static file here\n"

	if err != nil || got != want {
		t.Fatalf(
			"Response body not matching. Err: %v, Got: %v, Want: %v",
			err,
			got,
			want,
		)
	}
}
