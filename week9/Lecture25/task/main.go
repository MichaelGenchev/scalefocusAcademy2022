package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"os"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"time"
	// sq "github.com/Masterminds/squirrel"
)

type Stories struct {
	Title  []Story  `json:"top_stories"`
	PageTitle string `json:"page_title"`
}

func addStoryToDataBase(story Story){
	db, err := sqlx.Open("sqlite", dbName)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now().Unix()
	db.Exec("INSERT INTO stories (id, title, time , score) VALUES ($1, $2, $3, $4)", story.ID, story.Title, currentTime, story.Score)
	
}


type Story struct {
	Title string `json:"title"`
	Score int `json:"score"`
	Time int `json:"time"`
	ID int `json:"id"`

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

const TIME_TO_WAIT = 3600

func HandleTopStories() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){

		w.Header().Set("Content-Type", "application/json")
		db, err := sqlx.Open("sqlite", dbName)
		if err != nil {
			log.Fatal(err)
		}
		stories := []Story{}

		story := Story{}
		rows, err := db.Queryx("SELECT * FROM stories")
    	for rows.Next() {
			err := rows.StructScan(&story)
        	if err != nil {
				log.Fatalln(err)
			}
			currentTime := time.Now().Unix()
			if int(currentTime) - story.Time < TIME_TO_WAIT {
				stories = append(stories, Story{ID: story.ID, Title: story.Title, Score: story.Score, Time: story.Time})
			}
			
		}
		if len(stories) == 0 {
			// fmt.Printf("HAVE TO GET FROM HN")
			ids := getIds()
			dataCh := processIDS(ids)

			for story := range dataCh{
				stories = append(stories, story)
			}

			for _, story := range stories {
				addStoryToDataBase(story)
				
			}	
			object := Stories{Title: stories}
			// fmt.Println("FROM SERVER")
			json.NewEncoder(w).Encode(&object)
		}else {
			// fmt.Println("WILL GET FROM DATABASE")
			resultFromDatabase := Stories{Title: stories}
			json.NewEncoder(w).Encode(&resultFromDatabase)
		}
	}
}

const dbName = "stories.db"

func startDB(){
	os.Remove(dbName)

	db, err := sqlx.Open("sqlite", dbName)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	_, err = db.Exec(`Create table stories (id INT PRIMARY KEY, title VARCHAR(32), time INT, score INT)`)

	if err != nil {
		log.Print("Here")

		log.Fatal(err)
	}
}


func main() {

	startDB()

	mux := http.NewServeMux()

	tmpl := template.Must(template.ParseFiles("template.html"))
	
	// tmpl, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	mux.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		db, err := sqlx.Open("sqlite", dbName)
		if err != nil {
			log.Fatal(err)
		}
		stories := []Story{}

		story := Story{}
		rows, err := db.Queryx("SELECT * FROM stories")
    	for rows.Next() {
			err := rows.StructScan(&story)
        	if err != nil {
				log.Fatalln(err)
			}
			currentTime := time.Now().Unix()
			if int(currentTime) - story.Time < TIME_TO_WAIT {
				stories = append(stories, Story{ID: story.ID, Title: story.Title, Score: story.Score, Time: story.Time})
			}
			
		}
		if len(stories) == 0 {
			fmt.Printf("HAVE TO GET FROM HN")
			ids := getIds()


			dataCh := processIDS(ids)

			for story := range dataCh{
				stories = append(stories, story)
			}
			for _, story := range stories {
				addStoryToDataBase(story)
				
			}	

			data := Stories{Title: stories, PageTitle:"Top News"}

			tmpl.Execute(w, data)
			}else {
				fmt.Println("WILL GET FROM DATABASE")
				resultFromDatabase := Stories{Title: stories, PageTitle:"Top News"}
				tmpl.Execute(w, resultFromDatabase)
		}
	})

	mux.HandleFunc("/api/top", HandleTopStories())

	log.Fatal(http.ListenAndServe(":8001", mux))
}




func HandleFuncIds() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"


		req, _ := http.NewRequest("GET", urlID, nil)
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


func getIds() []int {
	urlID := "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"

	req, _ := http.NewRequest("GET", urlID, nil)
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

	req, _ := http.NewRequest("GET", urlStory, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	var data Story

	json.NewDecoder(res.Body).Decode(&data)
	return data
}