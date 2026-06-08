package main

import (
	"fmt"
	"os"
)

var (
	ans   [9][9]byte
	count int
)

func main() {
	if len(os.Args) != 10 {
		fail()
	}
	var b [9][9]byte
	for i, s := range os.Args[1:] {
		if len(s) != 9 {
			fail()
		}
		for j := 0; j < 9; j++ {
			if (s[j] < '1' || s[j] > '9') && s[j] != '.' {
				fail()
			}
			b[i][j] = s[j]
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] != '.' {
				c := b[i][j]
				b[i][j] = '.'
				if !valid(b, i, j, c) {
					fail()
				}
				b[i][j] = c
			}
		}
	}

	solve(&b)

	if count == 1 {
		for _, r := range ans {
			s := ""
			for _, c := range r {
				s += string(c) + " "
			}
			fmt.Println(s[:17]) // Trims trailing space
		}
	} else {
		fail()
	}
}

func fail() {
	fmt.Println("Error")
	os.Exit(0)
}

func valid(b [9][9]byte, r, c int, v byte) bool {
	for i := 0; i < 9; i++ {
		if b[r][i] == v || b[i][c] == v || b[r/3*3+i/3][c/3*3+i%3] == v {
			return false
		}
	}
	return true
}

func solve(b *[9][9]byte) {
	if count > 1 {





		
	================================================================================ SUDOKU SOLVER CODE EXPLANATION - MADE EASY FOR EVERYONE!

Imagine a Sudoku puzzle is a big grid made of 9 rows and 9 columns, creating 81 tiny square toy boxes. Some toy boxes have number blocks from 1 to 9 in them, and some are empty, marked with a dot ('.').

Our program is a smart robot whose job is to fill the empty boxes with the right numbers without breaking the rules, and make sure there is only ONE way to do it.

Let's look at the code line by line to see how it works!
PART 1: THE MEMORY LOGS (Global Variables)

var ( ans [9][9]byte count int )

    "var (...)" tells the computer to build some memory storage boxes.
    "ans [9][9]byte" is a secret backup grid. When the robot finds a perfect finished board, it takes a snapshot picture of it and saves it here.
    "count int" is our scoreboard. It counts how many different ways the puzzle can be solved. If it finds more than 1 way, the puzzle is broken!

PART 2: THE BOSS FUNCTION (func main)

func main() {

    This is the main button. When you start the program, this runs first.

    if len(os.Args) != 10 { fail() }

    "os.Args" is the bag of information you gave the program when you started it. It must have exactly 10 items inside: the program's name plus the 9 rows of the Sudoku puzzle. If it doesn't, someone made a mistake, so we press the "fail()" emergency button.

    var b [9][9]byte

    We create a fresh, empty 9x9 grid inside the computer's mind and name it "b".

    for i, s := range os.Args[1:] {

    This loop looks at every row string you typed in, one by one. "i" is the row number (0 to 8), and "s" is the text string for that row.

      if len(s) != 9 { fail() }

    If any row doesn't have exactly 9 characters, it is broken! We call "fail()".

      for j := 0; j < 9; j++ {

    Now we look closer at each individual character inside that row string, going from left to right. "j" is the column number (0 to 8).

          if (s[j] < '1' || s[j] > '9') && s[j] != '.' { fail() }
          b[i][j] = s[j]

    We look at the character "s[j]". It MUST be a number block from '1' to '9', or an empty dot '.'. If it's something weird like an 'X' or an 'A', we hit "fail()". If it's valid, we put it into our board grid "b" at row "i", column "j". } }

    for i := 0; i < 9; i++ { for j := 0; j < 9; j++ {

    Before solving, we need to double-check that the starting puzzle doesn't already break the rules (like having two '5' blocks in the same row). We look at every square box on the board.

          if b[i][j] != '.' {
              c := b[i][j]
              b[i][j] = '.'

    If a box is NOT empty, we remember its number inside a temporary variable "c". Then we lift the block out of the grid, making it a dot '.' temporarily. (We do this because if we leave it there, our rule-checker will see it and think it's a duplicate block!).

              if !valid(b, i, j, c) { fail() }
              b[i][j] = c

    We ask our referee function "valid" if it's okay to place the number "c" in this box. If the referee says "No!", the starting puzzle is illegal, so we hit "fail()". If it's "Yes!", we put the block back safely. } } }

    solve(&b)

    Now we hand the board over to our magical "solve" robot. The "&" symbol means we give it the real board so it can write directly on it.

    if count == 1 {

    Once the robot is done looking at all paths, we check our scoreboard. A good Sudoku must have EXACTLY ONE unique solution.

      for _, r := range ans {
          s := ""
          for _, c := range r { s += string(c) + " " }
          fmt.Println(s[:17])
      }

    If the count is exactly 1, we read our saved picture "ans" row by row ("r"). We glue the characters ("c") together into a text line with spaces between them, trim off the very last extra space, and print the finished puzzle onto the screen!

    } else { fail() } }

    If the scoreboard is 0 (unsolvable) or more than 1 (too many ways to solve), it's a bad puzzle, so we call "fail()".

PART 3: THE EMERGENCY BUTTON (func fail)

func fail() { fmt.Println("Error") os.Exit(0) }

    Whenever this helper is called, it loudly prints "Error" to the screen and instantly shuts down the whole program.

PART 4: THE REFEREE (func valid)

func valid(b [9][9]byte, r, c int, v byte) bool {

    This is our rules referee. It checks if putting a number block "v" at row "r" and column "c" is allowed.

    for i := 0; i < 9; i++ { if b[r][i] == v || b[i][c] == v || b[r/33+i/3][c/33+i%3] == v { return false } } return true }

    The loop runs 9 times (i from 0 to 8) to check all 3 Sudoku rules at once:
        "b[r][i] == v" checks if the number "v" is already somewhere in that row.
        "b[i][c] == v" checks if "v" is already somewhere in that column.
        "b[r/33+i/3][c/33+i%3] == v" uses clever integer math to find the 3x3 small square box that holds this cell and checks if "v" is already inside it.

    If any of these checks find a match, the referee shouts "False!" (Illegal move!).

    If the loop finishes cleanly, the referee says "True!" (It's safe to play!).

PART 5: THE PUZZLE ROBOT (func solve)

func solve(b *[9][9]byte) {

    This is the smart robot engine. It uses a trick called "backtracking" (try a number, move forward, and if you get stuck, erase it and try another one).

    if count > 1 { return }

    If our scoreboard already counted more than 1 solution, the puzzle is invalid. The robot stops searching immediately to save time.

    for i := 0; i < 9; i++ { for j := 0; j < 9; j++ { if b[i][j] == '.' {

    The robot scans the board like reading a book. It looks for the first empty dot.

              for v := byte('1'); v <= '9'; v++ {

    When it finds an empty spot, it tries guessing number blocks from '1' to '9'.

                  if valid(*b, i, j, v) {
                      b[i][j] = v
                      solve(b)
                      b[i][j] = '.'
                  }
              }

    "if valid(...)" asks the referee if the guess "v" is legal.

    If yes, the robot writes "v" into that box.

    "solve(b)" tells the robot to call itself recursively to go find the NEXT empty box.

    "b[i][j] = '.'" is the magic rewind step! If the robot hits a dead end later on and discovers this guess was a mistake, it returns back here, wipes the box clean to a '.', and loops to try the next higher number block.

              return

    If the robot tries all numbers from 1 to 9 in this box and none work out, it means an earlier choice was wrong. It returns to step backward. } } }

    count++ if count == 1 { ans = *b } }

    If the robot checks the entire board and can't find ANY empty dots ('.'), it means the board is completely and perfectly filled! It adds 1 point to our scoreboard.

    If this is our very first success, it takes a snapshot copy of this winning board and stores it in our global "ans" box before looking for other solutions. ================================================================================	