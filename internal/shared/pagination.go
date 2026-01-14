package shared

import (
	"database/sql"
	"math"
	"net/http"
	"strconv"
)

func CountRows(db *sql.DB, query string, args ...any) (int, error) {
	var total int
	err := db.QueryRow(query, args...).Scan(&total)
	return total, err
}

func GetPagination(r *http.Request) (page, limit, offset int) {
	page = 1
	limit = 10

	q := r.URL.Query()

	if p := q.Get("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	if l := q.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}

	offset = (page - 1) * limit
	return
}

func CalculateTotalPages(totalItems, limit int) int {
	if limit == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalItems) / float64(limit)))
}
