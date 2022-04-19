package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	// "strings"
)

// Book struct (Model)
type Stories struct {
	Title  []Story  `json:"top_stories"`
}




type Story struct {
	Title string `json:"title"`
	Score int `json:"score"`

}


func worker(wg *sync.WaitGroup, ch chan Story, id int) {
	defer wg.Done()
	strID := strconv.Itoa(id)
	story := getStory(strID)
	ch <- story
}

func monitorWorker(wg *sync.WaitGroup, ch chan Story){
	wg.Wait()
	close(ch)
}

func processIDS(ids []int) chan Story {
	wg := &sync.WaitGroup{}
	dataCh := make(chan Story)

	for i := 0; i < len(ids); i++ {
		id := ids[i]
		wg.Add(1)
		go worker(wg, dataCh, id)
	}
	go monitorWorker(wg, dataCh)

	return dataCh

}




func HandleTopStories() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")

		ids := getIds()

		stories := []Story{}
		// fmt.Println(ids)

		dataCh := processIDS(ids)

		// SLOW WAY WITHOUT GENERATOR

		// for _, id := range ids {
		// 	strID := strconv.Itoa(id)
		// 	story := getStory(strID)
		// 	stories = append(stories, story)
		// }

		//FAST WAY
		for story := range dataCh{
			stories = append(stories, story)
		}



		

		// stories = append(stories, Story{Title:"HTML", Score: 200})


		object := Stories{Title: stories}
		json.NewEncoder(w).Encode(&object)
	}
}


func getIds() []int {
	urlID := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

	req, err := http.NewRequest("GET", urlID, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer res.Body.Close()
		var data []int
	
		json.NewDecoder(res.Body).Decode(&data)
	
		var result []int

		for i := 0; i < 10; i++ {
			result = append(result, data[i])
		}
		return result
}

func getStory(id string) Story {

	urlStory := "https://hacker-news.firebaseio.com/v0/item/" + id + ".json?print=pretty"

	req, err := http.NewRequest("GET", urlStory, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	var data Story

	json.NewDecoder(res.Body).Decode(&data)
	return data
}


func main() {
	mux := http.NewServeMux()



	mux.HandleFunc("/top", HandleTopStories())
	mux.HandleFunc("/ids", HandleFuncIds())

	log.Fatal(http.ListenAndServe(":8000", mux))
}




func HandleFuncIds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

		// response, err := http.Get(urlID)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer response.Body.Close()

		// responseData, err := ioutil.ReadAll(response.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// ids := strings.Split(string(responseData), ", ")
		// fmt.Println(ids[0])
		req, err := http.NewRequest("GET", urlID, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer res.Body.Close()
		var data []int
	
		json.NewDecoder(res.Body).Decode(&data)
	
		var result []int

		for i := 0; i < 10; i++ {
			result = append(result, data[i])
		}
		fmt.Println(result)


		
	}
}