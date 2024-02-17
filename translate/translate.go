// An exercise in doing string comparison and replacement

package main

import (
	"fmt"
	"strings"
)

var sentence = "Last week the Royal Festival Hall saw the first performance of a new symphony by one of the world leading modern composers, Arthur 'Two Sheds' Jackson"

var englishToDutch = map[string]string{
	"Last":        "Laatst",
	"week":        "week",
	"the":         "de",
	"Royal":       "Koninklijk",
	"Festival":    "Feest",
	"Hall":        "Hal",
	"saw":         "zaag",
	"first":       "eerst",
	"performance": "optreden",
	"of":          "van",
	"a":           "een",
	"new":         "nieuw",
	"symphony":    "symphonie",
	"by":          "bij",
	"one":         "een",
	"world":       "wereld",
	"leading":     "leidend",
	"modern":      "modern",
	"composer":    "componist",
	"composers":   "componisten",
	"Two":         "Twee",
	"Sheds":       "Schuren",
}

func main() {

	fmt.Println("The original English sentence is:")
	fmt.Println(sentence)
	fmt.Println()

	// Remove the punctuation, so it's not seen as part of the words
	// TODO: I'd like to be able to put the punctuation back
	removeApost := strings.ReplaceAll(sentence, "'", "")
	removeComma := strings.ReplaceAll(removeApost, ",", "")

	// Convert the string to a slice for processing
	original := strings.Split(removeComma, " ")

	// This will get all the translated text
	translated := ""

	// Go through the sentence and find translation if it exists
	for _, word := range original {
		holder := word
		// In order to match the map, word must be in quotes
		fmt.Scan("%q", &word)
		// Look for the translation
		replacement := englishToDutch[word]
		// If there is a translation, add it to "translated" with a space following
		if replacement != "" {
			translated += replacement
			translated += " "
			// If there's no translation, add the default word to "translated" with following space
		} else {
			translated += holder
			translated += " "
		}
	}
	// The last word will be followed by a space, so we strip that and print the translation
	strings.TrimSpace(translated)
	fmt.Println("Translated to Dutch is:")
	fmt.Println(translated)
}
