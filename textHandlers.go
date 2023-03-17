package main

import "strings"

// Make Line Breaks on Submissions
// This function will make line breaks in a string after a certain number of characters.
// separateLines will take a string and return a string with line breaks after 30 characters.
func separateLines(message string) string {
	var lines []string
	line := ""
	for _, word := range strings.Fields(message) {
		if len(line)+len(word) > 60 {
			lines = append(lines, line)
			line = ""
		}
		line += word + " "
	}
	lines = append(lines, line)
	return strings.Join(lines, "\n")
}
