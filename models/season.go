package models

import (
	"github.com/gocolly/colly"
	"strings"
	"sync"
	"time"
)

type Season struct {
	Id       string     `json:"id"`
	SeriesId string     `json:"seriesId"`
	Episodes []*Episode `json:"episodes"`
}

// GetEpisodes fetches all episodes for one season of the series
func (s *Season) GetEpisodes(wg *sync.WaitGroup, errChn chan<- error) {

	defer wg.Done()

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := c.Limit(
		&colly.LimitRule{
			DomainGlob:  "*",
			RandomDelay: 2 * time.Second,
		},
	)

	if err != nil {
		errChn <- err
		return
	}

	// episode names
	var titles []string
	var ids []string
	// rating of the each episode
	//var ratings []string
	// episode summary
	var descriptions []string

	c.OnHTML("a[title]", func(e *colly.HTMLElement) {
		if e.Attr("itemprop") == "name" {
			titles = append(titles, e.Attr("title"))                 // episode title
			ids = append(ids, strings.Split(e.Attr("href"), "/")[2]) // episode id
		}
	})

	//c.OnHTML("span.ipl-rating-star__rating", func(e *colly.HTMLElement) {
	//	fmt.Println(e.Text)
	//})

	c.OnHTML("div.item_description", func(e *colly.HTMLElement) {
		descriptions = append(descriptions, strings.TrimSpace(e.Text)) // episode description
	})

	err = c.Visit("https://www.imdb.com/title/" + s.SeriesId + "/episodes?season=" + s.Id)
	if err != nil {
		return
	}

	c.Wait()

	var eps []*Episode

	for i := 0; i < len(titles); i++ {
		e := &Episode{
			Id:          ids[i],
			SeriesId:    s.SeriesId,
			Season:      s.Id,
			Name:        titles[i],
			Rating:      0,
			Description: descriptions[i],
			PosterId:    "",
			PosterLink:  "",
		}

		eps = append(eps, e)
	}

	s.Episodes = eps

	return
}
