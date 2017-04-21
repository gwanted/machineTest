package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tealeg/xlsx"

	"app"
	"common"
	"conf"
	"model"
	"mydb"
)

func main() {
	mydb.InitDB(conf.App.DBAddress)

	http.HandleFunc("/analysis", app.Analysis)
	http.HandleFunc("/register", app.Register)

	http.ListenAndServe(":8888", nil)
}

func ImportDataLinks() {
	file, err := xlsx.OpenFile("data/links.xlsx")
	if err != nil {
		panic(err.Error())
	}

	linksAll := []model.Links{}
	err = common.UnmarshalSheet(file.Sheets[0], &linksAll)
	if err != nil {
		panic(err.Error())
		return
	}
	for _, links := range linksAll {
		err = links.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func ImportDataGenomeTags() {
	file, err := os.Open("data/genome-tags.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tags := model.GenomeTags{}
		for i, x := range record {
			if i == 0 {
				tags.TagID = x
			} else {
				tags.Tag = x
			}
		}
		err = tags.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func ImportDataGenomeScores() {
	file, err := os.Open("data/genome-scores.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		scores := model.GenomeScores{}
		for i, x := range record {
			if i == 0 {
				scores.MovieID = x
			} else if i == 1 {
				scores.TagID = x
			} else {
				scores.Relevance = x
			}
		}
		err = scores.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func ImportDataMovies() {
	file, err := os.Open("data/movies.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		movies := model.Movies{}
		for i, x := range record {
			if i == 0 {
				movies.MovieID = x
			} else if i == 1 {
				movies.Title = x
			} else {
				movies.Genres = x
			}
		}
		err = movies.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func ImportDataRatings() {
	file, err := os.Open("data/ratings.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		ratings := model.Ratings{}
		for i, x := range record {
			if i == 0 {
				ratings.UserID = x
			} else if i == 1 {
				ratings.MovieID = x
			} else if i == 2 {
				ratings.Rating = x
			}
		}
		err = ratings.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func ImportDataTags() {
	file, err := os.Open("data/tags.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		tags := model.Tags{}
		for i, x := range record {
			if i == 0 {
				tags.UserID = x
			} else if i == 1 {
				tags.MovieID = x
			} else if i == 2 {
				tags.Tag = x
			}
		}
		err = tags.Insert()
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}
