package trello

import (
	"fmt"
	"log"
	"os"

	"github.com/ozgio/strutil"
	"github.com/sirupsen/logrus"

	"github.com/adlio/trello"
	"github.com/joho/godotenv"
)

var (
	Boards = []string{"Template", "LivGolf", "Olympics", "somosunaola", "bedfiles", "Go!", "herrumbrevivo", "LaPorrA", "Grow"}
)

func GetBoardIDFromBoardName(boardName string) (boardID string) {

	// Normalize board name
	boardNameSlug := strutil.Slugify(boardName)
	// fmt.Println("Board Name:", boardName)
	// fmt.Println("Board Name Slug:", boardNameSlug)
	switch boardNameSlug {
	case "template":
		return "66f705b52988ebc62a01e29c"
	case "livgolf":
		return "66462e4bca554f21f82b8e4a"
	case "olympics":
		return "617c56690fcb27430e740522"
	case "somosunaola":
		return "64f3290c7cacf838026f13d4"
	case "bedfiles":
		return "66ae6b96c6c1b51cd9f04d90"
	case "go!":
		return "64bfb38d1819c09c10cc406d"
	case "herrumbrevivo":
		return "59cd45ec717c12d2a2269df7"
	case "laporra":
		return "66757258661e38a299a7e687"
	case "grow":
		return "64bfd0b65f48b2e94a085631"
	default:
		fmt.Println("Unknown board name: ", boardName)
		os.Exit(1)
	}
	return boardName
}

// CustomGetCards obtiene las tarjetas de una lista incluyendo los campos personalizados
func CustomGetCards(client *trello.Client, listID string, extraArgs trello.Arguments) ([]*trello.Card, error) {
	path := fmt.Sprintf("lists/%s/cards", listID)
	// Añadir el parámetro customFieldItems=true
	if extraArgs == nil {
		extraArgs = trello.Arguments{"customFieldItems": "true"}
	}

	var cards []*trello.Card
	err := client.Get(path, extraArgs, &cards)
	if err != nil {
		return nil, err
	}
	// panic(fmt.Sprintf("%+v\n", cards[0]))

	return cards, nil
}
func GetListAndCardsFromBoardAndMonth(client *trello.Client, board *trello.Board, month string) (trello.List, []*trello.Card, error) {

	// Get the list of cards for the month
	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println(err)
	}
	var list trello.List
	var monthCards = []*trello.Card{}
	for _, list := range lists {
		// Take the list name
		listName := list.Name
		listNameSlug := strutil.Slugify(listName)
		monthSlug := strutil.Slugify(month)
		if listNameSlug == monthSlug {
			// Get cards from list
			// cards, err := list.GetCards(trello.Defaults())

			// if err != nil {
			// 	fmt.Println(err)
			// }

			cards, err := CustomGetCards(client, list.ID, nil)
			if err != nil {
				fmt.Println("Error obteniendo tarjetas con campos personalizados:", err)
				return *list, monthCards, err
			}

			for _, card := range cards {
				monthCards = append(monthCards, card)
			}
		}
	}

	return list, monthCards, nil
}
func InitTrello(boardName string) (*trello.Client, *trello.Member, *trello.Board) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	appKey := os.Getenv("TRELLO_APP_KEY")
	token := os.Getenv("TRELLO_TOKEN")

	if appKey == "" || token == "" {
		fmt.Println("Missing TRELLO_APP_KEY or TRELLO_TOKEN")
		os.Exit(1)
	}

	client := trello.NewClient(appKey, token)
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	// client.Logger = logger

	// Get Trello member
	extraArgs := make(map[string]string)
	member, err := client.GetMember("juanandresmorenorubio", extraArgs)
	if err != nil || member == nil {
		log.Fatalf("Error fetching Trello member: %v", err) // Ensure we handle the error properly
	}

	// Get Trello board
	boardID := GetBoardIDFromBoardName(boardName)
	boardT, err := client.GetBoard(boardID, extraArgs)
	if err != nil || boardT == nil {
		log.Fatalf("Error fetching Trello board: %v", err) // Ensure we handle the error properly
	}

	return client, member, boardT
}
