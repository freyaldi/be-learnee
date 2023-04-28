package util

import (
	"strings"

	"git.garena.com/sea-labs-id/batch-06/ferza-reyaldi/stage01-project-backend/dto"
)

func FormatCourseQuery(query *dto.CourseRequestQuery) *dto.CourseRequestQuery {
	if query.Limit == 0 {
		query.Limit = 8
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" || query.SortBy == "" {
		query.SortBy = "created_at"
	}

	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}