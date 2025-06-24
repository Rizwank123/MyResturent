package integration

import (
	"net/http"
	"testing"

	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/tests/helper"
	"github.com/stretchr/testify/assert"
)

func TestCreateRating(t *testing.T) {
	t.Run("Create rating", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// create and send request to get the resturent
		filter := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.ResturentController.Filter, http.MethodPost, "/resturent/filter", nil, nil, filter)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var resturents []domain.Resturent
		helper.ParseEntityData(t, rawResp["data"], &resturents)

		if len(resturents) == 0 {
			t.Fatalf("expected at least one resturent, got %v", resturents)
		}

		// create and send request to create rating
		reqBody := domain.CreateRatingInput{
			ResturentID: resturents[0].ID,
			Rating:      5,
			Name:        "Alex t",
			Review:      "Good",
			Suggestion:  "Make sure reduce serving time",
		}

		rec, err = helper.SendRequest(e, api.RatingController.CreateRating, http.MethodPost, "/rating", nil, nil, reqBody)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		var rawResp1 map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp1)

		var created domain.Rating
		helper.ParseEntityData(t, rawResp1["data"], &created)

		if created.ID.IsNil() {
			t.Fatalf("expected non-zero ID, got %v", created.ID)
		}

		if created.Name != "Alex t" {
			t.Fatalf("expected name 'Alex t', got %v", created.Name)
		}

		if created.Rating != 5 {
			t.Fatalf("expected rating '5', got %v", created.Rating)
		}

	})
}

func TestGetRating(t *testing.T) {
	// TODO implement test
	t.Run("Filter rating", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.RatingController.Filter, http.MethodPost, "/rating/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var ratings []domain.MenuCard
		helper.ParseEntityData(t, rawResp["data"], &ratings)

		if len(ratings) == 0 {
			t.Fatalf("expected at least one menu card, got %v", ratings)
		}
	})

	t.Run("Get rating by id", func(t *testing.T) {
		// TODO implement test
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.RatingController.Filter, http.MethodPost, "/rating/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var ratings []domain.Rating
		helper.ParseEntityData(t, rawResp["data"], &ratings)

		if len(ratings) == 0 {
			t.Fatalf("expected at least one menu card, got %v", ratings)
		}
		pathParam := map[string]string{"id": ratings[0].ID.String()}
		rec, err = helper.SendRequest(e, api.RatingController.FindByID, http.MethodGet, "/rating/{id}", pathParam, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

	})
}

func TestUpdateRating(t *testing.T) {
	t.Run("Update rating", func(t *testing.T) {
		api, e, teardown := helper.SetupSuite(t)
		defer teardown(t)
		// Create and send request to create menu card
		filterInput := domain.FilterInput{}

		rec, err := helper.SendRequest(e, api.RatingController.Filter, http.MethodPost, "/rating/filter", nil, nil, filterInput)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var rawResp map[string]interface{}
		helper.ParseResponse(t, rec, &rawResp)

		var ratings []domain.Rating
		helper.ParseEntityData(t, rawResp["data"], &ratings)

		if len(ratings) == 0 {
			t.Fatalf("expected at least one menu card, got %v", ratings)
		}
		pathParam := map[string]string{"id": ratings[0].ID.String()}

		in := domain.UpdateRatingInput{
			Name:       "Alex t",
			Rating:     4,
			Review:     "Good",
			Suggestion: "Make sure reduce serving time",
		}
		rec, err = helper.SendRequest(e, api.RatingController.UpdateRating, http.MethodPut, "/rating/"+ratings[0].ID.String(), pathParam, nil, in)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

	})
}
