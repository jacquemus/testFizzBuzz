package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_integration_end_to_end(t *testing.T) {
	router := router()
	ctx := context.TODO()

	DefaultData := `{"data":["1","2","3","4","5","6","fizz","8","buzz","10","11","12","13","fizz","15","16","17","buzz","19","20","fizz","22","23","24","25","26","buzz","fizz","29","30","31","32","33","34","fizz","buzz","37","38","39","40","41","fizz","43","44","buzz","46","47","48","fizz","50","51","52","53","buzz","55","fizz","57","58","59","60","61","62","fizzbuzz","64","65","66","67","68","69","fizz","71","buzz","73","74","75","76","fizz","78","79","80","buzz","82","83","fizz","85","86","87","88","89","buzz","fizz","92","93","94","95","96","97","fizz","buzz","100","101","102","103","104","fizz","106","107","buzz","109","110","111","fizz","113","114","115","116","buzz","118","fizz","120","121","122","123","124","125","fizzbuzz","127","128","129","130","131","132","fizz","134","buzz","136","137","138","139","fizz","141","142","143","buzz","145","146","fizz","148","149","150","151","152","buzz","fizz","155","156","157","158","159","160","fizz","buzz","163","164","165","166","167","fizz","169","170","buzz","172","173","174","fizz","176","177","178","179","buzz","181","fizz","183","184","185","186","187","188","fizzbuzz","190","191","192","193","194","195","fizz","197","buzz","199","200"],"message":"success!"}`
	Limit10Data := `{"data":["1","2","first","4","second","first","7","8","first","second"],"message":"success!"}`

	t.Run("When calling the fizzbuzz Should return the correct default response", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(ctx, "GET", "/fizzbuzz", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, DefaultData, w.Body.String())
	})

	t.Run("When calling the fizzbuzz and setting the params Should return the correct response", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(ctx, "GET", "/fizzbuzz", nil)
		q := req.URL.Query()
		q.Add("limit", "10")
		q.Add("modulo1", "3")
		q.Add("modulo2", "5")
		q.Add("word1", "first")
		q.Add("word2", "second")

		req.URL.RawQuery = q.Encode()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, Limit10Data, w.Body.String())
	})
}

func Test_integration_error(t *testing.T) {
	router := router()
	ctx := context.TODO()

	t.Run("When calling the fizzbuzz and setting the params incorrectly a params Should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(ctx, "GET", "/fizzbuzz", nil)
		q := req.URL.Query()
		q.Add("limit", "testValue")

		req.URL.RawQuery = q.Encode()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("When calling the fizzbuzz and setting the params with an invalid a params Should return an error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(ctx, "GET", "/fizzbuzz", nil)
		q := req.URL.Query()
		q.Add("limit", "0")

		req.URL.RawQuery = q.Encode()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
