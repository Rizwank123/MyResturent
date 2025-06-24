package service

import (
	"github.com/gofrs/uuid/v5"
	"github.com/rizwank123/myResturent/internal/domain"
)

func SetupResturentIDFilter(in *domain.FilterInput) {
	if in.ResturentID != uuid.Nil {
		in.Fields = append(in.Fields, domain.FilterFieldPredicate{
			Field:    "resturent_id",
			Operator: domain.FilterOpEq,
			Value:    in.ResturentID,
		})
	}
}
