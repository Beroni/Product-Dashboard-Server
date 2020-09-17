package test

import (
	"bytes"
	gin "cms/src/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SignInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Should be create User and return response 201.
func TestSignIn(t *testing.T) {
	var r gin.Routes

	user := SignUpCredentials{"signinaccount@gmail.com", "SignIn", "123456"}
	userBytes := new(bytes.Buffer)

	json.NewEncoder(userBytes).Encode(user)

	server := r.StartGin()
	ts := httptest.NewServer(server)

	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/users", "application/json", userBytes)

	assert.Equal(t, resp.StatusCode, 201, "They should be equal")

	loginUser := SignInCredentials{"signinaccount@gmail.com", "123456"}

	loginBytes := new(bytes.Buffer)

	json.NewEncoder(loginBytes).Encode(loginUser)

	resp, _ = http.Post(ts.URL+"/sessions", "application/json", loginBytes)

	assert.Equal(t, resp.StatusCode, 200, "They should be equal")

}
