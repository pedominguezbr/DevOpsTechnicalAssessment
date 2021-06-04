package auth

import (
	"bytes"
	"encoding/json"
	"framework-go/pkg/features/devOps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	req := &http.Request{Method: "POST"}
	req.URL, _ = url.Parse("http://www.google.com/search?q=foo&q=bar")
	if q := req.FormValue("q"); q != "foo" {
		t.Errorf(`req.FormValue("q") = %q, want "foo"`, q)
	}
}

func TestDevops(t *testing.T) {
	user := devOps.RequestDevops{
		To: "jwt@email.com",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/Devops", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	print(request)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	assert.Equal(t, 200, w.Code)

}
func TestDevopsInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/Devops", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	print(request)
	w := httptest.NewRecorder()

	assert.Equal(t, 200, w.Code)
}
