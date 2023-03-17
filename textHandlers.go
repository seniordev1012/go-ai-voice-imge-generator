package main

import "strings"

// Make Line Breaks on Submissions
// This function will make line breaks in a string after a certain number of characters.
// separateLines will take a string and return a string with line breaks after 60 characters.
func separateLines(message string) string {
	var lines []string
	line := ""
	paragraph := ""
	count := 0
	hasFullStop := false

	words := strings.Fields(message)
	for i, word := range words {
		if len(line)+len(word) > 60 {
			// Check if we need to add a new paragraph
			if count%5 == 0 && hasFullStop && i < len(words)-1 {
				// Add the paragraph and reset variables
				lines = append(lines, paragraph)
				paragraph = ""
				count = 0
				hasFullStop = false
			}
			// Add the line and reset variables
			lines = append(lines, line)
			line = ""
			count++
		}
		line += word + " "
		if strings.HasSuffix(word, ".") {
			hasFullStop = true
		}
	}

	// Add any remaining text to the last line
	lines = append(lines, line)

	// Add the final paragraph if it hasn't been added yet
	if len(paragraph) > 0 {
		lines = append(lines, paragraph)
	}

	// Join the lines with newlines to form the final string
	return strings.Join(lines, "\n")
}
