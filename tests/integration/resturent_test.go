package integration

import (
	"net/http"
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/tests/helper"
)

func TestCreateAndGetResturent(t *testing.T) {
	t.Run("should create resturent", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)

		reqBody := domain.CreateResturentInput{
			Name:    "Test Bistro",
			Address: domain.Address{City: "Go Town"},
			License: "license",
		}

		rec, err := helper.SendRequest(e, api.ResturentController.CreateResturent, http.MethodPost, "/resturent", nil, nil, reqBody)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		if rec.Code != http.StatusCreated {
			t.Fatalf("expected status 200, got %v", rec.Code)
		}

		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var created domain.Resturent
		helper.ParseEntityData(t, rawResp["data"], &created)

		if created.ID.IsNil() {
			t.Fatalf("expected non-zero ID, got %v", created.ID)
		}
		if created.Name != "Test Bistro" {
			t.Fatalf("expected name 'Test Bistro', got %v", created.Name)
		}
		if created.Address.City != "Go Town" {
			t.Fatalf("expected location 'Go Town', got %v", created.Address.City)
		}
	})

	t.Run("should return not found for invalid id", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)

		// Simulate GET /resturent/:id with a non-existent ID
		id := uuid.Must(uuid.NewV4())
		pathParams := map[string]string{
			"id": id.String(), // assuming this doesn't exist
		}
		rec, err := helper.SendRequest(e, api.ResturentController.FindById, http.MethodGet, "/resturent/"+id.String(), pathParams, nil, nil)
		if err != nil && err != pgx.ErrNoRows {
			t.Fatalf("error sending request: %v", err)
		}
		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %v", rec.Code)
		}
	})

	t.Run("should list all resturents", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		input := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.ResturentController.Filter, http.MethodGet, "/resturent/filter", nil, nil, input)
		if err != nil {
			t.Fatalf("error sending request: %v", err)
		}
		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %v", rec.Code)
		}

		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var entityData []domain.Resturent
		helper.ParseEntityData(t, rawResp["data"], &entityData)

		if len(entityData) == 0 {
			t.Fatalf("expected at least 1 resturent in response")
		}
	})
}
