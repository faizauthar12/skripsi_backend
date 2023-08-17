package utils

import (
	"fmt"
	"math"
)

func HandleAdminPagination(pages int64, pageLink string, itemCount int64) (int64, string, string) {
	var maxPages int64

	if itemCount > 10 {
		maxPages = int64(math.Ceil(float64(itemCount) / 10))
	} else {
		maxPages = 1
	}

	var previousPageLink string

	nextPageLink := fmt.Sprintf("/admin/%s?pages=%d", pageLink, pages+1)
	if pages > 0 {
		previousPageLink = fmt.Sprintf("/admin/%s?pages=%d", pageLink, pages-1)
	}

	return maxPages, previousPageLink, nextPageLink
}
