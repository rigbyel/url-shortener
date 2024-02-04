package tests

import (
	"net/http"
	"net/url"
	"path"
	"testing"
	"url-shortener/internal/http-server/handlers/url/save"
	"url-shortener/internal/lib/random"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
)

const (
	host = "localhost:8081"
)



func TestURLShortener_SimplePath(t *testing.T) {
	u := url.URL {
		Scheme: "http",
		Host: host,
	}
	e := httpexpect.Default(t, u.String())

	e.POST("/url").
		WithJSON(save.Request{
			URL: gofakeit.URL(),
			Alias: random.NewRandomString(10),
		}).
		WithBasicAuth("bigadmin", "mypass").
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("alias")
}

func TestURLShortener_SaveRedirectRemove(t *testing.T) {
	testCases := []struct {
		name string
		url string
		alias string
		error string
	} {
		{
			name: "Valid URL",
			url: gofakeit.URL(),
			alias: gofakeit.Word() + gofakeit.Word(),
		},
		{
			name: "Invalid URL",
			url: "mksll",
			alias: gofakeit.Word(),
			error: "field URL is not a valid URL",
		},
		{
			name: "Empty alias",
			url: gofakeit.URL(),
			alias: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := url.URL {
				Scheme: "http",
				Host: host,
			}

			e := httpexpect.Default(t, u.String())

			resp := e.POST("/url").
			WithJSON(save.Request{
				URL: tc.url,
				Alias: tc.alias,
			}).
			WithBasicAuth("bigadmin", "mypass").
			Expect().
			Status(http.StatusOK).
			JSON().Object()

			// testing save

			if tc.error != "" {
				resp.NotContainsKey("alias")

				resp.Value("error").String().IsEqual(tc.error)

				return
			}

			alias := tc.alias

			if tc.alias != "" {
				resp.Value("alias").String().IsEqual(alias)
			} else {
				resp.Value("alias").String().NotEmpty()
				alias = resp.Value("alias").String().Raw()
			}

			// TODO: test redirect


			// testing remove

			reqDel := e.DELETE("/"+path.Join("url", alias)). 
				WithBasicAuth("bigadmin", "mypass").
				Expect().
				Status(http.StatusOK).
				JSON().Object()

			reqDel.Value("status").String().IsEqual("OK")
		})

	}
}
