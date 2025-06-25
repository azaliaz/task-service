package rest

import "time"

func mapDuration(startedAt, completedAt time.Time) string {
	duration := ""
	if !startedAt.IsZero() {
		if !completedAt.IsZero() {
			duration = completedAt.Sub(startedAt).String()
		} else {
			duration = time.Since(startedAt).String()
		}
	}

	return duration
}
