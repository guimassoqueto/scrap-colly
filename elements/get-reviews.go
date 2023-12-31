package elements

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func extractReviews(eText string) int {
	regex, _ := regexp.Compile(`[\d\.]+`)
	match := regex.FindString(eText)
	removedDots := strings.ReplaceAll(match, ".", "")
	reviews, err := strconv.Atoi(removedDots)
	if err != nil {
		return 0
	}
	return reviews
}

func GetReviews(c *colly.Collector, reviews *int) {
	c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
		*reviews = extractReviews(e.Text)
	})
}