package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/rizwank123/myResturent/internal/dependency"
	"github.com/rizwank123/myResturent/internal/http/api"
	"github.com/rizwank123/myResturent/internal/http/transport"
	"github.com/rizwank123/myResturent/internal/pkg/config"
)

type EchoHandler func(c echo.Context) error
type TearDownSuite func(tb testing.TB)

func SetupSuite(tb testing.TB) (a *api.ResturnetApi, e *echo.Echo, td TearDownSuite) {
	opts := config.Options{
		ConfigSource: config.SourceEnv,
		ConfigFile:   "../../test.env",
	}

	cfg, err := dependency.NewConfig(opts)
	if err != nil {
		tb.Fatalf("failed to load config: %v", err)
	}

	db, err := dependency.NewDatabase(cfg)
	if err != nil {
		tb.Fatalf("failed to connect DB: %v", err)
	}

	e = echo.New()
	e.Validator = &transport.CustomValidator{Validator: validator.New()}

	a, err = dependency.NewResturnetApi(cfg, db)
	if err != nil {
		tb.Fatalf("failed to init API deps: %v", err)
	}
	td = func(tb testing.TB) {
		db.Close()
	}

	return a, e, td
}

// SendRequest sends a test request to the Echo handler
func SendRequest(e *echo.Echo, handler EchoHandler, method, path string, pathParams map[string]string, queryParams map[string]string, body interface{}) (*httptest.ResponseRecorder, error) {
	var req *http.Request

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
		b, _ := json.Marshal(body)
		req = httptest.NewRequest(method, path, bytes.NewReader(b))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if pathParams != nil {
		names := []string{}
		values := []string{}
		for k, v := range pathParams {
			names = append(names, k)
			values = append(values, v)
		}
		ctx.SetParamNames(names...)
		ctx.SetParamValues(values...)
	}

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	err := handler(ctx)
	return rec, err
}

func ParseResponse(t *testing.T, rec *httptest.ResponseRecorder, target interface{}) {
	if err := json.Unmarshal(rec.Body.Bytes(), target); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
}

func ParseEntityData(t *testing.T, wrapper interface{}, entity interface{}) {
	raw, err := json.Marshal(wrapper)
	if err != nil {
		t.Fatalf("marshal wrapper: %v", err)
	}
	if err := json.Unmarshal(raw, entity); err != nil {
		t.Fatalf("unmarshal entity: %v", err)
	}
}
