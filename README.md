## Snake and Ladder Game (Golang CLI)

This is a command-line Snake and Ladder game simulation written in Go, using the [Cobra](https://github.com/spf13/cobra) library for CLI functionality.

### Features
- **Configurable Board:** Set the board size, number of dice, number of players, ladders, and snakes.
- **Randomized Dice Rolls:** Uses Go's `math/rand` for dice simulation.
- **Custom Ladders & Snakes:** Define ladders and snakes as start-end pairs.
- **Multiple Players:** Supports any number of players.

### Technical Details
- **Language:** Go
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Main File:** `main.go`
- **Game Logic:**
	- The board, dice, players, ladders, and snakes are all configurable via CLI flags.
	- Ladders and snakes are provided as key-value pairs (e.g., `-l 2=38,7=14` for ladders).
	- Each player rolls the dice in turn; landing on a ladder moves up, on a snake moves down.
	- The first player to reach the final square wins.

### CLI Usage

#### Build
```sh
go build -o main main.go
```

#### Run
```sh
./main \
	--board 100 \
	--numOfDice 2 \
	--numOfPlayers 5 \
	--ladders 12=25,20=29 \
	--snakes 26=20,15=5
```

#### Short Flags
- `-b` : Board size (e.g., 100)
- `-d` : Number of dice (default: 1)
- `-p` : Number of players (default: 2)
- `-l` : Ladders (e.g., 2=38,7=14)
- `-s` : Snakes (e.g., 16=6,48=26)

#### Example
```sh
./main -b 100 -d 2 -p 5 -l 12=25,20=29 -s 26=20,15=5
```

### Output
The game prints the board size, dice, players, ladders, and snakes, then simulates turns until a player wins. Each move is printed to the console.

