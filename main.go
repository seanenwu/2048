package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
	"strings"
)

const Size = 4

type Node struct {
	Arr  [Size][Size]int
	Point int
	Next *Node
}

func main() {
	var head, tail, ptr *Node
	var highscore, score, step int
	var game [Size][Size]int
	start := true
	rand.Seed(time.Now().UnixNano())

	for{
		if(start){
			ptr = &Node{}
			head = ptr
			tail = ptr
			tail.Next = nil
			for i := 0; i < Size; i++ {
				for j := 0; j < Size; j++ {
					game[i][j] = 0
				}
			}
			randomgame(&game);
			score = 0
			step = 0
			arrCpy(&ptr.Arr, &game)
			ptr.Point = 0
			data, err := os.ReadFile("number.txt")
			if err != nil {
				highscore=0;
				os.WriteFile("2048hsc.txt", []byte(strconv.Itoa(highscore)), 0644)
			}else{
				highscore, _ = strconv.Atoi(strings.TrimSpace(string(data)))
			}
			start = false
		}

		showGame(highscore,game,score,step)
		if step != 0 {
			fmt.Println("Press < or , to undo")
		}
		if ptr.Next != nil {
			fmt.Println("Press > or . to redo")
		}
		fmt.Println("Press R to restart")
		if gameOver(game) {
			fmt.Println("Game Over!")
		}

		var input string
		fmt.Scanln(&input)
		switch input {
			case "A", "a":
				pushLeft(&game, &score)
				if(arrDif(&ptr.Arr, &game)){
					randomgame(&game)
					if score > highscore {
						highscore = score
						os.WriteFile("2048hsc.txt", []byte(strconv.Itoa(highscore)), 0644)
					}
					step++
					if ptr.Next != nil {
						ptr.Next = nil
					}
					tail = &Node{}
					ptr.Next = tail
					tail.Next = nil
					ptr = tail
					ptr.Point = score
					arrCpy(&ptr.Arr, &game)
				}
			case "D", "d":
				pushRight(&game, &score)
				if(arrDif(&ptr.Arr, &game)){
					randomgame(&game)
					if score > highscore {
						highscore = score
						os.WriteFile("2048hsc.txt", []byte(strconv.Itoa(highscore)), 0644)
					}
					step++
					if ptr.Next != nil {
						ptr.Next = nil
					}
					tail = &Node{}
					ptr.Next = tail
					tail.Next = nil
					ptr = tail
					ptr.Point = score
					arrCpy(&ptr.Arr, &game)
				}
			case "W", "w":
				pushUp(&game, &score)
				if(arrDif(&ptr.Arr, &game)){
					randomgame(&game)
					if score > highscore {
						highscore = score
						os.WriteFile("2048hsc.txt", []byte(strconv.Itoa(highscore)), 0644)
					}
					step++
					if ptr.Next != nil {
						ptr.Next = nil
					}
					tail = &Node{}
					ptr.Next = tail
					tail.Next = nil
					ptr = tail
					ptr.Point = score
					arrCpy(&ptr.Arr, &game)
				}
			case "S", "s":
				pushDown(&game, &score)
				if(arrDif(&ptr.Arr, &game)){
					randomgame(&game)
					if score > highscore {
						highscore = score
						os.WriteFile("2048hsc.txt", []byte(strconv.Itoa(highscore)), 0644)
					}
					step++
					if ptr.Next != nil {
						ptr.Next = nil
					}
					tail = &Node{}
					ptr.Next = tail
					tail.Next = nil
					ptr = tail
					ptr.Point = score
					arrCpy(&ptr.Arr, &game)
				}
				case "R", "r":
					start = true
				case "<", ",":
					if step > 0 {
						ptr = head
						for i := 0; i < step-1; i++ {
							ptr = ptr.Next
						}
						step--
						score = ptr.Point
						arrCpy(&game, &ptr.Arr)
					}
			case ">", ".":
				if ptr.Next != nil {
					ptr = ptr.Next
					step++
					score = ptr.Point
					arrCpy(&game, &ptr.Arr)
				}

		}
	}
}

func randomgame(game *[Size][Size]int) {
    for {
        row, col := rand.Intn(Size), rand.Intn(Size)
        if (*game)[row][col] == 0 {
            R := rand.Intn(10)
            if R == 0 {
                (*game)[row][col] = 4
            } else {
                (*game)[row][col] = 2
            }
            return
        }
    }
}

func arrCpy(dst, src *[Size][Size]int) {
    for i := 0; i < Size; i++ {
        for j := 0; j < Size; j++ {
            (*dst)[i][j] = (*src)[i][j]
        }
    }
}

func arrDif(a, b *[Size][Size]int) bool {
    for i := 0; i < Size; i++ {
        for j := 0; j < Size; j++ {
            if (*a)[i][j] != (*b)[i][j] {
                return true
            }
        }
    }
    return false
}

func showGame(highscore int, game [Size][Size]int, score, step int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("%23s\n", "highscore: "+strconv.Itoa(highscore))

	for i := 0; i < Size; i++ {
		fmt.Println("     |     |     |")
		fmt.Println("     |     |     |")
		for j := 0; j < Size; j++ {
			val := game[i][j]
			valStr := ""
			if val != 0 {
				valStr = strconv.Itoa(val)
			}
			fmt.Printf("%5s", valStr)
			if j < Size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		fmt.Println("     |     |     |")
		fmt.Println("     |     |     |")
		if i < Size-1 {
			fmt.Println("-----+-----+-----+-----")
		}
	}
	fmt.Printf("Score: %d\n", score)
	fmt.Printf("Step: %d\n", step)
	fmt.Println("Press WASD to control")
}

func pushLeft(game *[Size][Size]int, score *int) {
	for i := 0; i < Size; i++ {
		var combine = false
		var n = 0
		for j := 0; j < Size; j++ {
			if (*game)[i][j] != 0 {
				if n == 0 || combine {
					(*game)[i][n] = (*game)[i][j]
					n++
					combine = false
				} else {
					if (*game)[i][n-1] == (*game)[i][j] {
						(*game)[i][n-1] *= 2
						*score += (*game)[i][n-1]
						combine = true
					} else {
						(*game)[i][n] = (*game)[i][j]
						n++
					}
				}
			}
		}
		for j := n; j < Size; j++ {
			(*game)[i][j] = 0
		}
	}
}

func pushRight(game *[Size][Size]int, score *int) {
	tmp := [Size][Size]int{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			tmp[i][j] = (*game)[i][Size-1-j]
		}
	}
	pushLeft(&tmp, score)
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			(*game)[i][j] = tmp[i][Size-1-j]
		}
	}
}

func pushUp(game *[Size][Size]int, score *int) {
	tmp := [Size][Size]int{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			tmp[i][j] = (*game)[j][i]
		}
	}
	pushLeft(&tmp, score)
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			(*game)[i][j] = tmp[j][i]
		}
	}
}

func pushDown(game *[Size][Size]int, score *int) {
	tmp := [Size][Size]int{}
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			tmp[i][j] = (*game)[j][i]
		}
	}
	pushRight(&tmp, score)
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			(*game)[i][j] = tmp[j][i]
		}
	}
}

func gameOver(game [Size][Size]int) bool {
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if game[i][j] == 0 {
				return false
			}
		}
	}

	for i := 0; i < Size; i++ {
		for j := 0; j < Size-1; j++ {
			if game[i][j] == game[i][j+1] || game[j][i] == game[j+1][i] {
				return false
			}
		}
	}
	return true
}
//10.21