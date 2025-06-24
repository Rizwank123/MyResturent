package integration

import (
	"net/http"
	"testing"

	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestUserEndpoints(t *testing.T) {
	t.Run("Register User", func(t *testing.T) {
		api, e, td := helper.SetupSuite(t)
		defer td(t)

		payload := domain.CreateUserInput{
			Name:     "Test User",
			Email:    "testuser@example.com",
			Password: "secure123",
			Role:     "OWNER",
			Mobile:   "1234567890",
		}

		rec, err := helper.SendRequest(e, api.UserController.Register, http.MethodPost, "/user/register", nil, nil, payload)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var created domain.User
		helper.ParseEntityData(t, rawResp["data"], &created)

		assert.Equal(t, "Test User", created.Name)
		assert.Equal(t, "testuser@example.com", created.Email)
	})

	t.Run("Login User", func(t *testing.T) {
		api, e, td := helper.SetupSuite(t)
		defer td(t)

		payload := domain.LoginUserInput{
			Email:    "testuser@example.com",
			Password: "secure123",
		}

		rec, err := helper.SendRequest(e, api.UserController.Login, http.MethodPost, "/user/login", nil, nil, payload)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var login domain.LoginResponse
		helper.ParseEntityData(t, rawResp["data"], &login)
		if login.Token == "" {
			t.Fatalf("Expected JWT token, got %v", login.Token)
		}
	})

	t.Run("Filter Users", func(t *testing.T) {
		api, e, td := helper.SetupSuite(t)
		defer td(t)

		filter := domain.FilterInput{
			Fields: []domain.FilterFieldPredicate{
				{
					Field:    "email",
					Operator: domain.FilterOpEq,
					Value:    "testuser@example.com",
				},
			},
		}

		rec, err := helper.SendRequest(e, api.UserController.Filter, http.MethodPost, "/user/filter", nil, nil, filter)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp domain.PaginationResponse
		helper.ParseResponse(t, rec, &resp)

		var users []domain.User
		helper.ParseEntityData(t, resp.Data, &users)

		assert.NotEmpty(t, users)
		assert.Equal(t, "testuser@example.com", users[0].Email)
	})

	t.Run("Update User", func(t *testing.T) {
		api, e, td := helper.SetupSuite(t)
		defer td(t)

		filter := domain.FilterInput{
			Fields: []domain.FilterFieldPredicate{
				{
					Field:    "email",
					Operator: domain.FilterOpEq,
					Value:    "testuser@example.com",
				},
			},
		}

		rec1, err := helper.SendRequest(e, api.UserController.Filter, http.MethodPost, "/user/filter", nil, nil, filter)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec1.Code)

		var resp domain.PaginationResponse
		helper.ParseResponse(t, rec1, &resp)

		var users []domain.User
		helper.ParseEntityData(t, resp.Data, &users)

		payload := domain.UpdateUserInput{
			Name: "Updated Test User",
		}

		pathParams := map[string]string{
			"id": users[0].ID.String(),
		}

		rec, err := helper.SendRequest(e, api.UserController.UpdateUser, http.MethodPut, "/user/:id", pathParams, nil, payload)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var updated domain.User
		helper.ParseEntityData(t, rawResp["data"], &updated)

		assert.Equal(t, "Updated Test User", updated.Name)
		assert.Equal(t, users[0].ID, updated.ID)
	})
}
