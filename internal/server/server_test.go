package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/mishalalajmi/mimic/internal/mock"
	"github.com/mishalalajmi/mimic/internal/project"
)

func TestServer_StartAndServe(t *testing.T) {
	port := freePort()
	s := New(fmt.Sprintf("127.0.0.1:%d", port), testProject())

	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	waitForServer(t, port)

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/users", port))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}

	if string(body) != `{"users":[]}` {
		t.Fatalf("expected body %q, got %q", `{"users":[]}`, string(body))
	}

	shutdownServer(t, s, errCh)
}

func TestServer_UnmatchedRoute(t *testing.T) {
	port := freePort()
	s := New(fmt.Sprintf("127.0.0.1:%d", port), testProject())

	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	waitForServer(t, port)

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/unknown", port))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	shutdownServer(t, s, errCh)
}

func TestServer_ShutdownBeforeStart(t *testing.T) {
	s := New(":8080", testProject())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		t.Fatalf("expected shutdown to succeed, got %v", err)
	}
}

func testProject() *project.Project {
	return &project.Project{
		Name: "test-project",
		Definition: []mock.Definition{
			{
				Name: "users",
				Routes: []mock.Route{
					{
						Name: "get-users",
						Request: mock.Request{
							Method: http.MethodGet,
							Path:   "/users",
						},
						Response: mock.Response{
							Status: http.StatusOK,
							Headers: map[string]string{
								"Content-Type": "application/json",
							},
							Body: map[string]any{
								"users": []any{},
							},
						},
					},
				},
			},
		},
	}
}

func freePort() int {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port
}

func waitForServer(t *testing.T, port int) {
	t.Helper()

	address := fmt.Sprintf("127.0.0.1:%d", port)
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", address, 50*time.Millisecond)
		if err == nil {
			conn.Close()
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("server did not start on %s", address)
}

func shutdownServer(t *testing.T, s *Server, errCh chan error) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown failed: %v", err)
	}

	select {
	case err := <-errCh:
		if !errors.Is(err, ErrServerClosed) {
			t.Fatalf("expected ErrServerClosed, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("server did not stop")
	}
}
