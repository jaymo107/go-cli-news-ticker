package main

import (
	"fmt"
	"strconv"

	"github.com/jaymo107/go-cli-news-ticker/news"
)

const (
	topStories = iota + 1
	QUIT       = "q"
)

func main() {
	var choice string

	printChoices()

	for {
		fmt.Scanln(&choice)

		if choice == QUIT {
			break
		}

		if val, _ := strconv.Atoi(choice); val == topStories {
			if stories, err := news.GetNews(); err == nil {
				printNews(stories)
			} else {
				fmt.Println("failed to get news", err)
			}
		}
	}
}

func printChoices() {
	fmt.Println("Please enter your choice below then press enter:")
	fmt.Println("(1) Get top stories from all sources")
	fmt.Println("(q) Quit")
}

func printNews(stories []news.Story) {
	for _, story := range stories {
		fmt.Println(story.String())
	}
}
