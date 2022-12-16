package utils

import (
	"fmt"
	"math"
	"strings"
)

func ForHumans(o int) string {
	r := [][]any{
		{int(math.Floor(float64(o) / 31536000)), "years"},
		{int(math.Floor(float64(o%31536000) / 86400)), "days"},
		{int(math.Floor(float64((o%31536000)%86400) / 3600)), "hours"},
		{int(math.Floor(float64(((o%31536000)%86400)%3600) / 60)), "minutes"},
		{(((o % 31536000) % 86400) % 3600) % 60, "seconds"},
	}

	var e []string

	for t := 0; t < len(r); t++ {
		if r[t][0] != 0 {
			e = append(e, fmt.Sprintf("%d %s", r[t][0], r[t][1]))
		}
	}

	return strings.TrimSpace(strings.Join(e, " "))
}
