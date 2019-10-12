package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type Tweet struct {
	TwitterID string `gorm:"unique_index"`
	Tweet     string `gorm:"type:varchar(512)"`
}

func main() {
	db, err := gorm.Open("sqlite3", "tweets.db")
	if err != nil {
		fmt.Errorf("Tweets db open failed.\n")
		panic(err)
	}
	defer db.Close()

	csvFile, err := os.Create("tweets.csv")
	if err != nil {
		fmt.Errorf("Tweets csv open failed.\n")
		panic(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	var tweets []Tweet

	db.Find(&tweets)

	for _, tweet := range tweets {
		row := []string{tweet.TwitterID, tweet.Tweet}
		if err := csvWriter.Write(row); err != nil {
			panic(err)
		}
	}

	csvWriter.Flush()
}
