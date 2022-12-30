package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func R2UScparer(corectedWord string) (string, error) {
	c := colly.NewCollector()
	resultSls := make([]string, 0)

	c.OnHTML("#table_1 > tbody tr", func(e *colly.HTMLElement) {
		scrapWord := e.DOM.Text()
		scrapWord, _, found := strings.Cut(scrapWord, " ")
		scrapWord = strings.ToLower(scrapWord)
		if found && !strings.Contains(scrapWord, "-") && (len(resultSls) == 0 || len(resultSls) > 0 && resultSls[0] != scrapWord) {
			resultSls = append(resultSls, scrapWord)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	fmt.Printf("Visiting(sync): %s\n", corectedWord)
	c.Visit(fmt.Sprintf("https://r2u.org.ua/vesum/?w=%s&all_forms=on", corectedWord))
	if len(resultSls) == 1 {
		return resultSls[0], nil
	}
	return "", fmt.Errorf("it's some problem with scraping this word -> %s", corectedWord)
}
