package commands

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/bryanl/doit-server"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestAuthCommand(t *testing.T) {
	cmd := Auth()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "login")
}

func TestAuth_retrieveCredentials(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac := doitserver.AuthCredentials{ID: "abc", CS: "def"}
		err := json.NewEncoder(w).Encode(&ac)
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	dsa := newDoitServerAuth()
	dsa.url = ts.URL
	ac, err := dsa.retrieveAuthCredentials()
	if err != nil {
		t.Fatal(err)
	}

	expectedAC := &doitserver.AuthCredentials{ID: "abc", CS: "def"}
	if got, want := ac, expectedAC; !reflect.DeepEqual(got, expectedAC) {
		t.Fatalf("retrieveAuthCredentials() = %#v; got %#v", got, want)
	}
}

func TestAuth_createAuthURL(t *testing.T) {
	dsa := newDoitServerAuth()
	dsa.url = "http://example.com"

	ac := &doitserver.AuthCredentials{
		ID: "id",
		CS: "cs",
	}
	u, err := dsa.createAuthURL(ac)
	if err != nil {
		t.Fatal(err)
	}

	newU, err := url.Parse(u)

	if got, want := newU.Host, "example.com"; got != want {
		t.Fatalf("createAuthURL() Host = %q, want = %q", got, want)
	}

	if got, want := newU.Path, "/auth/digitalocean"; got != want {
		t.Fatalf("createAuthURL() Path = %q, want = %q", got, want)
	}

	q := newU.Query()

	if got, want := q.Get("cs"), "cs"; got != want {
		t.Fatalf("createAuthURL() cs param = %q, want = %q", got, want)
	}

	if got, want := q.Get("id"), "id"; got != want {
		t.Fatalf("createAuthURL() id param = %q, want = %q", got, want)
	}
}

func TestAuth_createAuthURL_with_keypairs(t *testing.T) {
	dsa := newDoitServerAuth()
	dsa.url = "http://example.com"
	ac := &doitserver.AuthCredentials{
		ID: "id",
		CS: "cs",
	}
	u, err := dsa.createAuthURL(ac, keyPair{k: "foo", v: "bar"})
	if err != nil {
		t.Fatal(err)
	}

	newU, err := url.Parse(u)

	if got, want := newU.Host, "example.com"; got != want {
		t.Fatalf("createAuthURL() Host = %q, want = %q", got, want)
	}

	if got, want := newU.Path, "/auth/digitalocean"; got != want {
		t.Fatalf("createAuthURL() Path = %q, want = %q", got, want)
	}

	q := newU.Query()

	if got, want := q.Get("cs"), "cs"; got != want {
		t.Fatalf("createAuthURL() cs param = %q, want = %q", got, want)
	}

	if got, want := q.Get("id"), "id"; got != want {
		t.Fatalf("createAuthURL() id param = %q, want = %q", got, want)
	}

	if got, want := q.Get("foo"), "bar"; got != want {
		t.Fatalf("createAuthURL() foo param = %q, want = %q", got, want)
	}
}

func TestAuth_initAuthCLI(t *testing.T) {
	dsa := newDoitServerAuth()

	ac := &doitserver.AuthCredentials{
		ID: "id",
		CS: "cs",
	}

	go func() {
		s, err := dsa.initAuthCLI(ac)
		if err != nil {
			t.Fatalf("initAuthCLI() unexpected error: %v", err)
		}

		if got, want := s, "token"; got != want {
			t.Fatalf("initAuthCLI() = %q; want = %q", got, want)
		}
	}()

	os.Stdin.Write([]byte("token\n"))
}

func TestAuth_initAuth(t *testing.T) {
	dsa := newDoitServerAuth()
	dsa.url = "http://example.com"
	dsa.browserOpen = func(u string) error {
		return nil
	}
	dsa.monitorAuth = func(u string, ac *doitserver.AuthCredentials) (*doitserver.TokenResponse, error) {
		return &doitserver.TokenResponse{
			AccessToken: "access-token",
		}, nil
	}
	ac := &doitserver.AuthCredentials{
		ID: "id",
		CS: "cs",
	}

	token, err := dsa.initAuth(ac)
	if err != nil && err != UnknownTerminalErr {
		t.Fatalf("initAuth() unexpected error: %v", err)
	} else if err == UnknownTerminalErr {
		return
	}

	if got, want := token, "access-token"; got != want {
		t.Fatalf("initAuth() token = %q; want = %q", got, want)
	}
}

func TestAuth_monitorAuthWS(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("websocket upgrade error: %v", err)
		}

		var readAC doitserver.AuthCredentials
		err = ws.ReadJSON(&readAC)
		if err != nil {
			t.Fatalf("read auth credentials error: %v", err)
		}

		if got, want := readAC.CS, "cs"; got != want {
			t.Fatalf("server AuthCredentials CS = %q; want = %q", got, want)
		}

		if got, want := readAC.ID, "id"; got != want {
			t.Fatalf("server AuthCredentials ID = %q; want = %q", got, want)
		}

		writeTR := doitserver.TokenResponse{
			AccessToken: "access-token",
		}

		err = ws.WriteJSON(&writeTR)
		if err != nil {
			t.Fatalf("write token respose error: %v", err)
		}

		err = ws.Close()
		if err != nil {
			t.Fatalf("websock close: %v", err)
		}
	}))

	ac := &doitserver.AuthCredentials{
		ID: "id",
		CS: "cs",
	}

	tr, err := monitorAuthWS(ts.URL, ac)
	if err != nil {
		t.Fatalf("monitorAuthWS() unexpected error: %v", err)
	}

	if got, want := tr.AccessToken, "access-token"; got != want {
		t.Fatalf("monitorAuthWS() AccessToken = %q; want = %q", got, want)
	}
}
