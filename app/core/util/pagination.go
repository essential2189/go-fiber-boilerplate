package util

import (
	"math"
)

func GetOffset(page int, limit int) int {
	return (page - 1) * limit
}

func CalculatePagination(limit int64, offset int64, total_count int64) (currentPage int64, totalPages int64) {
	currentPage = offset/limit + 1
	totalPages = int64(math.Ceil(float64(total_count) / float64(limit)))
	return
}
