package test

import (
	"bytes"
	model "cms/src/models"
	gin "cms/src/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

type createProduct struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity uint8   `json:"quantity"`
}

type sessionResponse struct {
	User  model.User
	Token string
}

// Should be create User and return response 201.
func TestProduct(t *testing.T) {
	var r gin.Routes

	userEmail := faker.Email()

	requestBody := createProduct{"Nice Product", 99.99, 2}
	reqBodyBytes := new(bytes.Buffer)

	json.NewEncoder(reqBodyBytes).Encode(requestBody)

	server := r.StartGin()
	ts := httptest.NewServer(server)

	user := SignUpCredentials{userEmail, "SignIn", "123456"}
	userBytes := new(bytes.Buffer)

	json.NewEncoder(userBytes).Encode(user)

	defer ts.Close()

	fmt.Println()

	resp, _ := http.Post(ts.URL+"/users", "application/json", userBytes)

	assert.Equal(t, resp.StatusCode, 201, "They should be equal")

	loginUser := SignInCredentials{userEmail, "123456"}

	loginBytes := new(bytes.Buffer)

	json.NewEncoder(loginBytes).Encode(loginUser)

	resp, _ = http.Post(ts.URL+"/sessions", "application/json", loginBytes)

	assert.Equal(t, resp.StatusCode, 200, "They should be equal")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(body)

	response := sessionResponse{}

	json.Unmarshal([]byte(bodyString), &response)

	req, _ := http.NewRequest("POST", ts.URL+"/products", reqBodyBytes)

	req.Header.Set("Authorization", "Bearer "+response.Token)

	client := &http.Client{}

	resp, _ = client.Do(req)

	assert.Equal(t, resp.StatusCode, 201, "They should be equal")

}
