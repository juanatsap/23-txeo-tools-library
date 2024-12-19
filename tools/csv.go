package tools

import (
	"bytes"
	"encoding/csv"
)

func ToCSV(header []string, rows [][]string) string {
	var b bytes.Buffer
	w := csv.NewWriter(&b)
	w.Write(header)
	for _, row := range rows {
		w.Write(row)
	}
	w.Flush()
	return b.String()
}
