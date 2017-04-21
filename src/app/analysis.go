package app

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"

	"common"
	"model"
)

type userDis struct {
	UserID string
	Dis    float64
}

type userDisList []userDis

func (a userDisList) Len() int {
	return len(a)
}
func (a userDisList) Less(i, j int) bool {
	return a[i].Dis < a[j].Dis
}
func (a userDisList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Analysis get the movie list which will be push to current user
func Analysis(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("analysis start %v\n", time.Now())
	userID := r.FormValue("userID")
	if userID == "" {
		common.ReturnEFormat(w, 1, "userID is null")
		return
	}
	ratings, err := model.FindRatings(bson.M{"userID": userID})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}

	sumRate := 0.0

	movieIDs := []string{}
	for _, rating := range ratings {
		tmpRate, _ := strconv.ParseFloat(rating.Rating, 64)
		sumRate += tmpRate
		movieIDs = append(movieIDs, rating.MovieID)
	}
	averageRate := sumRate / float64(len(ratings))
	fmt.Printf("averageRate %v\n", averageRate)
	tmp, err := model.FindRatings(bson.M{"movieID": bson.M{"$in": movieIDs}})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}

	sxs := map[string][]*model.Ratings{}
	neighbors := []string{}
	for _, x := range tmp {
		sxs[x.UserID] = append(sxs[x.UserID], x)
		neighbors = append(neighbors, x.UserID)
	}

	userDisLists := userDisList{}
	for _, neighbor := range neighbors {
		dis := getCosDistance(ratings, sxs[neighbor])
		userDisTmp := userDis{}
		userDisTmp.UserID = neighbor
		userDisTmp.Dis = dis
		userDisLists = append(userDisLists, userDisTmp)
	}

	sort.Sort(sort.Reverse(userDisList(userDisLists)))
	movieList := []*model.Ratings{}
	movieReturnIDs := []string{}

	neighborIDs := []string{}
	for i := 0; i < 5; i++ {
		neighborIDs = append(neighborIDs, neighbors[i])
	}

	ratingNeighbors, err := model.FindRatings(bson.M{"userID": bson.M{"$in": neighborIDs}})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}

	//for i := 0; i < len(userDisLists); i++ {
	for _, y := range ratingNeighbors {

		sign := false
		for _, k := range ratings {
			if y.MovieID == k.MovieID {
				sign = true
				break
			}
		}
		if !sign {
			movieList = append(movieList, y)
		}
	}

	finalMovieList := []*model.Ratings{}

	for _, m := range movieList {
		tmpRate, _ := strconv.ParseFloat(m.Rating, 64)
		if tmpRate > averageRate {
			movieReturnIDs = append(movieReturnIDs, m.MovieID)
			finalMovieList = append(finalMovieList, m)
		}
	}

	movies, err := model.FindMovies(bson.M{"movieID": bson.M{"$in": movieReturnIDs}})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}

	fmt.Printf("analysis end %v\n", time.Now())

	common.ReturnFormat(w, 0, movies, "SUCCESS")
}

func getCosDistance(user1, user2 []*model.Ratings) float64 {
	sumx := 0.0
	sumy := 0.0
	sumxy := 0.0
	for _, u1 := range user1 {
		for _, u2 := range user2 {
			if u1.MovieID == u2.MovieID {
				r1, _ := strconv.ParseFloat(u1.Rating, 64)
				r2, _ := strconv.ParseFloat(u2.Rating, 64)

				sumx += r1 * r1
				sumy += r2 * r2
				sumxy += r1 * r2
			}
		}
	}
	if sumxy == 0.0 {
		return 0.0
	}
	demo := math.Sqrt(sumx * sumy)
	return sumxy / demo
}
