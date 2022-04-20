package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Stories struct {
	Title  []Story  `json:"top_stories"`
	PageTitle string `json:"page_title"`
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

		dataCh := processIDS(ids)

		for story := range dataCh{
			stories = append(stories, story)
		}

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

	tmpl := template.Must(template.ParseFiles("template.html"))
	
	// tmpl, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }



	mux.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ids := getIds()

		stories := []Story{}

		dataCh := processIDS(ids)

		for story := range dataCh{
			stories = append(stories, story)
		}

		data := Stories{Title: stories, PageTitle:"Top News"}

		tmpl.Execute(w, data)
	})
	mux.HandleFunc("/api/top", HandleTopStories())

	log.Fatal(http.ListenAndServe(":8000", mux))
}




func HandleFuncIds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(result)



	}
}