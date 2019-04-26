package main

import (
	"flag"
	trello "github.com/adlio/trello"
	"log"
	"os"
	// "reflect"
)

func main() {
	// get new client
	appKey := os.Getenv("TRELLO_KEY")
	token := os.Getenv("TRELLO_TOKEN")
	if appKey == "" || token == "" {
		log.Fatalln("You need Trello key or token.")
	}
	client := trello.NewClient(appKey, token)

	// get member
	user := os.Getenv("TRELLO_USER")
	if user == "" {
		log.Fatalln("You need Trello user.")
	}
	member, err := client.GetMember(user, trello.Defaults())
	if err != nil {
		log.Fatalf("Get Member ERROR: %s", err.Error())
	}

	// get boards
	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		log.Fatalf("Client ERROR: %s", err.Error())
	}
	var (
		BOARD_NAME = flag.String("board", "Work", "trello board name")
		LIST_NAME  = flag.String("list", "In progress", "trello list name")
	)
	flag.Parse()
	var boardTargetNum int
	for i, board := range boards {
		if board.Name == *BOARD_NAME {
			boardTargetNum = i
		}
	}

	// get list
	lists, err := boards[boardTargetNum].GetLists(trello.Defaults())
	if err != nil {
		log.Fatalf("Get Lists ERROR: %s", err.Error())
	}
	// log.Println(lists)
	var listTargetNum int
	for i, list := range lists {
		if list.Name == *LIST_NAME {
			listTargetNum = i
		}
	}

	// get cards
	cards, err := lists[listTargetNum].GetCards(trello.Defaults())
	if err != nil {
		log.Fatalf("Get cards ERROR: %s", err.Error())
	}

	if len(cards) != 0 {
		for _, card := range cards {
			log.Println(card.Name, card.URL)
		}
		log.Fatalln("Exist cards in In progress")
	}
	log.Println("OK")
}
