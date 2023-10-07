# Tic-Tac-Toe Game

This is a simple command-line Tic-Tac-Toe game implemented in Go using the Bubble Tea framework, along with BubbleZone for mouse controls. You can play the game in your terminal with your keyboard or mouse.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/tic-tac-toe.git
   ```

2. Change the directory to the project folder:

   ```bash
   cd tic-tac-toe
   ```

3. Build and run the game:

   ```bash
   go build
   ./tic-tac-toe
   ```

4. Use the following controls to play:

   - Arrow keys: Move the cursor
   - Enter: Place your symbol (X or O) on the selected cell
   - Q: Quit the game

   Alternatively, you can use your mouse to interact with the game board:

   - Move your mouse cursor to select a cell.
   - Left-click to place your symbol in the selected cell.

5. Win the game by getting three of your symbols in a row horizontally, vertically, or diagonally.

6. Enjoy playing Tic-Tac-Toe!

## Game Rules

- The game is played on a 3x3 grid.
- Player X starts the game.
- Players take turns to place their symbol (X or O) on an empty cell.
- The game ends when one player wins or when the board is full (a draw).
- To win, a player must have three of their symbols in a row horizontally, vertically, or diagonally.
- If the board is full and no player has won, the game is a draw.

## Controls

Keyboard Controls:
- Use arrow keys to move the cursor to select a cell.
- Press Enter to place your symbol in the selected cell.
- Press Q to quit the game at any time.

Mouse Controls (BubbleZone):
- Move your mouse cursor to select a cell.
- Left-click to place your symbol in the selected cell.
  
<img src="https://go.dev/images/gophers/ladder.svg" width="48" alt="Go Gopher climbing a ladder." align="right">

## Todo

- [ ] Implement AI for single-player mode

Enjoy the game and challenge your friends or the AI (coming soon) in this classic game of Tic-Tac-Toe!
