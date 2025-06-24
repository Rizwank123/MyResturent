package integration

import (
	"net/http"
	"testing"

	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGetMenuCard(t *testing.T) {
	t.Run("should create resturent", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)

		reqBody := domain.CreateResturentInput{
			Name:    "Test Bistro",
			Address: domain.Address{City: "Go Town"},
			License: "license",
		}

		rec, err := helper.SendRequest(e, api.ResturentController.CreateResturent, http.MethodPost, "/resturent", nil, nil, reqBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
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

		// create and send request to create menu card
		reqBody1 := domain.CreateMenuCardInput{
			Name:        "Chicken Masala",
			ResturentID: created.ID,
			Price:       800,
			Size:        "Full Plate",
			Category:    "VEG",
			FoodType:    "MAINCOURSE",
			MealType:    "BREAKFAST",
			Image:       "https://www.pexels.com/photo/pot-with-stew-9609849/",
			IsAvailable: true,
			Description: "best dish of the day",
		}

		rec, err = helper.SendRequest(e, api.MenuCardController.CreateMenuCard, http.MethodPost, "/menu-card", nil, nil, reqBody1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		var rawResp1 map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp1)

		var createdMenuCard domain.MenuCard
		helper.ParseEntityData(t, rawResp1["data"], &createdMenuCard)

		if createdMenuCard.ID.IsNil() {
			t.Fatalf("expected non-zero ID, got %v", createdMenuCard.ID)
		}
		if createdMenuCard.Name != "Chicken Masala" {
			t.Fatalf("expected name 'Chicken Masala', got %v", createdMenuCard.Name)
		}

		// get menu card
		pathParams := map[string]string{"id": createdMenuCard.ID.String()}
		rec, err = helper.SendRequest(e, api.MenuCardController.GetMenuCard, http.MethodGet, "/menu-card/"+createdMenuCard.ID.String(), pathParams, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp2 map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp2)

		var menuCard domain.MenuCard
		helper.ParseEntityData(t, rawResp2["data"], &menuCard)

		if menuCard.ID.IsNil() {
			t.Fatalf("expected non-zero ID, got %v", menuCard.ID)
		}
		if menuCard.Name != "Chicken Masala" {
			t.Fatalf("expected name 'Chicken Masala', got %v", menuCard.Name)
		}
	})

}
func TestFilterAndSortMenuCards(t *testing.T) {
	// TODO implement test
	t.Run("Filter menu cards", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.MenuCardController.Filter, http.MethodPost, "/menu-card/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var menuCards []domain.MenuCard
		helper.ParseEntityData(t, rawResp["data"], &menuCards)

		if len(menuCards) == 0 {
			t.Fatalf("expected at least one menu card, got %v", menuCards)
		}

	})
}
func TestUpdateMenuCard(t *testing.T) {
	// TODO implement test
	t.Run("Update menu card", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)

		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.MenuCardController.Filter, http.MethodPost, "/menu-card/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var menuCards []domain.MenuCard
		helper.ParseEntityData(t, rawResp["data"], &menuCards)

		if len(menuCards) == 0 {
			t.Fatalf("expected at least one menu card, got %v", menuCards)
		}
		// Create and send request to update menu card
		reqBody := domain.UpdateMenuCardInput{
			Name:        "Chicken Masala",
			ResturentID: menuCards[0].ResturentID,
			Price:       800,
			Size:        "Full Plate",
			Category:    "VEG",
			FoodType:    "MAINCOURSE",
			MealType:    "BREAKFAST",
			Image:       "https://www.pexels.com/photo/pot-with-stew-9609849/",
			IsAvailable: true,
			Description: "best dish of the day",
		}
		pathParam := map[string]string{"id": menuCards[0].ID.String()}
		rec, err = helper.SendRequest(e, api.MenuCardController.UpdateMenuCard, http.MethodPatch, "/menu-card/"+menuCards[0].ID.String(), pathParam, nil, reqBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp1 map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp1)
		var menuCard domain.MenuCard
		helper.ParseEntityData(t, rawResp1["data"], &menuCard)

		if menuCard.ID.IsNil() {
			t.Fatalf("expected non-zero ID, got %v", menuCard.ID)
		}
		if menuCard.Name != "Chicken Masala" {
			t.Fatalf("expected name 'Chicken Masala', got %v", menuCard.Name)
		}

	})
}

func TestDeleteMenuCard(t *testing.T) {
	// TODO implement test
	t.Run("Delete menu card", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.MenuCardController.Filter, http.MethodPost, "/menu-card/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var menuCards []domain.MenuCard
		helper.ParseEntityData(t, rawResp["data"], &menuCards)

		if len(menuCards) == 0 {
			t.Fatalf("expected at least one menu card, got %v", menuCards)
		}
		// Create and send request to delete menu card
		pathParam := map[string]string{"id": menuCards[0].ID.String()}
		rec, err = helper.SendRequest(e, api.MenuCardController.DeleteMenuCard, http.MethodDelete, "/menu-card/"+menuCards[0].ID.String(), pathParam, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)

	})
}
