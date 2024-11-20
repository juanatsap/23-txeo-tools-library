package tools

import (
	"strings"
	"unicode"
)

func StartsWithNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	r := []rune(s)
	return unicode.IsDigit(r[0])
}
func IsValidDate(s string) bool {
	// Implementar validación de fecha si es necesario
	months := []string{"Ene", "Feb", "Mar", "Abr", "May", "Jun", "Jul", "Ago", "Sep", "Oct", "Nov", "Dic"}
	for _, month := range months {
		if strings.Contains(s, month) {
			return true
		}
	}
	return false
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
