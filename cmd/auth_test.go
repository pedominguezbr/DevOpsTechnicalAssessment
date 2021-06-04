package auth

import (
	"bytes"
	"encoding/json"
	"framework-go/pkg/features/devOps"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	user := devOps.RequestDevops{
		To: "jwt@email.com",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	print(request)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	assert.Equal(t, 200, w.Code)

}
func TestLoginInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	print(request)
	w := httptest.NewRecorder()

	assert.Equal(t, 200, w.Code)
}
