package test

import (
	"bytes"
	gin "cms/src/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

type SignUpCredentials struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Should be create User and return response 201.
func TestSignUp(t *testing.T) {
	var r gin.Routes

	userEmail := faker.Email()

	requestBody := SignUpCredentials{userEmail, "John Doe", "123456"}
	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(requestBody)

	server := r.StartGin()
	ts := httptest.NewServer(server)

	defer ts.Close()

	resp, _ := http.Post(ts.URL+"/users", "application/json", reqBodyBytes)

	assert.Equal(t, 201, resp.StatusCode, "They should be equal")

}
