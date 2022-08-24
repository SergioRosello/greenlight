package validator

import (
	"github.com/SergioRosello/greenlight/internal/data"
)

func (v *Validator) ValidateFilters(filter data.Filters) {
	v.Check(filter.Page >= 1, "page", "must be greater than zero")
	v.Check(filter.Page <= 10_000_000, "page", "must be lesser than ten million")
	v.Check(filter.PageSize >= 1, "page", "must be greater than zero")
	v.Check(filter.PageSize <= 100, "page", "must be lesser than one hundred")
	v.Check(PermittedValue(filter.Sort, filter.SortSafelist...), "sort", "invalid sort value")
}
