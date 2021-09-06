package crawler

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
	db "crawl-IMDB/storage"
	model "crawl-IMDB/model"
)


func Crawler() {

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %s", r.URL)
		fmt.Println()
	})
	c.OnError(func(r *colly.Response, e error) {
		log.Println("Error: ", e)
	})


	c.OnHTML("tr", func(h *colly.HTMLElement) {

		movie := model.Movie{}
		movie.Name = h.ChildText(".titleColumn > a")
		movie.Year = h.ChildText(".titleColumn .secondaryInfo")
		movie.Rate = h.ChildText(".ratingColumn > strong")
		movie.ImgSrc = "imdb.com/"+h.ChildAttr(".posterColumn > a","href")

		newMovie,err := db.AddToDB(movie)

		if err!= nil{
			log.Println(err)
		}

		log.Println(newMovie)
		
	
	})

	
	c.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")

}
