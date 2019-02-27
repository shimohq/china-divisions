package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yalp/jsonpath"
)

type responseModel struct {
	Message interface{} `json:"message"`
}

func TestDistrictGetNoParams(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/district", nil)
	var res responseModel
	getRes(w, req, router, &res)
	out, err := jsonpath.Read(res.Message, "$[0].name")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "北京市", out)
}

func TestDistrictGetWithCodeProvinces(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/district?code=11", nil)
	var res responseModel
	getRes(w, req, router, &res)
	out, err := jsonpath.Read(res.Message, "$.name")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "北京市", out)
}

func TestDistrictGetWithCodeProvincesAndSub2(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/district?code=11&&subdistrict=2", nil)
	var res responseModel
	getRes(w, req, router, &res)
	out, err := jsonpath.Read(res.Message, "$.children[1].name")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "西城区", out)
}

func TestDistrictGetWithCodeCityAndSub3(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/district?code=110101&&subdistrict=3", nil)
	var res responseModel
	getRes(w, req, router, &res)
	out, err := jsonpath.Read(res.Message, "$.name")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "东城区", out)
}

func getRes(w *httptest.ResponseRecorder, req *http.Request, router *gin.Engine, res *responseModel) {
	router.ServeHTTP(w, req)
	result := w.Result()
	body, _ := ioutil.ReadAll(result.Body)
	json.Unmarshal(body, &res)
}
