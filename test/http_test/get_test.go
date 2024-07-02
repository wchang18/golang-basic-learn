package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockGet(t *testing.T) {
	url := "http://example.com"
	path := "/user"
	gock.New(url).Get(path).Reply(200).
		JSON(map[string]string{
			"name": "tony",
		})

	res, err := http.Get(url + path)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, res.StatusCode)
	jsonStr, _ := ioutil.ReadAll(res.Body)
	var user map[string]string
	json.Unmarshal(jsonStr, &user)
	assert.Equal(t, "tony", user["name"])
	assert.Equal(t, gock.IsDone(), true)
}
