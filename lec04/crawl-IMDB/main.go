package main

import (
    // "log"
    db "crawl-IMDB/storage"
	cr "crawl-IMDB/crawler"
    _ "github.com/go-sql-driver/mysql"
)

func main(){

    db.InitDB()

	cr.Crawler()


}
