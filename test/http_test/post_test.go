package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestMockPost(t *testing.T) {
	url := "http://example.com"
	path := "/add"
	str := `{"name": "tony"}`
	gock.New(url).Post(path).JSON(map[string]string{"name": "tony"}).Reply(200).
		JSON(map[string]string{
			"id":   "123",
			"name": "tony",
		})

	body := strings.NewReader(str)
	res, err := http.Post(url+path, "application/json", body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, res.StatusCode)
	jsonStr, _ := ioutil.ReadAll(res.Body)
	var user map[string]string
	json.Unmarshal(jsonStr, &user)
	assert.Equal(t, "tony", user["name"])
	assert.Equal(t, "123", user["id"])
	assert.Equal(t, gock.IsDone(), true)
}
