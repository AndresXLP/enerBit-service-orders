package utils

import (
	"fmt"
	"strings"
	"time"
)

func BuildQueryByStatus(query *string, status int32, operator string) {
	if len(*query) == 0 && status >= 0 {
		*query = fmt.Sprintf("status = '%s'", defineStatus(status))
	}
}

func defineStatus(status int32) string {
	var statusToSearch string
	switch status {
	case 0:
		statusToSearch = "new"
		break
	case 1:
		statusToSearch = "done"
		break
	case 2:
		statusToSearch = "cancelled"
		break
	}

	return statusToSearch
}

func BuildQueryByDate(query *string, since, until time.Time, operator string) {
	if len(*query) == 0 && !since.IsZero() && !until.IsZero() {
		*query = fmt.Sprintf("DATE(planned_date_begin) >= DATE('%s') AND DATE(planned_date_end) <= DATE('%s')", since.Format("2006-01-02"), until.Format("2006-01-02"))
	} else if len(*query) > 0 && !since.IsZero() && !until.IsZero() && len(operator) > 0 {
		*query = fmt.Sprintf(" %s %s %s", *query, strings.ToTitle(operator), fmt.Sprintf("DATE(planned_date_begin) >= DATE('%s') AND DATE(planned_date_end) <= DATE('%s')", since.Format("2006-01-02"), until.Format("2006-01-02")))
	}
}
