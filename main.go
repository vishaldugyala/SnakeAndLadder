package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type game struct {
	players   []*player
	numOfDice int
	board     int
	ladders   map[int]int // Stores the ladder position from low index to higher index
	snakes    map[int]int // Stores the snake position from lower index to higher index
}

type player struct {
	name     string
	position int
}

// Define flag variables at the package level
var (
	boardSize    int
	numOfDice    int
	numOfPlayers int
	players      []*player
	ladders      map[string]string
	snakes       map[string]string
)

// Define the main game command
var gameCommand = &cobra.Command{
	Use:   "SnakeAndLadder",
	Short: "snake and ladder application cli",
	Long:  "A command line application developed in golang for snake and ladder game",
	Run:   runGame,
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Add flags to the command
	gameCommand.PersistentFlags().IntVarP(&boardSize, "board", "b", 0, "defines the board size for the game")
	gameCommand.PersistentFlags().IntVarP(&numOfDice, "numOfDice", "d", 1, "defines the number of dice to be used in the game")
	gameCommand.PersistentFlags().IntVarP(&numOfPlayers, "numOfPlayers", "p", 2, "defines the number of players")
	gameCommand.PersistentFlags().StringToStringVarP(&ladders, "ladders", "l", make(map[string]string), "defines the ladders")
	gameCommand.PersistentFlags().StringToStringVarP(&snakes, "snakes", "s", make(map[string]string), "defines the snakes")

	// Execute the command
	if err := gameCommand.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

// Create the game board based on parsed flags
func createBoard() *game {
	return &game{
		numOfDice: numOfDice,
		board:     boardSize,
		ladders:   stringToInt(ladders),
		snakes:    stringToInt(snakes),
		players:   players,
	}
}

// Convert string-to-string map to int-to-int map
func stringToInt(input map[string]string) map[int]int {
	intMap := make(map[int]int)
	for k, v := range input {
		keyInt, _ := strconv.Atoi(k)
		valueInt, _ := strconv.Atoi(v)
		intMap[keyInt] = valueInt
	}
	return intMap
}

// Run the game logic
func runGame(cmd *cobra.Command, args []string) {

	players = make([]*player, numOfPlayers)
	for i := 0; i < numOfPlayers; i++ {
		players[i] = &player{
			position: 0,
		}
	}

	game := createBoard()

	// Print the contents of the game configuration
	fmt.Println("Board Size:", game.board)
	fmt.Println("Number of Dice:", game.numOfDice)
	fmt.Println("Number of Players:", numOfPlayers)
	fmt.Println("Players", game.players)
	fmt.Println("Ladders:", game.ladders)
	fmt.Println("Snakes:", game.snakes)

	for {
		//simulate the game
		for idx, player := range game.players {
			numRolled := game.rollDice()
			//fmt.Printf("Player :%d has rolled %d and current position is :%d\n", idx, numRolled, player.position)
			if player.position+numRolled <= boardSize {
				player.position += numRolled
			}
			if player.position == boardSize {
				fmt.Printf("Player :%d has won the game\n", idx)
				return
			}
			if value, ok := game.ladders[player.position]; ok {
				player.position = value
				fmt.Printf("Player :%d has taken a ladder to position:%d\n", idx, player.position)
			}
			if value, ok := game.snakes[player.position]; ok {
				player.position = value
				fmt.Printf("Player :%d has been bitten by snake :%d\n", idx, player.position)
			}
			fmt.Printf("Player :%d has rolled %d and new position is :%d\n", idx, numRolled, player.position)
		}
	}

}

func (game *game) rollDice() int {
	numberRolled := 0
	for i := 0; i < game.numOfDice; i++ {
		numberRolled += rand.Int() % 6
	}
	//fmt.Printf("%d dices has been rolled and the output is %d\n", game.numOfDice, numberRolled)
	return numberRolled
}
